package entity

import ()

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
	return true
}

// agenda delUser
// delete a user.
func (as *AgendaService) DeleteUser(username string, password string) bool {
	return true
}

// agenda query
// list all users from storage
// return the list result.
func (as *AgendaService) ListAllUsers() []User {
	return nil
}

// agenda cm
// add a meeting.
func (as *AgendaService) AddMeeting(sponsor string, title string, start string, end string, participator []string) {

}

// agenda qm
// query meetings by username and time interval.
func (as *AgendaService) MeetingQuery(username string, start string, end string) []Meeting {
	return nil
}

// agenda dm
// delete a meeting by sponsor name and title.
func (as *AgendaService) DeleteMeeting(sponsor string, title string) bool {
	return true
}

// agenda clear
// delete all meetings by sponsor.
func (as *AgendaService) DeleteAllMeetings(sponsor string) bool {
	return true
}

// agenda em
// exit from a meeting by username and meeting title.
func (as *AgendaService) ExitFromMeeting(username string, title string) bool {
	return true
}

// agenda addPart
// add a participator to a meeting
func (as *AgendaService) AddParticipator(username string, title string) bool {
	return true
}

// agenda delPart
// delete a participator to a meeting
func (as *AgendaService) DelParticipator(username string, title string) bool {
	return true
}

func (as *AgendaService) IsParticipator(username string, title string) bool {
	return true
}
