package entity

import (
	"errors"
)

var (
	nullAgumentError = errors.New("[error]: aguments shouldn't be null")
	timeFormatError  = errors.New("[error]: time format should be like yyyy-mm-dd/hh:mm")
)

type AgendaService struct {
	AgendaStorage *storage
}

// new an agendaservice.
func NewAgendaService() *AgendaService {
	return &AgendaService{Storage()}
}

// agenda login
// check if the username match password.
func (as *AgendaService) UserLogin(username string, password string) error {

	if err := as.AgendaStorage.readUsers(); err != nil {
		return err
	}
	res := as.AgendaStorage.QueryUser(func(user User) bool {
		return username == user.Name && password == user.Password
	})

	//fmt.Println(res)

	if len(res) > 0 {
		// set current user
		//fmt.Println(res[0])
		return as.AgendaStorage.setCurUser(res[0])

	}
	return errors.New("[error]: Invalid username or password")
}

// agenda logout
// user logout.
func (as *AgendaService) UserLogout() error {
	if err := as.AgendaStorage.readCurUser(); err != nil {
		return err
	}
	return as.AgendaStorage.eraseCurUser()
}

// agenda register
// regist a user.
func (as *AgendaService) UserRegister(username string, password string, email string, phone string) error {

	if err := as.AgendaStorage.readUsers(); err != nil {
		return err
	}
	user := NewUser(username, password, email, phone)
	userList := as.AgendaStorage.QueryUser(func(user User) bool {
		return user.Name == username
	})

	if len(userList) > 0 {
		return errors.New("[error]: User has been registerd")
	}
	//register successfully, wirte to outfile
	as.AgendaStorage.AddUser(*user)
	return as.AgendaStorage.writeUsers()
}

// agenda delUser
// delete a user.
func (as *AgendaService) DeleteUser(username string, password string) error {
	var ret int

	if err := as.AgendaStorage.readUsers(); err != nil {
		return err
	}

	if err := as.AgendaStorage.readCurUser(); err != nil {
		return err
	}
	if username != as.AgendaStorage.CurUser.Name || password != as.AgendaStorage.CurUser.Password {
		return errors.New("[error]: invalid username or password")
	}
	ret = as.AgendaStorage.DeleteUser(func(user User) bool {
		return user.Name == username && user.Password == password
	})
	if ret > 0 {
		//if the deleted user is the current user,
		//logout the current user
		if username == as.AgendaStorage.CurUser.Name {
			if err := as.AgendaStorage.eraseCurUser(); err != nil {
				return err
			}
		}
		return as.AgendaStorage.writeUsers()
	}
	return nil
}

// agenda query
// list all users from storage
// return the list result.
func (as *AgendaService) ListAllUsers() ([]User, error) {
	if err := as.AgendaStorage.readUsers(); err != nil {
		return nil, err
	}

	return as.AgendaStorage.ListAllusers(), nil
}

// agenda cm
// add a meeting.
func (as *AgendaService) AddMeeting(title string, startdate string, enddate string, participators []string) error {
	if title == "" || startdate == "" || enddate == "" || participators == nil {
		return nullAgumentError
	}
	start, err := StringToDate(startdate)
	if err != nil {
		return timeFormatError
	}
	end, err := StringToDate(enddate)
	if err != nil {
		return timeFormatError
	}
	if err := as.AgendaStorage.readMeetings(); err != nil {
		return err
	}

	if err := as.AgendaStorage.readCurUser(); err != nil {
		return err
	}
	sponsor := as.AgendaStorage.CurUser.Name
	meeting := NewMeeting(sponsor, participators, start, end, title)
	//title conflict
	mIndex := as.AgendaStorage.QueryIndexOfMeeting(func(m Meeting) bool {
		return m.Title == title
	})
	if mIndex >= 0 {
		return errors.New("[error]: There is a title conflict with another title")
	}
	// invaliable date
	if start.After(end) {
		return errors.New("[error]: start date can't be later than end date")
	}

	// query current  user sponsors other meeting with time conflict
	mIndex = as.AgendaStorage.QueryIndexOfMeeting(func(m Meeting) bool {

		return m.Sponsor == sponsor && m.OverLap(start, end)
	})
	if mIndex >= 0 {
		return errors.New("[error]: You have a time conflict meeting as sponsor")
	}
	// the sponsor is Participator in some other meetings and time conflict
	mIndex = as.AgendaStorage.QueryIndexOfMeeting(func(m Meeting) bool {
		return m.ParticipatorIndex(sponsor) >= 0 && m.OverLap(start, end)
	})
	if mIndex >= 0 {
		return errors.New("[error]: You have a time conflict meeting as participator")
	}
	// some participators time conflict
	mIndex = as.AgendaStorage.QueryIndexOfMeeting(func(m Meeting) bool {
		for _, par := range participators {
			if (m.ParticipatorIndex(par) > 0 || m.Sponsor == par) && m.OverLap(start, end) {
				return true
			}

		}
		return false
	})
	if mIndex >= 0 {
		return errors.New("[error]: Can't create meeting since some participators have time conflict")
	}
	// no conflict, add it into meetingList of database)
	as.AgendaStorage.addMeeting(meeting)
	return as.AgendaStorage.writeMeetings()
}

