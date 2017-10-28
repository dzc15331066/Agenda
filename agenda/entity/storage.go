package entity

import (
	"sync"
)

type storage struct {
}

var (
	s    *storage
	once sync.Once
)

// create a thread safe singleton of storage.
func Storage() *storage {
	once.Do(func() {
		s = &storage{}
	})
	return s

}

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

// sync with the file

func (s *storage) Sync() {

}
