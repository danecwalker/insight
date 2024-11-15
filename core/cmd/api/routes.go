package main

import (
	"net/http"

	"github.com/danecwalker/insight/core/internal/auth"
)

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()
	v1 := http.NewServeMux()

	mux.Handle("/api/v1/", prettyLog(http.StripPrefix("/api/v1", v1)))

	services := SetupDependencies()

	// Register the auth handler
	auth.RegisterAuthHandler(v1, services.UserService, services.MagicService, services.EmailService)

	return mux
}
