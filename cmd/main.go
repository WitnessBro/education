package main

import (
	"log"

	_ "github.com/lib/pq"

	"net/http"

	"github.com/WitnessBro/education/internal/app"
	"github.com/WitnessBro/education/internal/config"
	"github.com/WitnessBro/education/pkg/db"
)

func main() {
	config, _ := config.LoadConfig("github.com/WitnessBro/education/internal/config/config.yaml")

	//connect to database
	db, err := db.Connect(config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//create the table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")

	if err != nil {
		log.Fatal(err)
	}
	// Инициализация роутера
	router := app.NewRouter(db)

	// Запуск сервера
	log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", app.JsonContentTypeMiddleware(router)))
}
