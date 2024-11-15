package main

import (
	"database/sql"

	"github.com/danecwalker/insight/core/internal/magic"
	"github.com/danecwalker/insight/core/internal/users"
)

type Storage struct {
	Users users.UserRepository
	Magic magic.MagicRepository
}

func SetupStorage(db *sql.DB) *Storage {
	// Setup the storage
	return &Storage{
		Users: users.NewSqliteRepository(db),
		Magic: magic.NewSqliteRepository(db),
	}
}
