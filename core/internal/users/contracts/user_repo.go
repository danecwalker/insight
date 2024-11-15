package user_contracts

type UserStore interface {
	CreateTables() error
}
