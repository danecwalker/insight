package users

import "database/sql"

type UsersSqlite struct {
	db *sql.DB
}

func NewSqliteRepository(db *sql.DB) *UsersSqlite {
	return &UsersSqlite{
		db: db,
	}
}

func (u *UsersSqlite) Save(user User) error {
	_, err := u.db.Exec("INSERT INTO users (email) VALUES (?)", user.email)
	if err != nil {
		return err
	}

	return nil
}

func (u *UsersSqlite) GetByEmail(email string) (User, error) {
	row := u.db.QueryRow("SELECT * FROM users WHERE email = ?", email)
	var user User
	err := row.Scan(&user.id, &user.email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *UsersSqlite) DeleteByEmail(email string) error {
	_, err := u.db.Exec("DELETE FROM users WHERE email = ?", email)
	if err != nil {
		return err
	}

	return nil
}
