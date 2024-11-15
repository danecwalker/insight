//go:build windows
// +build windows

package main

import (
	"log"
	"os"
)

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
