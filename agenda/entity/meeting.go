package entity

import "time"

type Meeting struct {
	Sponsor       string
	Participators []string
	Start         time.Time
	End           time.Time
	Title         string
}

func NewMeeting(sponsor string, participator []string, start time.Time, end time.Time, title string) Meeting {
	participators := make([]string, 1)
	for _, p := range participator {
		participators = append(participators, p)
	}
	return Meeting{sponsor, participators, start, end, title}
}

func (m *Meeting) IsParticipator(username string) bool {
	for _, par := range m.Participators {
		if par == username {
			return true
		}
	}
	return false
}
