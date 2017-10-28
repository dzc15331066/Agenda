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
func (s *storage) AddUser(u User) {

}

// query users and return a query list.
func (s *storage) QueryUser(filter func(User) bool) []User {
	return nil
}

// delete users from the user list
// return the number of deleted users.
func (s *storage) DeleteUser(filter func(User) bool) int {
	return 0
}

// add a meeting to meeting list.
func (s *storage) addMeeting(m Meeting) {

}

// query meetings
// return a query list.
func (s *storage) QueryMeeting(filter func(Meeting) bool) []Meeting {
	return nil
}

// delete meetings
// return the number of deleted meetings.
func (s *storage) DeleteMeeting(filter func(Meeting) bool) int {
	return 0
}

// sync with the files.
func (s *storage) Sync() {
	s.writeToFile()
	s.saveCurUser()
}

// read current user from file "curUser.txt".
func (s *storage) readCurUser() *User {
	return &User{}
}

// save current user to file "curUser.txt".
func (s *storage) saveCurUser() bool {
	if s.CurUser == nil {

	} else {

	}
	return true
}
