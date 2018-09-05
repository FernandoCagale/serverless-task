package routers

import (
	"os"

	"github.com/FernandoCagale/serverless-task/api/handlers"
	"github.com/FernandoCagale/serverless-task/pkg/task"
	"github.com/FernandoCagale/serverless-task/pkg/task/repository"
	"github.com/FernandoCagale/serverless-task/pkg/task/usecase"
	"github.com/gorilla/mux"
)

//MakeTaskHandlers make url handlers
func MakeTaskHandlers(r *mux.Router) {
	service := makeTaskGorm()

	r.Handle("/v1/api/task/{id}", handlers.FindByID(service)).Methods("GET")
	r.Handle("/v1/api/task/{id}", handlers.UpdateByID(service)).Methods("PUT")
	r.Handle("/v1/api/task/{id}", handlers.DeleteByID(service)).Methods("DELETE")
	r.Handle("/v1/api/task", handlers.FindAll(service)).Methods("GET")
	r.Handle("/v1/api/task", handlers.Create(service)).Methods("POST")
}

//makeTaskGorm database postgres
func makeTaskGorm() task.UseCase {
	return usecase.NewService(repository.NewGormRepository(os.Getenv("DATASTORE_URL")))
}

//makeTaskInmemory database memory
func makeTaskInmemory() task.UseCase {
	return usecase.NewService(repository.NewInmemRepository())
}
