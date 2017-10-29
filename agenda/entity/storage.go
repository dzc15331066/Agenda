package entity

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"sync"
)

const (
	curUserFilename = "curUser.txt"
	userFilename    = "userList.json"
	meetingFilename = "meetingList.json"
	logfilename     = "log.txt"
)

var log = logrus.New()

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
		InitLog()
		s = &storage{}
		s.UserList = make([]User, 0)
		s.MeetingList = make([]Meeting, 0)
		s.CurUser = &User{}
		s.readFromFile()
	})
	return s

}

// read.
func (s *storage) readFromFile() bool {
	readFromFile(&s.UserList, userFilename)
	readFromFile(&s.MeetingList, meetingFilename)
	readFromFile(s.CurUser, curUserFilename)
	return true
}

// write.
func (s *storage) writeToFile() bool {
	if len(s.MeetingList) > 0 {
		writeToFile(s.UserList, userFilename)
	}
	if len(s.UserList) > 0 {
		writeToFile(s.MeetingList, meetingFilename)
	}
	if s.CurUser != nil {
		writeToFile(*s.CurUser, curUserFilename)
	}
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
	log.Out.(*os.File).Close()
}

// intit a log
func InitLog() {
	logfile, err := os.OpenFile(logfilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		log.Out = logfile
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}

// read json to datalist from file.
func readFromFile(datalist interface{}, filename string) {
	//read data from file
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	if len(data) > 0 {
		err = json.Unmarshal(data, &datalist)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

}

// write datalist to file.
func writeToFile(datalist interface{}, filename string) {
	data, err := json.Marshal(datalist)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		log.Fatal(err)
	}

}
