//go:build !windows

package app

import (
	"log/slog"
	"os"
)

func MkDir(name string) {
	err := os.Mkdir(name, 0755)
	if err != nil {
		slog.Error("Failed to create data directory")
		os.Exit(1)
	}
}
