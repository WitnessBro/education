package cmd

import (
	"fmt"
	"net/http"

	"github.com/WitnessBro/education/internal/app"
	"github.com/WitnessBro/education/internal/app/logger"
	"github.com/WitnessBro/education/internal/config"
	"github.com/WitnessBro/education/pkg/db"
)

func Execute() {
	config, _ := config.LoadConfig("internal/config/config.yaml")
	log := logger.NewLogger()
	db, err := db.Connect(config.DatabaseURL)
	if err != nil {
		log.Error("Can't connect")
	}
	defer db.Close()

	//create the table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")

	if err != nil {
		log.Error("Database already exist")
	}

	router := app.NewRouter(db)

	//log.Log(context.Background(), slog.LevelDebug, "Server started on "+config.Port)

	//TODO: don't know what to do here
	fmt.Println(config.Port)
	http.ListenAndServe(config.Port, app.JsonContentTypeMiddleware(router))
}
