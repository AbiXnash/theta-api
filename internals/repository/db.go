package db

import (
	"database/sql"
	"os"

	"github.com/gookit/slog"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func ConnectDB() {
	url := os.Getenv("TURSO_DATABASE_URL")

	db, err := sql.Open("libsql", url)
	if err != nil {
		slog.Error("failed to connect to turso")
		os.Exit(1)
	} else {
		slog.Info("Connected to Turso")
	}

	defer db.Close()
}
