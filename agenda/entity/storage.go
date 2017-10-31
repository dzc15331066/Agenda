package entity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
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
	CurUser     User
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
		s.CurUser = User{}
		s.readFromFile()
	})
	return s

}

// read.
func (s *storage) readFromFile() bool {
	readFromFile(&s.UserList, userFilename)
	readFromFile(&s.MeetingList, meetingFilename)
	readFromFile(&s.CurUser, curUserFilename)
	return true
}

// write.
func (s *storage) writeToFile() bool {
	writeToFile(s.UserList, userFilename)
	writeToFile(s.MeetingList, meetingFilename)
	if s.CurUser != (User{}) {
		writeToFile(s.CurUser, curUserFilename)
	}
	return true
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

	for index, m := range s.MeetingList {
		if filter(m) {
			m.Participators = append(m.Participators, username)
			fmt.Println(m.Participators)
			//这里应该引用来赋值，直接赋值改变不了
			s.MeetingList[index].Participators = m.Participators
			return true
		}
	}
	fmt.Println("here wrong")
	return false
}

// delete a participator form some meeting by meeting's title
// if successful return true, else return false.
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

}

// write datalist to file.
func writeToFile(datalist interface{}, filename string) bool {
	data, err := json.Marshal(datalist)
	if err != nil {
		log.Fatal(err)
		return false
	}

	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

// erase current user file while logout.
func eraseCurUser() bool {
	file, err := os.OpenFile(curUserFilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = file.Truncate(0)
	if err != nil {
		log.Fatal(err)
	}
	return true
}
