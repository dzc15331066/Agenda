package test

import (
	"github.com/dzc15331066/Agenda/entity"
	"io/ioutil"
	"testing"
)

func TestUserRegister(t *testing.T) {
	as := entity.NewAgendaService()
	name := "Username"
	pass := "pass"
	email := "1235@qq.com"
	phone := "1315766578"
	err := as.UserRegister(name, pass, email, phone)
	if err != nil {
		t.Error(err)
	}
	bytes, err := ioutil.ReadFile("userList.json")
	if err != nil {
		t.Error(err)
	} else {
		res := string(bytes)
		want := `[{"Name":"Username","Password":"pass","Email":"1235@qq.com","Phone":"1315766578"}]`
		if res != want {
			t.Errorf("want %q but get %q", want, res)
		}
	}

}

func TestUserLogin(t *testing.T) {
	as := entity.NewAgendaService()
	name := "Username"
	pass := "pass"
	err := as.UserLogin(name, pass)
	if err != nil {
		t.Error(err)
	}
	bytes, err := ioutil.ReadFile("curUser.txt")
	if err != nil {
		t.Error(err)
	} else {
		res := string(bytes)
		want := `{"Name":"Username","Password":"pass","Email":"1235@qq.com","Phone":"1315766578"}`
		if res != want {
			t.Errorf("want %q but get %q", want, res)
		}
	}
}

func TestUserLogout(t *testing.T) {
	as := entity.NewAgendaService()
	err := as.UserLogout()
	if err != nil {
		t.Error(err)
	}
	bytes, err := ioutil.ReadFile("curUser.txt")
	if err != nil {
		t.Error(err)
	} else {
		res := string(bytes)
		want := ""
		if res != want {
			t.Errorf("want %q but get %q", want, res)
		}
	}
}

func TestListAllUsers(t *testing.T) {
	as := entity.NewAgendaService()
	err := as.UserLogin("Username", "pass")
	users, err := as.ListAllUsers()
	eq := true
	if err != nil {
		t.Error(err)
	}
	want := make([]entity.User, 0)
	want = append(want, entity.User{"Username", "pass", "1235@qq.com", "1315766578"})
	for i := 0; i < len(want); i++ {
		if want[i] != users[i] {
			eq = true
			break
		}
	}
	if !eq {
		t.Errorf("want %q but get %q", want, users)
	}

}

func TestDeleteUser(t *testing.T) {
	as := entity.NewAgendaService()
	name := "Username"
	pass := "pass"

	err := as.DeleteUser(name, pass)
	if err != nil {
		t.Error(err)
	}
	bytes, err := ioutil.ReadFile("userList.json")
	if err != nil {
		t.Error(err)
	} else {
		res := string(bytes)
		want := "[]"
		if res != want {
			t.Errorf("want %q but get %q", want, res)
		}
	}
	bytes, err = ioutil.ReadFile("curUser.txt")
	if err != nil {
		t.Error(err)
	} else {
		res := string(bytes)
		want := ""
		if res != want {
			t.Errorf("want %q but get %q", want, res)
		}
	}

}

func TestAddMeeting(t *testing.T) {
	as := entity.NewAgendaService()
	as.UserRegister("Username", "pass", "1341415@qq.com", "12343")
	as.UserRegister("part1", "pass1", "12143@qq.com", "13443443535")
	as.UserRegister("part2", "pass2", "121434@qq.com", "1334345345")
	as.UserLogin("Username", "pass")
	title := "meeting"
	startdate := "2001-11-11/12:00"
	enddate := "2005-12-11/13:00"
	participators := []string{"part1", "part2"}
	err := as.AddMeeting(title, startdate, enddate, participators)
	if err != nil {
		t.Error(err)
	}
	bytes, err := ioutil.ReadFile("meetingList.json")
	if err != nil {
		t.Error(err)
	} else {
		res := string(bytes)
		want := `[{"Sponsor":"Username","Participators":["part1","part2"],"Start":"2001-11-11T12:00:00Z","End":"2005-12-11T13:00:00Z","Title":"meeting"}]`
		if res != want {
			t.Errorf("want %q but get %q", want, res)
		}
	}
}

