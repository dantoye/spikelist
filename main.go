package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func upload(eq events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	store := NewStore()
	body, err := HandleUpload(eq.QueryStringParameters["text"], eq.Headers["Nightbot-User"], store)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
	}
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: body}, nil
}

func rate(eq events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	store := NewStore()
	body, err := HandleRate(eq.QueryStringParameters["text"], eq.Headers["Nightbot-User"], store)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
	}
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: body}, nil
}

func register(eq events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	store := NewStore()
	body, err := HandleRegister(eq.QueryStringParameters["text"], eq.Headers["Nightbot-User"], store)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
	}
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: body}, nil
}

func main() {
	handlers := map[string]interface{}{
		"upload":   upload,
		"rate":     rate,
		"register": register,
	}
	lambda.Start(handlers[os.Getenv("HANDLER_NAME")])
}
