package service

import (
	"encoding/json"
	"net/http"

	openapi "github.com/WitnessBro/education/internal/app/handlers/user"
	"github.com/WitnessBro/education/internal/storage"
)

var _ openapi.ServerInterface = (*UserService)(nil)

type UserService struct {
	repo *storage.UserStorage
}

func NewUserService(storage *storage.UserStorage) *UserService {
	return &UserService{repo: storage}
}

func (t *UserService) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := t.repo.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
func (t *UserService) PostUsers(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get task"))
}
func (t *UserService) DeleteUsersId(w http.ResponseWriter, r *http.Request, id int) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get task"))
}
func (t *UserService) GetUsersId(w http.ResponseWriter, r *http.Request, id int) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get task"))
}
func (t *UserService) PutUsersId(w http.ResponseWriter, r *http.Request, id int) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get task"))
}
