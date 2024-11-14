package repository

import (
	"database/sql"

	"github.com/danecwalker/insight/core/internal/users/models"
	"github.com/danecwalker/insight/core/internal/utils"
)

type SqliteUserRepository struct {
	db *sql.DB
}

func NewSqliteUserRepository(db *sql.DB) *SqliteUserRepository {
	return &SqliteUserRepository{db: db}
}

func (r *SqliteUserRepository) CreateTable() error {
	_, err := r.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL,
			password TEXT NOT NULL
			)
	`)

	if err != nil {
		return err
	}

	return nil
}

func (r *SqliteUserRepository) InsertUser(user models.User) error {
	_, err := r.db.Exec(`
		INSERT INTO users (email, password) VALUES (?, ?)
	`, user.GetEmail(), utils.HashAndSalt(user.GetPassword()))

	if err != nil {
		return err
	}

	return nil
}

func (r *SqliteUserRepository) GetUserByEmail(email string) (models.User, error) {
	rows, err := r.db.Query(`
		SELECT id, email, password FROM users WHERE email = ? LIMIT 1
	`, email)

	if err != nil {
		return models.User{}, err
	}

	var user_id int
	var user_email string
	var user_password string

	row := rows.Next()

	if !row {
		return models.User{}, nil
	}

	err = rows.Scan(&user_id, &user_email, &user_password)

	if err != nil {
		return models.User{}, err
	}

	user := models.NewUser(user_email, user_password)
	user.SetID(user_id)
	return user, nil
}
