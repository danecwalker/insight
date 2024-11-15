package store

import (
	"database/sql"

	user_contracts "github.com/danecwalker/insight/core/internal/users/contracts"
	user_repos "github.com/danecwalker/insight/core/internal/users/repos"
)

type Storage struct {
	Users user_contracts.UserStore
}

func NewSqliteStorage(db *sql.DB) *Storage {
	users := user_repos.NewSqliteUserStore(db)

	users.CreateTables()

	return &Storage{
		Users: users,
	}
}
