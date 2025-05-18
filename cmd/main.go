package main

import (
	"log/slog"
	"os"

	"github.com/mktkhr/backend-go/config"
)

var (
	appConfig *config.Config
)

func main() {
	slog.Info("Starting Server...")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	config.GetConfig()
}
