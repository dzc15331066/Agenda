package entity

import (
	"sync"
)

const (
	curUserFile     = "curUser.txt"
	userListFile    = "userList.json"
	meetingListFile = "meetingList.json"
)

type storage struct {
	UserList    []User
	MeetingList []Meeting
	Dirty       bool
	CurUser     *User
}

var (
	s    *storage
	once sync.Once
)

// create a thread safe singleton of storage.
func Storage() *storage {
	once.Do(func() {
		s = &storage{}
		s.readFromFile()
		s.readCurUser()
	})
	return s

}

// read
func (s *storage) readFromFile() bool {
	return true
}

func (s *storage) writeToFile() bool {
	return true
}

// add a user to user list.
func (s *storage) AddUser(user User) {
	s.UserList = append(s.UserList, user)
}

// query users and return a query list.
func (s *storage) QueryUser(filter func(User) bool) []User {
	var userList = make([]User, 1)
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
	var ret int = 0
	for index, usr := range s.UserList {
		if filter(usr) {
			s.UserList = append(s.UserList[:index], s.UserList[index+1:]...)
			ret++
		}
	}
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

// query meetings
// return a query list.
func (s *storage) QueryMeeting(filter func(Meeting) bool) []Meeting {
	meetings := make([]Meeting, 1)
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

	var count int
	for index, m := range s.MeetingList {
		if filter(m) {
			s.MeetingList = append(s.MeetingList[:index], s.MeetingList[index+1:]...)
			count++
		}
	}

	return count
}

// exit meetings
// return true if exit successfully,false if failing
func (s *storage) ExitFromMeetings(filter func(Meeting) int) bool {

	for _, m := range s.MeetingList {
		if index := filter(m); index >= 0 {
			m.Participators = append(m.Participators[:index], m.Participators[index+1:]...)
			return true
		}
	}
	return false
}

// add a participator from some meeting by meeting's title
// if successful return true, else retuern false
func (s *storage) AddParticipator(username string, filter func(Meeting) bool) bool {

	for _, m := range s.MeetingList {
		if filter(m) {
			m.Participators = append(m.Participators, username)
			return true
		}
	}
	return false
}

// delete a participator form some meeting by meeting's title
// if successful return true, else retuen false

func (s *storage) DelParticipator(username string, filter func(Meeting) int) bool {
	for _, m := range s.MeetingList {
		if index := filter(m); index >= 0 {
			m.Participators = append(m.Participators[:index], m.Participators[index+1:]...)
			return true
		}
	}
	return false
}

// sync with the files.
func (s *storage) Sync() {
	s.writeToFile()
	s.saveCurUser()
}

// read current user from file "curUser.txt".
func (s *storage) readCurUser() *User {
	return nil
}

// save current user to file "curUser.txt".
func (s *storage) saveCurUser(user User) bool {
	if s.CurUser == nil {

	} else {

	}
	return true
}