// agenda qm
// query meetings by username and time interval.
func (as *AgendaService) QueryMeeting(startdate string, enddate string) ([]Meeting, error) {
	if startdate == "" || enddate == "" {
		return nil, nullAgumentError
	}
	if err := as.AgendaStorage.readMeetings(); err != nil {
		return nil, err
	}

	if err := as.AgendaStorage.readCurUser(); err != nil {
		return nil, err
	}
	start, err := StringToDate(startdate)
	if err != nil {
		return nil, timeFormatError
	}
	end, err := StringToDate(enddate)
	if err != nil {
		return nil, timeFormatError
	}
	username := as.AgendaStorage.CurUser.Name
	meetings := as.AgendaStorage.QueryMeeting(func(m Meeting) bool {
		// username is the sponsor or a participator of some meeting
		return (m.Sponsor == username || m.ParticipatorIndex(username) >= 0) && m.OverLap(start, end)
	})
	if len(meetings) == 0 {
		return nil, errors.New("[error]:You have no meeting between this time interval")
	}
	return meetings, nil
}

// agenda dm
// delete a meeting by sponsor name(the current user's name) and title.
func (as *AgendaService) DeleteMeeting(title string) error {
	if title == "" {
		return nullAgumentError
	}
	if err := as.AgendaStorage.readMeetings(); err != nil {
		return err
	}

	if err := as.AgendaStorage.readCurUser(); err != nil {
		return err
	}
	num := as.AgendaStorage.DeleteMeeting(func(meeting Meeting) bool {
		return meeting.Sponsor == as.AgendaStorage.CurUser.Name && meeting.Title == title
	})

	if num > 0 {
		return as.AgendaStorage.writeMeetings()
	}
	return errors.New("[error]: You have no such a meeting")
}

// agenda clear
// delete all meetings by sponsor.
func (as *AgendaService) DeleteAllMeetings() error {
	if err := as.AgendaStorage.readMeetings(); err != nil {
		return err
	}

	if err := as.AgendaStorage.readCurUser(); err != nil {
		return err
	}
	res := as.AgendaStorage.DeleteMeeting(func(meeting Meeting) bool {
		return meeting.Sponsor == as.AgendaStorage.CurUser.Name

	})
	if res > 0 {
		return as.AgendaStorage.writeMeetings()
	}
	return errors.New("[error]: You don't have any meeting")
}

// agenda em
// exit from a meeting by username and meeting title.
func (as *AgendaService) ExitFromMeeting(title string) error {
	if title == "" {
		return nullAgumentError
	}
	if err := as.AgendaStorage.readMeetings(); err != nil {
		return err
	}

	if err := as.AgendaStorage.readCurUser(); err != nil {
		return err
	}
	username := as.AgendaStorage.CurUser.Name
	res := as.AgendaStorage.ExitFromMeetings(func(meeting Meeting) int {
		if meeting.Title == title {
			return meeting.ParticipatorIndex(username)
		}
		return -1
	})
	if res {
		return as.AgendaStorage.writeMeetings()
	}
	return errors.New("[error]: You have no such a meeting")
}

// agenda addPart
// add a participator to a meeting sponsored by current user.
func (as *AgendaService) AddParticipator(usernames []string, title string) error {
	if usernames == nil || title == "" {
		return nullAgumentError
	}
	if err := as.AgendaStorage.readMeetings(); err != nil {
		return err
	}
	if err := as.AgendaStorage.readUsers(); err != nil {
		return err
	}
	if err := as.AgendaStorage.readCurUser(); err != nil {
		return err
	}
	sponsor := as.AgendaStorage.CurUser.Name
	mIndex := as.AgendaStorage.QueryIndexOfMeeting(func(m Meeting) bool {
		return m.Sponsor == sponsor && m.Title == title
	})
	if mIndex == -1 {
		return errors.New("[error]: You are not the sponsor of this meeting")
	}

	m := &as.AgendaStorage.MeetingList[mIndex]
	dirty := false
	// check registery of participators
	for _, username := range usernames {
		users := as.AgendaStorage.QueryUser(func(user User) bool {
			return username == user.Name
		})

		if len(users) == 0 {
			return errors.New("[error]: " + username + " doesn't register")
		} else {
			dirty = true
			index := as.AgendaStorage.QueryIndexOfMeeting(func(meeting Meeting) bool {
				return (meeting.ParticipatorIndex(username) > 0 || meeting.Sponsor == username) && meeting.OverLap(m.Start, m.End)
			})
			if index != -1 {
				return errors.New("[error]: " + username + " has time conflict")
			}
			m.Participators = append(m.Participators, username)
		}

		// participator index

	}
	// check
	if dirty {
		return as.AgendaStorage.writeMeetings()
	}
	return nil

}

// agenda delPart
// delete a participator to a meeting sponsored by current user.
func (as *AgendaService) DelParticipator(usernames []string, title string) error {
	if usernames == nil || title == "" {
		return nullAgumentError
	}
	if err := as.AgendaStorage.readMeetings(); err != nil {
		return err
	}
	if err := as.AgendaStorage.readUsers(); err != nil {
		return err
	}
	if err := as.AgendaStorage.readCurUser(); err != nil {
		return err
	}
	sponsor := as.AgendaStorage.CurUser.Name
	mIndex := as.AgendaStorage.QueryIndexOfMeeting(func(m Meeting) bool {
		return m.Sponsor == sponsor && m.Title == title
	})
	if mIndex == -1 {
		return errors.New("[error]: You are not the sponsor of this meeting")
	}
	m := &as.AgendaStorage.MeetingList[mIndex]
	dirty := false
	for _, username := range usernames {
		pIndex := m.ParticipatorIndex(username)
		if pIndex != -1 {
			dirty = true

			m.Participators = append(m.Participators[:pIndex], m.Participators[pIndex+1:]...)
		} else {
			return errors.New("[error]: " + username + " is not in the meeting")
		}

		// participator index

	}
	if dirty {
		return as.AgendaStorage.writeMeetings()
	}
	return nil
}
