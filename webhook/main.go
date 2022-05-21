package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//mode := request.Body //req.query['hub.mode'];
	fmt.Println(request.Body)
	//let token = req.query['hub.verify_token'];
	//let challenge = req.query['hub.challenge'];
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", "ok"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
