package main

import (
	"log/slog"
	"os"

	"github.com/archey347/go-http-api-template/internal"
)

func main() {
	l := slog.With("service", "go-http-api-template")
	l.Info("Starting")

	// Load Config
	c, err := internal.LoadConfig(os.Getenv("GO_HTTP_API_CONFIG"))
	if err != nil {
		l.Error("Failed to load config", "error", err.Error())
		return
	}

	err = internal.Start(c, l)
	if err != nil {
		l.Error("Failed to start", "error", err.Error())
	}

	l.Info("Stopped")
}
