package users

type User struct {
	id    int
	email string
}

func NewUser(email string) *User {
	return &User{
		email: email,
	}
}

func (u *User) ID() int {
	return u.id
}

func (u *User) Email() string {
	return u.email
}
