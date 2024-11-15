package main

import (
	"database/sql"
	"log"

	"github.com/danecwalker/insight/core/internal/store"
	"github.com/danecwalker/insight/core/internal/utils"
	"github.com/resend/resend-go/v2"

	_ "modernc.org/sqlite"
)

func main() {
	p := utils.GetEnvString("DB_PATH", ".insight/data.db")
	db, err := sql.Open("sqlite", p)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	app := application{
		config: config{
			addr: ":8080",
		},
		store:        store.NewSqliteStorage(db),
		resendClient: resend.NewClient(utils.GetEnvString("RESEND_API_KEY", "")),
	}

	mux := app.mount()
	if err := app.run(mux); err != nil {
		log.Fatal(err)
	}
}
