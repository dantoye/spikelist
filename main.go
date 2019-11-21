package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func upload(eq events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	res.Body = fmt.Sprintf("%s:%s", eq.QueryStringParameters["text"], eq.Headers["Nightbot-User"])
	res.StatusCode = 200
	log.Printf("responding with %s", res.Body)
	return
}

func rate(eq events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	res.Body = fmt.Sprintf("%s:%s", eq.QueryStringParameters["text"], eq.Headers["Nightbot-User"])
	res.StatusCode = 200
	log.Printf("responding with %s", res.Body)
	return
}

func register(eq events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	res.Body = fmt.Sprintf("%s:%s", eq.QueryStringParameters["text"], eq.Headers["Nightbot-User"])
	res.StatusCode = 200
	log.Printf("responding with %s", res.Body)
	return
}

func main() {
	handlers := map[string]interface{}{
		"upload":   upload,
		"rate":     rate,
		"register": register,
	}
	lambda.Start(handlers[os.Getenv("HANDLER_NAME")])
}
