package main

import (
	"github.com/FernandoCagale/serverless-infra/jwt"
	"github.com/FernandoCagale/serverless-task/api/routers"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/urfave/negroni"
)

func main() {
	router := routers.NewRouter()
	task := routers.NewRouter()

	routers.MakeTaskHandlers(task)

	mw := jwtmiddleware.New(*jwt.GetConfigJWT("secret"))

	an := negroni.New(negroni.HandlerFunc(mw.HandlerWithNext), negroni.Wrap(task))
	router.PathPrefix("/v1/api").Handler(an)

	n := negroni.Classic()
	n.UseHandler(router)

	n.Run("127.0.0.1:3000")
}
