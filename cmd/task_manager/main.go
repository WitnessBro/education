package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/WitnessBro/education/internal/app"
	"github.com/WitnessBro/education/internal/config"
	"github.com/WitnessBro/education/pkg/db"
)

func main() {
	config, _ := config.LoadConfig("configs/config.yaml")
	// Настраиваем логгер с JSON-форматом
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo, // Уровень логирования (можно изменить на Debug, Warn, Error)
	}))

	// Устанавливаем глобальный логгер
	slog.SetDefault(logger)

	// Примеры логирования
	logger.Info("Service started", slog.String("env", "development"))
	logger.Debug("This is a debug message", slog.String("module", "main"))
	logger.Error("Something went wrong", slog.String("error", "example error"))
	//log := logger.NewLogger()
	db, err := db.Connect(config.DatabaseURL)
	if err != nil {
		logger.Error("Can't connect")
	}
	defer db.Close()

	//create the table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")

	if err != nil {
		logger.Error("Database already exist")
	}

	router := app.NewRouter(db)

	//log.Log(context.Background(), slog.LevelDebug, "Server started on "+config.Port)

	//TODO: don't know what to do here
	fmt.Println(config.Port)
	fmt.Println(config.LogLevel)
	http.ListenAndServe(config.Port, app.JsonContentTypeMiddleware(router))
}
