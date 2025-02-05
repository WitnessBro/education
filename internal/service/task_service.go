package service

import (
	"net/http"

	openapi "github.com/WitnessBro/education/internal/app/handlers/http"
	"github.com/WitnessBro/education/internal/storage"
)

var _ openapi.ServerInterface = (*TaskManagerService)(nil)

type TaskManagerService struct {
	repo *storage.UserStorage
}

func (t *TaskManagerService) GetTasks(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get task"))
}

func (t *TaskManagerService) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create task"))
}

func (t *TaskManagerService) DeleteTask(w http.ResponseWriter, r *http.Request, taskId string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete task"))
}

func (t *TaskManagerService) UpdateTask(w http.ResponseWriter, r *http.Request, taskId string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update task"))
}

func (t *TaskManagerService) GetTaskById(w http.ResponseWriter, r *http.Request, taskId string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get task by id"))
}
