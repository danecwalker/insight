package user_repos

import "database/sql"

type UsersSqlite struct {
	db *sql.DB
}

func NewSqliteUserStore(db *sql.DB) *UsersSqlite {
	return &UsersSqlite{
		db: db,
	}
}

func (u *UsersSqlite) CreateTables() error {
	_, err := u.db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL
	);`)
	if err != nil {
		return err
	}

	return nil
}
