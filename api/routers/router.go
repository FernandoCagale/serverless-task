package routers

import (
	"github.com/FernandoCagale/serverless-infra/infra"
	"github.com/gorilla/mux"
)

//NewRouter infra
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.NotFoundHandler = infra.NotFoundHandler()
	router.MethodNotAllowedHandler = infra.NotAllowedHandler()
	return router
}
