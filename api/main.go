package main

import (
	"github.com/FernandoCagale/serverless-infra/jwt"
	"github.com/FernandoCagale/serverless-task/api/routers"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/negroni"
	"github.com/urfave/negroni"
)

var initialized = false

var negroniLambda *negroniadapter.NegroniAdapter

func handlers(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if !initialized {
		router := routers.NewRouter()
		task := routers.NewRouter()

		routers.MakeTaskHandlers(task)

		mw := jwtmiddleware.New(*jwt.GetConfigJWT("secret"))

		an := negroni.New(negroni.HandlerFunc(mw.HandlerWithNext), negroni.Wrap(task))
		router.PathPrefix("/v1/api").Handler(an)

		n := negroni.Classic()
		n.UseHandler(router)

		negroniLambda = negroniadapter.New(n)
		initialized = true
	}

	return negroniLambda.Proxy(req)
}

func main() {
	lambda.Start(handlers)
}
