package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func getLevels(req events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	res.Body = "hello world"
	res.StatusCode = 200
	return
}

func main() {
	handlers := map[string]interface{}{
		"get_levels": getLevels,
	}
	lambda.Start(handlers[os.Getenv("HANDLER_NAME")])
}
