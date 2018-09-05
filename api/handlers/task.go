package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FernandoCagale/serverless-infra/errors"
	"github.com/FernandoCagale/serverless-infra/logger"
	"github.com/FernandoCagale/serverless-infra/render"
	"github.com/FernandoCagale/serverless-task/pkg/entity"
	"github.com/FernandoCagale/serverless-task/pkg/task"
	"github.com/gorilla/mux"
)

func Create(service task.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var task *entity.Task

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&task); err != nil {
			render.ResponseError(w, errors.AddBadRequestError("Invalid request payload"))
			return
		}

		logger.WithFields(logger.Fields{
			"ID":   task.ID,
			"Name": task.Name,
		}).Info("create")

		defer r.Body.Close()

		if err := service.Create(task); err != nil {
			switch err {
			case entity.ErrInvalidPayload:
				render.ResponseError(w, errors.AddBadRequestError(err.Error()))
			default:
				render.ResponseError(w, errors.AddInternalServerError(err.Error()))
			}
			return
		}

		render.Response(w, task, http.StatusCreated)
	})
}

func UpdateByID(service task.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			render.ResponseError(w, errors.AddBadRequestError("Invalid task ID"))
			return
		}

		var task entity.Task
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&task); err != nil {
			render.ResponseError(w, errors.AddBadRequestError("Invalid request payload"))
			return
		}

		logger.WithFields(logger.Fields{
			"ID":   id,
			"Name": task.Name,
		}).Info("updateByID")

		defer r.Body.Close()

		err = service.Update(id, &task)
		if err != nil {
			switch err {
			case entity.ErrNotFound:
				render.ResponseError(w, errors.AddNotFoundError(err.Error()))
			case entity.ErrInvalidPayload:
				render.ResponseError(w, errors.AddBadRequestError(err.Error()))
			default:
				render.ResponseError(w, errors.AddInternalServerError(err.Error()))
			}
			return
		}

		render.Response(w, map[string]string{"updated": "true"}, http.StatusOK)
	})
}

func DeleteByID(service task.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			render.ResponseError(w, errors.AddBadRequestError("Invalid task ID"))
			return
		}

		err = service.Delete(id)
		if err != nil {
			switch err {
			case entity.ErrNotFound:
				render.ResponseError(w, errors.AddNotFoundError(err.Error()))
			default:
				render.ResponseError(w, errors.AddInternalServerError(err.Error()))
			}
			return
		}

		render.Response(w, map[string]string{"deleted": "true"}, http.StatusOK)
	})
}

func FindByID(service task.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			render.ResponseError(w, errors.AddBadRequestError("Invalid task ID"))
			return
		}

		task, err := service.FindByID(id)
		if err != nil {
			switch err {
			case entity.ErrNotFound:
				render.ResponseError(w, errors.AddNotFoundError(err.Error()))
			default:
				render.ResponseError(w, errors.AddInternalServerError(err.Error()))
			}
			return
		}

		render.Response(w, task, http.StatusOK)
	})
}

func FindAll(service task.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tasks, err := service.FindAll()
		if err != nil {
			render.ResponseError(w, errors.AddInternalServerError(err.Error()))
			return
		}

		render.Response(w, tasks, http.StatusOK)
	})
}
