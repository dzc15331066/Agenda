package entity

type User struct {
	Name     string
	Password string
	Email    string
	Phone    string
}

func NewUser(name string, password string, email string, phone string) *User {
	return &User{name, password, email, phone}
}
