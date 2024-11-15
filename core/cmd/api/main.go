package main

import (
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	app := application{
		config: config{
			addr: ":8080",
		},
	}

	mux := app.mount()
	if err := app.run(mux); err != nil {
		log.Fatal(err)
	}
}