func TestQueryMeeting(t *testing.T) {
	as := entity.NewAgendaService()
	ms, err := as.QueryMeeting("2001-11-12/11:00", "2004-11-12/11:00")
	if err != nil {
		t.Error(err)
	}
	start, _ := entity.StringToDate("2001-11-11/12:00")
	end, _ := entity.StringToDate("2005-12-11/13:00")
	want := make([]entity.Meeting, 0)
	m := entity.Meeting{"Username", []string{"part1", "part2"}, start, end, "meeting"}
	want = append(want, m)
	l := min(len(want), len(ms))
	for i := 0; i < l; i++ {
		if !equalMeeting(ms[i], want[i]) {
			t.Errorf("want %q but get %q", want, ms)

		}
	}

}

func equalMeeting(m1 entity.Meeting, m2 entity.Meeting) bool {
	eq := m1.Sponsor == m2.Sponsor && m1.Title == m2.Title && m1.Start == m2.Start && m1.End == m2.End
	l := min(len(m1.Participators), len(m2.Participators))
	for i := 0; i < l; i++ {
		if m1.Participators[i] != m2.Participators[i] {
			return false
		}
	}
	return eq
}

func min(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestDeleteMeeting(t *testing.T) {
	as := entity.NewAgendaService()
	err := as.DeleteMeeting("meeting")
	if err != nil {
		t.Error(err)
	}
	bytes, err := ioutil.ReadFile("meetingList.json")
	if err != nil {
		t.Error(err)
	}
	res := string(bytes)
	want := "[]"
	if res != want {
		t.Errorf("want %q but get %q", want, res)
	}
}

func TestDeleteAllMeetings(t *testing.T) {
	as := entity.NewAgendaService()
	title := "meeting"
	startdate := "2001-11-11/12:00"
	enddate := "2005-12-11/13:00"
	participators := []string{"part1", "part2"}
	as.AddMeeting(title, startdate, enddate, participators)
	err := as.DeleteAllMeetings()
	if err != nil {
		t.Error(err)
	}
	bytes, err := ioutil.ReadFile("meetingList.json")
	if err != nil {
		t.Error(err)
	}
	res := string(bytes)
	want := "[]"
	if res != want {
		t.Errorf("want %q but get %q", want, res)
	}
}

func TestExitFromMeeting(t *testing.T) {
	as := entity.NewAgendaService()
	title := "meeting"
	startdate := "2001-11-11/12:00"
	enddate := "2005-12-11/13:00"
	participators := []string{"part1", "part2"}
	as.AddMeeting(title, startdate, enddate, participators)
	err := as.UserLogin("part1", "pass1")
	if err != nil {
		t.Error(err)
	}
	err = as.ExitFromMeeting("meeting")
	if err != nil {
		t.Error(err)
	}
	bytes, err := ioutil.ReadFile("meetingList.json")
	if err != nil {
		t.Error(err)
	}
	want := `[{"Sponsor":"Username","Participators":["part2"],"Start":"2001-11-11T12:00:00Z","End":"2005-12-11T13:00:00Z","Title":"meeting"}]`
	res := string(bytes)
	if res != want {
		t.Errorf("want %q but get %q", want, res)
	}
}

func TestAddParticipator(t *testing.T) {
	as := entity.NewAgendaService()
	as.UserRegister("part3", "pass3", "22534@qq.com", "12324434")
	err := as.UserLogin("Username", "pass")
	if err != nil {
		t.Error(err)
	}
	err = as.AddParticipator([]string{"part1", "part3"}, "meeting")
	if err != nil {
		t.Error(err)
	}
	ms := as.AgendaStorage.QueryMeeting(func(m entity.Meeting) bool {
		return m.Title == "meeting"
	})
	if len(ms) == 0 {
		t.Error(`can't find a meeting named "meeting"`)
	} else {
		if ms[0].ParticipatorIndex("part3") == -1 || ms[0].ParticipatorIndex("part1") == -1 {
			t.Error(`can't add "part1" or "part3" to "meeting"`)
		}
	}
}

func TesDelParticipartors(t *testing.T) {
	as := entity.NewAgendaService()
	err := as.UserLogin("Username", "pass")
	if err != nil {
		t.Error(err)
	}
	err = as.DelParticipator([]string{"part1", "part3"}, "meeting")
	if err != nil {
		t.Error(err)
	}
	ms := as.AgendaStorage.QueryMeeting(func(m entity.Meeting) bool {
		return m.Title == "meeting"
	})
	if len(ms) == 0 {
		t.Error(`can't find a meeting named "meeting"`)
	} else {
		if ms[0].ParticipatorIndex("part3") != -1 || ms[0].ParticipatorIndex("part1") != -1 {
			t.Error("can't delete participators")
		}
	}
}
