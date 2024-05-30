package samplefolder

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

	// Return a successful response
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Request processed successfully",
	}, nil
}

func main() {
	lambda.Start(handler)
}
