package app

import (
	"database/sql"
	"net/http"

	openapiTask "github.com/WitnessBro/education/internal/app/handlers/http"
	"github.com/WitnessBro/education/internal/service"
	"github.com/WitnessBro/education/internal/storage"
	"github.com/go-chi/chi/v5"
)

func NewRouter(db *sql.DB) *chi.Mux {

	userRepo := storage.NewUserStorage(db)
	taskRepo := storage.NewTaskStorage(db)

	r := chi.NewRouter()
	task_manager_service := service.NewTaskManagerService(userRepo, taskRepo)
	openapiTask.HandlerFromMux(task_manager_service, r)
	return r
}

func JsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
