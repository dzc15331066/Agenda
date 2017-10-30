package entity

import "time"

const format = "2006-01-02/15:04"

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

func (m *Meeting) ParticipatorIndex(username string) int {
	for i, par := range m.Participators {
		if par == username {
			return i
		}
	}
	return -1
}

// change string "yyyy-mm-dd/hh:mm" to Date.
func StringToDate(str string) time.Time {
	loc, _ := time.LoadLocation("Local")
	date, _ := time.ParseInLocation(format, str, loc)
	return date
}

// change date to string format "yyyy-mm-dd/hh:mm".
func DateToString(date time.Time) string {
	return date.Format(format)
}
