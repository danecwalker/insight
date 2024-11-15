package store

import (
	"database/sql"

	"github.com/danecwalker/insight/core/internal/magic"
	"github.com/danecwalker/insight/core/internal/users"
	"github.com/danecwalker/insight/core/internal/users/infrastructure"
)

type Storage struct {
	Users users.UserStore
	Magic magic.MagicStore
}

func NewSqliteStorage(db *sql.DB) *Storage {
	users := infrastructure.NewSqliteUserStore(db)

	return &Storage{
		Users: users,
	}
}
