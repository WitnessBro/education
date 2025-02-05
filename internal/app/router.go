package app

import (
	"database/sql"
	"net/http"

	openapiTask "github.com/WitnessBro/education/internal/app/handlers/http"
	openapiUser "github.com/WitnessBro/education/internal/app/handlers/user"
	"github.com/WitnessBro/education/internal/service"
	"github.com/WitnessBro/education/internal/storage"
	"github.com/go-chi/chi/v5"
)

func NewRouter(db *sql.DB) *chi.Mux {

	userRepo := storage.NewUserRepository(db)

	r := chi.NewRouter()
	task_manager_service := &service.TaskManagerService{}
	user_service := service.NewUserService(userRepo)
	openapiTask.HandlerFromMux(task_manager_service, r)
	openapiUser.HandlerFromMux(user_service, r)
	return r
}

func JsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
