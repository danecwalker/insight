package infrastructure

import "database/sql"

type MagicSqlite struct {
	db *sql.DB
}

func NewSqliteMagicStore(db *sql.DB) *MagicSqlite {
	return &MagicSqlite{
		db: db,
	}
}
