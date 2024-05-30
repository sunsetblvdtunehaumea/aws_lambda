package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

type MyRequestBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type MyResponseBody struct {
	Message     string        `json:"message"`
	Data        MyRequestBody `json:"data"`
	PathParamID string        `json:"pathParamID"`
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

	// Retrieve the path parameter
	pathParamID := request.PathParameters["id"]

	// Do something with the body and path parameter, for example, log it
	fmt.Printf("Received body: %+v\n", body)
	fmt.Printf("Received path parameter 'id': %s\n", pathParamID)

	// Create a response object
	responseObject := MyResponseBody{
		Message:     "Request processed successfully",
		Data:        body,
		PathParamID: pathParamID,
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
	// Example of creating a stub request with path parameters
	testRequest := events.APIGatewayProxyRequest{
		Body: `{"name": "John Doe", "email": "john.doe@example.com", "age": 30}`,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		PathParameters: map[string]string{
			"id": "12345",
		},
	}

	// Call the handler function with the test request
	response, err := handler(testRequest)
	if err != nil {
		log.Fatalf("Error handling request: %s", err)
	}

	// Print the response
	fmt.Printf("Response: %+v\n", response)
}
