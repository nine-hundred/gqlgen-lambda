package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"gqlgen-lambda/routers"
)

const defaultPort = ":5000"

func main() {
	lambda.Start(routers.Handler)
}
