package storage

import (
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
