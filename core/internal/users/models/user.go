package models

type User struct {
	id       int
	email    string
	password string
}

func NewUser(email, password string) User {
	return User{
		email:    email,
		password: password,
	}
}

func (u *User) SetID(id int) {
	u.id = id
}

func (u User) GetID() int {
	return u.id
}

func (u User) GetEmail() string {
	return u.email
}

func (u User) GetPassword() string {
	return u.password
}
