package meeting

import "time"

type Meeting struct {
	Sponsor       string
	Participators []string
	Start         time.Time
	End           time.Time
	Title         string
}

func NewMeeting(sponsor string, participator string, start time.Time, end time.Time, title string) *Meeting {
	participators := make([]string, 1)
	participators = append(participators, participator)
	return &Meeting{sponsor, participators, start, end, title}
}

func (m *Meeting) IsParticipator(username string) bool {
	//add your codes here
	return true
}
