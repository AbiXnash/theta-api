package main

import (
	"os"

	"github.com/AbiXnash/theta-api/internals/components"
	db "github.com/AbiXnash/theta-api/internals/repository"
	"github.com/AbiXnash/theta-api/internals/router"
	"github.com/AbiXnash/theta-api/internals/service"
	"github.com/gookit/slog"
	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
	configureLogger()

	if err := db.ConnectDB(); err != nil {
		slog.Error("Failed to connect to database")
		os.Exit(1)
	}
}

func main() {
	port := getPort()
	slog.Info("Setting up router")

	r := router.GetRouter(updateServerStatus())
	slog.Info("Server starting on port: ", port)
	r.Run(":" + port)
}

func updateServerStatus() *components.Components {
	comps := components.New()
	comps.DB = db.GetDB()
	comps.Status = service.NewStatusService()

	return comps
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

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
