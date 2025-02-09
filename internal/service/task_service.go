package service

import (
	"encoding/json"
	"net/http"

	openapi "github.com/WitnessBro/education/internal/app/handlers/http"
	"github.com/WitnessBro/education/internal/models"
	"github.com/WitnessBro/education/internal/storage"
)

var _ openapi.ServerInterface = (*TaskManagerService)(nil)

type TaskManagerService struct {
	repoUser *storage.UserStorage
	repoTask *storage.TaskStorage
}

func NewTaskManagerService(userStorage *storage.UserStorage, taskStorage *storage.TaskStorage) *TaskManagerService {
	return &TaskManagerService{repoUser: userStorage, repoTask: taskStorage}
}

func (t *TaskManagerService) DeleteUsersId(w http.ResponseWriter, r *http.Request, id int) {
	err := t.repoUser.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetUsers implements http.ServerInterface.
func (t *TaskManagerService) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := t.repoUser.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

// GetUsersId implements http.ServerInterface.
func (t *TaskManagerService) GetUsersId(w http.ResponseWriter, r *http.Request, id int) {
	user, err := t.repoUser.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// PostUsers implements http.ServerInterface.
func (t *TaskManagerService) PostUsers(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := t.repoUser.CreateUser(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// PutUsersId implements http.ServerInterface.
func (t *TaskManagerService) PutUsersId(w http.ResponseWriter, r *http.Request, id int) {
	//vars := chi.URLParam("")
	var u *models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := t.repoUser.UpdateUser(id, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TaskManagerService) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := t.repoTask.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func (t *TaskManagerService) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := t.repoTask.CreateTask(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TaskManagerService) DeleteTask(w http.ResponseWriter, r *http.Request, taskId string) {
	err := t.repoTask.DeleteTask(taskId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TaskManagerService) UpdateTask(w http.ResponseWriter, r *http.Request, taskId string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update task"))
}

func (t *TaskManagerService) GetTaskById(w http.ResponseWriter, r *http.Request, taskId string) {
	tasks, err := t.repoTask.GetTaskById(taskId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)

}
