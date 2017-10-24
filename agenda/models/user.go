package model

import ()

type User struct {
	username string
	password string
	email    string
	phone    string
}

type UserList struct {
	users *Mlist
}
