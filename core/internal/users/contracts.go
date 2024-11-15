package users

type UserService interface {
	CreateUser(email string) (User, error)
	GetUserByEmail(email string) (User, error)
	DeleteUserByEmail(email string) error
}

type UserRepository interface {
	Save(user User) error
	GetByEmail(email string) (User, error)
	DeleteByEmail(email string) error
}
