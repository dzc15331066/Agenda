package storage

import (
	"../user/"
	"sync"
)

type storage struct {
}

var (
	s    *storage
	once sync.Once
)

//create a thread safe singleton of storage
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

func (s *storage) AddUser(u User) {

}
