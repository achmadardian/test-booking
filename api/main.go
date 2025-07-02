package main

import (
	"log/slog"
	"os"

	"github.com/achmadardian/test-booking-api/config"
	"github.com/achmadardian/test-booking-api/responses"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.Default()
	responses.SetLogger(logger)

	if err := godotenv.Load(); err != nil {
		logger.Error("faield to load .env")
		os.Exit(1)
	}

	DB, err := config.NewDatabase(logger)
	if err != nil {
		logger.Error("failed to connect database",
			slog.Any("error", err),
		)
		os.Exit(1)
	}
	defer DB.Close()
	logger.Info("successfully connect to database")
}
