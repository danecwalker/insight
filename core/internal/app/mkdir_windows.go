//go:build windows

package app

import (
	"log/slog"
	"os"
	"syscall"
)

func MkDir(name string) {
	err := os.Mkdir(name, 0755)
	if err != nil {
		slog.Error("Failed to create data directory")
		os.Exit(1)
	}

	fn, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		slog.Error("Failed to create UTF16PtrFromString")
		os.Exit(1)
	}

	err = syscall.SetFileAttributes(fn, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		slog.Error("Failed to hide data directory")
		os.Exit(1)
	}
}
