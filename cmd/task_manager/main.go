package main

import (
	"log/slog"
	"net/http"

	"github.com/WitnessBro/education/internal/app"
	"github.com/WitnessBro/education/internal/config"
	"github.com/WitnessBro/education/internal/migrations"
	"github.com/WitnessBro/education/pkg/db"
)

func main() {
	config, _ := config.LoadConfig("configs/config.yaml")
	// Настраиваем логгер с JSON-форматом
	db, err := db.Connect(config.DatabaseURL)
	if err != nil {
		slog.Error("Can't connect")
	}
	defer db.Close()

	//TODO Graceful shutdown
	// Connect - интерфейс с методами GetConnect и Close

	migrations.DoMigrations(db)
	router := app.NewRouter(db)

	//log.Log(context.Background(), slog.LevelDebug, "Server started on "+config.Port)
	http.ListenAndServe(config.Port, app.JsonContentTypeMiddleware(router))
}
