package entity

import "time"

type AgendaService struct {
	AgendaStorage *storage
}

// new an agendaservice.
func NewAgendaService() *AgendaService {
	return &AgendaService{Storage()}
}

// agenda login
// check if the username match password.
func (as *AgendaService) UserLogin(username string, password string) bool {
	res := as.AgendaStorage.QueryUser(func(user User) bool {
		if username == user.Name {
			return true
		}
		return false
	})

	if len(res) > 0 {
		as.AgendaStorage.saveCurUser(res[0])
	}
	return true
}

// agenda logout
// user logout.
func (as *AgendaService) UserLogout() bool {

	return true
}

// agenda register
// regist a user.
func (as *AgendaService) UserRegister(username string, password string, email string, phone string) bool {

	user := NewUser(username, password, email, phone)
	userList := as.AgendaStorage.QueryUser(func(user User) bool {
		if user.Name == username {
			return true
		}
		return false
	})

	if len(userList) > 0 {
		return false
	}

	as.AgendaStorage.AddUser(user)
	return true
}

// agenda delUser
// delete a user.
func (as *AgendaService) DeleteUser(username string, password string) bool {
	var ret int
	ret = as.AgendaStorage.DeleteUser(func(user User) bool {
		if user.Name == username {
			return true
		}
		return false
	})
	return ret > 0
}

// agenda query
// list all users from storage
// return the list result.
func (as *AgendaService) ListAllUsers() []User {

	return as.AgendaStorage.ListAllusers()
}

// agenda cm
// add a meeting.
func (as *AgendaService) AddMeeting(sponsor string, title string, start time.Time, end time.Time, participator []string) bool {
	// invaliable data
	if start.After(end) || len(sponsor) <= 0 || len(title) <= 0 {
		return false
	}
	//query the meeting to add in database, if the meeting has existed return false
	var meeting Meeting = NewMeeting(sponsor, participator, start, end, title)
	meetings := as.AgendaStorage.QueryMeeting(func(meeting Meeting) bool {

		if meeting.Sponsor == sponsor {
			// time conflict as for sponsor
			if meeting.Start.Before(start) && start.Before(meeting.End) || start.Before(meeting.Start) && meeting.Start.Before(end) || start.Before(meeting.Start) && end.After(meeting.End) || meeting.Start.Before(start) && meeting.End.After(end) {
				return true
			}
		}
		// the sponsor is Participator in some other meetings and time conflict
		for _, par := range meeting.Participators {
			if par == sponsor {
				if meeting.Start.Before(start) && start.Before(meeting.End) || start.Before(meeting.Start) && meeting.Start.Before(end) || start.Before(meeting.Start) && end.After(meeting.End) || meeting.Start.Before(start) && meeting.End.After(end) {
					return true
				}
			}
		}
		//title conflict
		if meeting.Title == title {
			return true
		}

		// the some of the participators is the sponsor of the other meeting and time conflict
		for _, par := range participator {
			if meeting.Sponsor == par {
				if meeting.Start.Before(start) && start.Before(meeting.End) || start.Before(meeting.Start) && meeting.Start.Before(end) || start.Before(meeting.Start) && end.After(meeting.End) || meeting.Start.Before(start) && meeting.End.After(end) {
					return true
				}
			}
		}

		// we collect those unsatisfied meeetings into meetings
		return false

	})
	//if those unsatisfied meetings exists, the meeting cannot be created
	if len(meetings) > 0 {
		return false
	}
	// not exists, add it into meetingList of database
	as.AgendaStorage.addMeeting(meeting)
	return true
}

// agenda qm
// query meetings by username and time interval.
func (as *AgendaService) QueryMeeting(username string, start time.Time, end time.Time) []Meeting {

	meetings := as.AgendaStorage.QueryMeeting(func(meeting Meeting) bool {
		// username is the sponsor of some meeting
		if meeting.Sponsor == username {
			if start.Before(meeting.Start) && end.After(meeting.End) || start.After(meeting.Start) && start.Before(meeting.End) || start.Before(meeting.End) && start.After(meeting.Start) {
				return true
			}
		}
		//username is a participator of some meeting
		for _, par := range meeting.Participators {
			if par == username {
				if start.Before(meeting.Start) && end.After(meeting.End) || start.After(meeting.Start) && start.Before(meeting.End) || start.Before(meeting.End) && start.After(meeting.Start) {
					return true
				}
			}
		}
		return false
	})

	return meetings
}

// agenda dm
// delete a meeting by sponsor name and title.
func (as *AgendaService) DeleteMeeting(sponsor string, title string) bool {

	num := as.AgendaStorage.DeleteMeeting(func(meeting Meeting) bool {
		if meeting.Sponsor == sponsor && meeting.Title == title {
			return true
		}
		return false
	})

	return num > 0
}

// agenda clear
// delete all meetings by sponsor.
func (as *AgendaService) DeleteAllMeetings(sponsor string) bool {
	res := as.AgendaStorage.DeleteMeeting(func(meeting Meeting) bool {
		if meeting.Sponsor == sponsor {
			return true
		}
		return false
	})
	return res > 0
}

// agenda em
// exit from a meeting by username and meeting title.
func (as *AgendaService) ExitFromMeeting(username string, title string) bool {

	res := as.AgendaStorage.ExitFromMeetings(func(meeting Meeting) int {
		if meeting.Title == title {
			for index, par := range meeting.Participators {
				if par == username {
					return index
				}
			}
			return -1
		}
		return -1
	})
	return res
}

// agenda addPart
// add a participator to a meeting
func (as *AgendaService) AddParticipator(username string, title string) bool {
	res := as.AgendaStorage.AddParticipator(username, func(m Meeting) bool {
		if m.Title == title {
			for _, par := range m.Participators {
				if par == username {
					return false
				}
			}
			return true
		}
		return false
	})
	return res
}

// agenda delPart
// delete a participator to a meeting
func (as *AgendaService) DelParticipator(username string, title string) bool {
	res := as.AgendaStorage.DelParticipator(username, func(m Meeting) int {
		if m.Title == title {
			for index, par := range m.Participators {
				if par == username {
					return index
				}
			}
			return -1
		}
		return -1
	})
	return res
}

func (as *AgendaService) IsParticipator(username string, title string) bool {
	res := as.AgendaStorage.QueryMeeting(func(m Meeting) bool {
		if m.Title == title {
			return m.IsParticipator(username)
		}
		return false
	})
	return len(res) > 0
}
