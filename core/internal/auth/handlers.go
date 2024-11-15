package auth

import (
	"net/http"

	"github.com/danecwalker/insight/core/internal/email"
	"github.com/danecwalker/insight/core/internal/magic"
	"github.com/danecwalker/insight/core/internal/users"
	"github.com/danecwalker/insight/core/internal/utils"
)

type authHandlerServices struct {
	User  users.UserService
	Magic magic.MagicService
	Email email.EmailService
}

type authHandler struct {
	services authHandlerServices
}

// NewAuthHandler creates a new auth handler.
func RegisterAuthHandler(mux *http.ServeMux, userService users.UserService, magicService magic.MagicService, emailService email.EmailService) {
	authHandler := &authHandler{
		services: authHandlerServices{
			User:  userService,
			Magic: magicService,
			Email: emailService,
		},
	}

	// Register the auth handler
	mux.HandleFunc("/auth/register", authHandler.Register)
}

// Register implements the registration endpoint.
func (a *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := utils.DecodeJson(r, &req); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	// Create the user
	user, err := a.services.User.CreateUser(req.Email)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	// Create the magic link
	magic, err := a.services.Magic.CreateMagic(user.Email())
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	// Send the magic link
	if err := a.services.Email.SendMagicLink(user.Email(), magic); err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
}
