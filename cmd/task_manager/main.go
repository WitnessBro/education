package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/WitnessBro/education/internal/app"
	"github.com/WitnessBro/education/internal/config"
	"github.com/WitnessBro/education/internal/migrations"
	"github.com/WitnessBro/education/pkg/db"
)

func main() {
	config, _ := config.LoadConfig("configs/config.yaml")
	db, err := db.Connect(config.DatabaseURL)
	if err != nil {
		slog.Error("Can't connect")
	}
	defer db.Close()

	_, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	wg.Add(1)

	migrations.DoMigrations(db)
	router := app.NewRouter(db)

	go http.ListenAndServe(config.Port, app.JsonContentTypeMiddleware(router))

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	<-signalCh
	slog.Info("\nGracefully shutting down service...")

	cancel()
	wg.Wait()

	slog.Info("Shutdown complete.")
	// Connect - интерфейс с методами GetConnect и Close
}
