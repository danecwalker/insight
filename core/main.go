package main

import "github.com/danecwalker/insight/core/internal/app"

func main() {
	insight := app.NewApp()

	insight.Setup()
	insight.Run()
}
