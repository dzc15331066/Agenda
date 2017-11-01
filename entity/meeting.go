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
	participators := make([]string, 0)
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
func StringToDate(str string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	date, err := time.ParseInLocation(format, str, loc)
	return date, err
}

// change date to string format "yyyy-mm-dd/hh:mm".
func DateToString(date time.Time) string {
	return date.Format(format)
}

// check whether meeting conflicts the dates.
func (meeting *Meeting) OverLap(start time.Time, end time.Time) bool {
	return ((meeting.Start.Before(start) || meeting.Start.Equal(start)) && (start.Before(meeting.End) || start.Equal(meeting.End))) || ((start.Before(meeting.Start) || start.Equal(meeting.Start)) && (meeting.Start.Before(end) || meeting.Start.Before(end))) || ((start.Before(meeting.Start) || start.Equal(meeting.Start)) && (end.After(meeting.End) || end.Equal(meeting.End))) || ((meeting.Start.Before(start) || meeting.Start.Equal(start)) && (meeting.End.Equal(end) || meeting.End.After(end)))
}
