package main

import (
	"database/sql"
	"log"

	"github.com/danecwalker/insight/core/internal/email"
	"github.com/danecwalker/insight/core/internal/magic"
	"github.com/danecwalker/insight/core/internal/users"
	"github.com/danecwalker/insight/core/internal/utils"
)

type Services struct {
	UserService  users.UserService
	MagicService magic.MagicService
	EmailService email.EmailService
}

func SetupDependencies() Services {
	// Setup the database connection
	p := utils.GetEnvString("DB_PATH", ".insight/data.db")
	db, err := sql.Open("sqlite", p)
	if err != nil {
		log.Fatal(err)
	}

	// Create the storage layer
	storage := SetupStorage(db)

	// Create the services
	services := Services{
		UserService:  users.NewUserService(storage.Users),
		MagicService: magic.NewMagicService(storage.Magic),
		EmailService: email.NewEmailService(email.NewSMPTEmailClient("resend", "re_KtAaPF8w_CnbbGGGri17Kv6PrV8LS1EuM", "smtp.resend.com", 465)),
	}

	return services
}
