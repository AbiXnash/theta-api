package main

import (
	"os"

	"github.com/AbiXnash/theta-api/internals/router"
	"github.com/gookit/slog"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Failed to load env")
	}

	slog.Info("Loaded env successfully")

	slog.Configure(func(logger *slog.SugaredLogger) {
		f := logger.Formatter.(*slog.TextFormatter)
		f.EnableColor = true
	})
	slog.Info("Logger configuration done")
}

func main() {
	slog.Info("Setting up routers")

	r := router.GetRouter()
	r.Run()

	port := os.Getenv("PORT")
	if port != "" {
		slog.Info("Server successfully start at PORT: ", port)
	}
	slog.Info("Server successfully start at default PORT: ", "8080")
}
