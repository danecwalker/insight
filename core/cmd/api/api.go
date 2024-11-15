package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/felixge/httpsnoop"
)

var cyan = color.New(color.FgCyan).SprintFunc()
var red = color.New(color.FgHiRed).SprintFunc()
var green = color.New(color.FgHiGreen).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()

type application struct {
	config config
}

type config struct {
	addr     string
	data_dir string
}

func prettyBytes(n int64) string {
	const unit = 1024
	if n < unit {
		return fmt.Sprintf("%d B", n)
	}
	div, exp := int64(unit), 0
	for n := n / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(n)/float64(div), "KMGTPE"[exp])
}

func prettyLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(next, w, r)
		c := fmt.Sprint(m.Code)
		switch c[0] {
		case '2':
			c = green(c)
		case '4':
			c = red(c)
		default:
			c = yellow(c)
		}
		log.Printf("[%s] %s %s took %s - %s", c, r.Method, cyan(r.URL.Path), m.Duration, prettyBytes(m.Written))
	})
}

func (app *application) run(mux *http.ServeMux) error {
	dataDir := app.config.data_dir
	if dataDir == "" {
		dataDir = ".insight"
	}

	// Create the data directory if it doesn't exist
	if err := createDataDir(dataDir); err != nil {
		return err
	}

	// Start the HTTP server
	svr := http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting server on %s", app.config.addr)
	return svr.ListenAndServe()
}

func createDataDir(dataDir string) error {
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		log.Printf("Creating data directory at %s\n", dataDir)
		if err := os.MkdirAll(dataDir, 0600); err != nil {
			return err
		}

		return hideDir(dataDir)
	}

	return nil
}
