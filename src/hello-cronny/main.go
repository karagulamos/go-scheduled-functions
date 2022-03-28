package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type requestBody struct {
	NextRun time.Time `json:"next_run"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	requestBody := requestBody{}

	if err := json.Unmarshal([]byte(request.Body), &requestBody); err != nil {
		log.Fatalf("scheduled function error: %s", err.Error())
	}

	log.Printf("Next run at %s\n", requestBody.NextRun.Local())

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusNoContent,
	}, nil
}

func main() {
	lambda.Start(handler)
}
