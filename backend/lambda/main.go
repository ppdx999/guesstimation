package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.Path {
	case "/hello":
		return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: "Hello, World!"}, nil
	case "/path1":
		return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: "Response from /path1"}, nil
	case "/path2":
		return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: "Response from /path2"}, nil
	default:
		return events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound, Body: "Not Found"}, nil
	}
}

func main() {
	lambda.Start(handler)
}
