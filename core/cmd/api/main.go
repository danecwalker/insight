package main

import (
	"database/sql"
	"log"

	"github.com/danecwalker/insight/core/internal/store"
	"github.com/danecwalker/insight/core/internal/utils"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	p := utils.GetEnvString("DB_PATH", "insight.db")
	db, err := sql.Open("sqlite3", "file:"+p+"")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	app := application{
		config: config{
			addr: ":8080",
		},
		store: store.NewSqliteStorage(db),
	}

	mux := app.mount()
	if err := app.run(mux); err != nil {
		log.Fatal(err)
	}
}
