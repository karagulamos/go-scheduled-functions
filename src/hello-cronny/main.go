package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	currentExecutionTime := time.Now().Format(time.RFC1123)

	fmt.Printf("Function executed at %s\n", currentExecutionTime)

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/plain"},
	}, nil
}

func main() {
	lambda.Start(handler)
}
