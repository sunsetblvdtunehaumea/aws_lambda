package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyRequestBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type MyResponseBody struct {
	Message string        `json:"message"`
	Data    MyRequestBody `json:"data"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body MyRequestBody

	// Unmarshal the request body into the MyRequestBody struct
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       fmt.Sprintf("Error parsing request body: %s", err),
		}, nil
	}

	// Do something with the body, for example, log it
	fmt.Printf("Received body: %+v\n", body)

	// Create a response object
	responseObject := MyResponseBody{
		Message: "Request processed successfully",
		Data:    body,
	}

	// Marshal the response object to JSON
	responseBody, err := json.Marshal(responseObject)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("Error creating response: %s", err),
		}, nil
	}

	// Return the response with JSON body
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBody),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
