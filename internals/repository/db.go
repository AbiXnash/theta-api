package db

import (
	"database/sql"
	"os"

	"github.com/gookit/slog"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB

func ConnectDB() error {
	url := os.Getenv("TURSO_DATABASE_URL")

	var err error
	DB, err = sql.Open("libsql", url)
	if err != nil {
		slog.Error("failed to connect to turso: ", err)
		return err
	}

	if err := DB.Ping(); err != nil {
		slog.Error("failed to ping turso: ", err)
		return err
	}

	slog.Info("Connected to Turso")
	return nil
}

func GetDB() *sql.DB {
	return DB
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
