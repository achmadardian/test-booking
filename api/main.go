package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	logger := slog.Default()

	if err := godotenv.Load(); err != nil {
		logger.Error("faield to load .env")
		os.Exit(1)
	}
}
