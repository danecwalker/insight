package contracts

import "github.com/danecwalker/insight/core/internal/users/models"

type UserRepository interface {
	CreateTable() error
	InsertUser(user models.User) error
	GetUserByEmail(email string) (models.User, error)
}
