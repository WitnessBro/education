package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/WitnessBro/education/internal/app"
	"github.com/WitnessBro/education/internal/config"
	"github.com/WitnessBro/education/internal/migrations"
	"github.com/WitnessBro/education/pkg/db"
)

func main() {
	//TODO брать конфиг из аргументов cmd, если нет конфига, то не запуститься
	config, _ := config.LoadConfig("configs/config.yaml")
	db, err := db.Connect(config.DatabaseURL)
	//ndb := NewDataBase
	//ndb.GetConnect()
	if err != nil {
		slog.Error("Can't connect")
	}
	defer db.Close()
	//defer ndb.Close()
	//conn := NewStorage
	_, cancel := context.WithCancel(context.Background())

	migrations.DoMigrations(db)
	router := app.NewRouter(db)

	go http.ListenAndServe(config.Port, app.JsonContentTypeMiddleware(router))

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	<-signalCh

	slog.Info("\nGracefully shutting down service...")

	//ndb.Close()

	cancel()

	slog.Info("Shutdown complete.")
	// TODO Connect - интерфейс с методами GetConnect и Close
}
