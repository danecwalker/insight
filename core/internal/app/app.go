package app

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type App struct {
	dataDir string
	dbFile  string
	mux     *http.ServeMux
}

func NewApp() *App {
	return &App{
		dataDir: ".insight",
		dbFile:  ".insight/data.db",
	}
}

func Json(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func (a *App) Setup() {
	slog.Info("Checking for data directory")
	if _, err := os.Stat(a.dataDir); os.IsNotExist(err) {
		slog.Info("Data directory does not exist. Creating...")
		MkDir(a.dataDir)

		slog.Info("Data directory created")
	} else {
		slog.Info("Data directory already exists")
	}

	slog.Info("Checking for database file")
	if _, err := os.Stat(a.dbFile); os.IsNotExist(err) {
		slog.Info("Database file does not exist. Creating...")
		_, err := os.Create(a.dbFile)
		if err != nil {
			slog.Error("Failed to create database file")
			os.Exit(1)
		}

		slog.Info("Database file created")
	} else {
		slog.Info("Database file already exists")
	}

	slog.Info("Initializing database")

	slog.Info("Creating tables")

	slog.Info("Database initialized")

	slog.Info("Setting up routes")

	a.mux = http.NewServeMux()

	a.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 30)),
		})

		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			slog.Error("Failed to sign token")
			Json(w, http.StatusInternalServerError, map[string]string{
				"error": "Internal server error",
			})
			return
		}

		Json(w, http.StatusOK, map[string]string{
			"token":    tokenString,
			"tokenB64": base64.URLEncoding.EncodeToString([]byte(tokenString)),
		})
	})

	a.mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.URL.Query().Get("token")
		if tokenString == "" {
			Json(w, http.StatusBadRequest, map[string]string{
				"error": "Bad token",
			})
			return
		}

		b, e := base64.URLEncoding.DecodeString(tokenString)
		if e != nil {
			Json(w, http.StatusBadRequest, map[string]string{
				"error": "Bad token",
			})
			return
		}

		token, err := jwt.Parse(string(b), func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				Json(w, http.StatusUnauthorized, map[string]string{
					"error": "Token expired",
				})
				return
			}

			Json(w, http.StatusUnauthorized, map[string]string{
				"error": "Unauthorized",
			})
			return
		}

		if !token.Valid {
			Json(w, http.StatusUnauthorized, map[string]string{
				"error": "Unauthorized",
			})
			return
		}

		Json(w, http.StatusOK, map[string]string{
			"error": "Authorized",
		})
	})

	slog.Info("Routes set up")
}

func (a *App) Run() {
	slog.Info("Running app on port 4565")

	err := http.ListenAndServe(":4565", a.mux)
	if err != nil {
		slog.Error("Failed to run app")
		os.Exit(1)
	}
}
