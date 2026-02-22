package main

import (
	"os"

	"github.com/AbiXnash/theta-api/internals/db"
	"github.com/AbiXnash/theta-api/internals/router"
	"github.com/gookit/slog"
	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
	configureLogger()
	connectDB()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		slog.Error("Failed to load env")
		return
	}
	slog.Info("Loaded env successfully")
}

func configureLogger() {
	slog.Configure(func(logger *slog.SugaredLogger) {
		f := logger.Formatter.(*slog.TextFormatter)
		f.EnableColor = true
	})
	slog.Info("Logger configured")
}

func connectDB() {
	db.ConnectDB()
}

func main() {
	port := getPort()
	slog.Info("Setting up router")

	r := router.GetRouter()
	slog.Info("Server starting on port: ", port)
	r.Run(":" + port)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
