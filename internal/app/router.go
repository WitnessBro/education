package app

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", GetUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser(db)).Methods("GET")
	router.HandleFunc("/users", CreateUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser(db)).Methods("DELETE")

	return router
}

func JsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
