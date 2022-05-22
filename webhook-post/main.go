package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println(request.Body)
	log.Println(request.QueryStringParameters)
	log.Println(request.Headers)

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprint(request.QueryStringParameters),
		StatusCode: http.StatusForbidden,
	}, nil
}

func main() {
	lambda.Start(handler)
}
