package users

type UserStore interface {
	CreateTables() error
}
