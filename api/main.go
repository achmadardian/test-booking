package main

import (
	"log/slog"
	"os"

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
}
