package handlers

import (
	"encoding/json"

	"net/http"

	"github.com/WitnessBro/education/internal/models"
	"github.com/WitnessBro/education/internal/storage"
	"github.com/gorilla/mux"
)

// func GetUsers(repo *storage.UserStorage) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		users, err := repo.GetUsers()
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		json.NewEncoder(w).Encode(users)
// 	}
// }

func GetUser(repo *storage.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		user, err := repo.GetUser(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

func CreateUser(repo *storage.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u models.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := repo.CreateUser(&u); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(u)
	}
}

func UpdateUser(repo *storage.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var u models.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := repo.UpdateUser(id, &u); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(u)
	}
}

func DeleteUser(repo *storage.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if err := repo.DeleteUser(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode("User deleted")
	}
}
