package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

// var logger *zap.Logger

// func init() {
// 	logger, _ := zap.NewProduction()
// 	defer logger.Sync()
// }

type Event struct {
	Msg string `json:"name"`
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	eventJson, _ := json.MarshalIndent(req, "", "  ")
	log.Printf("Incoming Request: %s", eventJson)
	log.Printf("REGION: %s", os.Getenv("AWS_REGION"))
	log.Println("ALL ENV VARS:")

	// for _, element := range os.Environ() {
	// 	log.Println(element)
	// }

	// request context
	lc, _ := lambdacontext.FromContext(ctx)
	log.Printf("REQUEST ID: %s", lc.AwsRequestID)
	// global variable
	log.Printf("FUNCTION NAME: %s", lambdacontext.FunctionName)
	// context method
	deadline, _ := ctx.Deadline()
	log.Printf("DEADLINE: %s", deadline)

	log.Printf("Complete Request : %s", fmt.Sprintf("%v", req))

	body := fmt.Sprintf("Hello, your Method is %s and your Path %s!", req.HTTPMethod, req.Path)
	// body := "Hello, your Method is"

	headers := map[string]string{
		"Content-Type": "text/plain",
		"X-My-Header":  "Hello World",
	}

	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: 200,
		Headers:    headers,
	}, nil
}

func main() {
	lambda.Start(handler)
}
