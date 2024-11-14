package utils

import (
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		slog.Error("Failed to hash password", "error", err)
	}
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		slog.Error("Failed to compare passwords", "error", err)
		return false
	}
	return true
}
