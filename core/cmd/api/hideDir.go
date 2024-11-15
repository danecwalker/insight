//go:build !windows

package main

func hideDir(dataDir string) error {
	return nil
}
