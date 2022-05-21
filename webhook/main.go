package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	queryParameters := request.QueryStringParameters
	mode := queryParameters["hub.mode"] // request.Body //req.query['hub.mode'];
	token := queryParameters["hub.verify_token"]
	challenge := queryParameters["hub.challenge"]
	log.Println(mode, token, challenge)
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", "ok2"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
