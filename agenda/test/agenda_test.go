package test

import (
	"Agenda/agenda/entity"
	"testing"
)

func TestUserRegister(t *testing.T) {
	as := entity.NewAgendaService()
	name := "Username"
	pass := "pass"
	email := "1235@qq.com"
	phone := "1315766578"
	t.Log(as.UserRegister(name, pass, email, phone))
}
