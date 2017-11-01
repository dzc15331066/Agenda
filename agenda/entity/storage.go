package entity

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sync"
)

const (
	curUserFilename = "curUser.txt"
	userFilename    = "userList.json"
	meetingFilename = "meetingList.json"
)

type storage struct {
	UserList    []User
	MeetingList []Meeting
	Dirty       bool
	CurUser     User
}

var (
	s    *storage
	once sync.Once
)

// create a thread safe singleton of storage.
func Storage() *storage {
	once.Do(func() {
		s = &storage{}
		s.UserList = make([]User, 0)
		s.MeetingList = make([]Meeting, 0)
	})
	return s

}

// read users from file.
func (s *storage) readUsers() error {
	return readFromFile(&s.UserList, userFilename)
}

// read meetings from file.
func (s *storage) readMeetings() error {
	return readFromFile(&s.MeetingList, meetingFilename)
}

// read curUses from file.
func (s *storage) readCurUser() error {
	if err := readFromFile(&s.CurUser, curUserFilename); err != nil {
		return err
	} else if s.CurUser == (User{}) {
		return errors.New("Failed! please login first")
	}
	return nil
}

// write users to file.
func (s *storage) writeUsers() error {
	return writeToFile(s.UserList, userFilename)
}

// write meetings to file.
func (s *storage) writeMeetings() error {
	return writeToFile(s.MeetingList, meetingFilename)
}

// write current to file.
func (s *storage) writeCurUser() error {
	if s.CurUser != (User{}) {
		return writeToFile(s.CurUser, curUserFilename)
	}
	return nil
}

// add a user to user list.
func (s *storage) AddUser(user User) {
	s.UserList = append(s.UserList, user)
	//fmt.Println(s.UserList)
}

// query users and return a query list.
func (s *storage) QueryUser(filter func(User) bool) []User {
	var userList = make([]User, 0)
	for _, user := range s.UserList {
		if filter(user) {
			userList = append(userList, user)
		}
	}
	return userList
}

// delete users from the user list
// return the number of deleted users.
func (s *storage) DeleteUser(filter func(User) bool) int {
	ret := 0
	users := make([]User, 0)
	for _, usr := range s.UserList {
		if filter(usr) {
			ret++
		} else {
			users = append(users, usr)
		}
	}
	s.UserList = users
	//if the deleted user is the current user,
	//should clear the CurUser.txt and logout the user

	return ret
}

//query all users and return them as an Array
func (s *storage) ListAllusers() []User {
	return s.UserList //is this safe ?
}

// add a meeting to meeting list.
func (s *storage) addMeeting(m Meeting) {
	s.MeetingList = append(s.MeetingList, m)

}

// query meeting index
// return a index
func (s *storage) QueryIndexOfMeeting(filter func(Meeting) bool) int {
	index := -1
	for i, m := range s.MeetingList {
		if filter(m) {
			index = i
			break
		}
	}
	return index
}

// query meetings
// return a query list.
func (s *storage) QueryMeeting(filter func(Meeting) bool) []Meeting {
	meetings := make([]Meeting, 0)
	for _, m := range s.MeetingList {
		if filter(m) {
			meetings = append(meetings, m)
		}
	}
	return meetings
}

// delete meetings
// return the number of deleted meetings.
func (s *storage) DeleteMeeting(filter func(Meeting) bool) int {
	meetings := make([]Meeting, 0)
	var count int
	for _, m := range s.MeetingList {
		if filter(m) {
			count++
		} else {
			//not need to be deleted, so add to the meetings
			meetings = append(meetings, m)
		}
	}
	s.MeetingList = meetings
	return count
}

// exit meetings
// return true if exit successfully,false if failing
func (s *storage) ExitFromMeetings(filter func(Meeting) int) bool {

	for i, m := range s.MeetingList {
		if index := filter(m); index >= 0 {
			m.Participators = append(m.Participators[:index], m.Participators[index+1:]...)
			//如果退出后参与者为零
			if len(m.Participators) == 0 {
				s.MeetingList = append(s.MeetingList[:i], s.MeetingList[i+1:]...)
			} else {
				s.MeetingList[i].Participators = m.Participators
			}
			return true
		}
	}
	return false
}

// read json to datalist from file.
func readFromFile(datalist interface{}, filename string) error {
	//read data from file
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if len(data) > 0 {
		return json.Unmarshal(data, &datalist)
	}
	return nil

}

// write datalist to file.
func writeToFile(datalist interface{}, filename string) error {
	data, err := json.Marshal(datalist)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0666)
}

// erase current user file while logout.
func (s *storage) eraseCurUser() error {
	file, err := os.OpenFile(curUserFilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	defer file.Close()
	if err != nil {
		return err
	}

	return file.Truncate(0)
}

func (s *storage) setCurUser(user User) error {
	s.CurUser = user
	return s.writeCurUser()
}
