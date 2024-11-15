//go:build windows

package main

import (
	"path/filepath"
	"syscall"
)

func hideDir(dataDir string) error {
	base := filepath.Base(dataDir)
	if len(filepath.Base(dataDir)) > 1 && base[0] == '.' {
		// hide the directory
		name := syscall.StringToUTF16Ptr(dataDir)
		return syscall.SetFileAttributes(name, syscall.FILE_ATTRIBUTE_HIDDEN)
	}

	return nil
}
