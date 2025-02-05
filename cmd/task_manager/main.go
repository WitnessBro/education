package main

import (
	"log/slog"
	"net/http"

	"github.com/WitnessBro/education/internal/app"
	"github.com/WitnessBro/education/internal/config"
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

	//create the table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")

	if err != nil {
		slog.Error("Database already exist")
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS tasks (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")

	if err != nil {
		slog.Error("Database already exist")
	}

	router := app.NewRouter(db)

	//log.Log(context.Background(), slog.LevelDebug, "Server started on "+config.Port)
	http.ListenAndServe(config.Port, app.JsonContentTypeMiddleware(router))
}
