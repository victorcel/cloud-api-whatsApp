package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"os"
)

type Webhook struct {
	HubMode        string `json:"hub.mode" validate:"required" `
	HubVerifyToken string `json:"hub.verify_token" validate:"required"`
	HubChallenge   string `json:"hub.challenge" validate:"required"`
}

func deserialize(queryParameters map[string]string) Webhook {
	mode := queryParameters["hub.mode"]
	token := queryParameters["hub.verify_token"]
	challenge := queryParameters["hub.challenge"]

	return Webhook{
		HubMode:        mode,
		HubVerifyToken: token,
		HubChallenge:   challenge,
	}
}
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	validate := validator.New()

	responseWebhook := deserialize(request.QueryStringParameters)

	err := validate.Struct(responseWebhook)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	token := os.Getenv("VERIFY_TOKEN")
	log.Println("token=>", token)
	if responseWebhook.HubMode == "subscribe" && token == responseWebhook.HubVerifyToken {
		return events.APIGatewayProxyResponse{
			Body:       responseWebhook.HubChallenge,
			StatusCode: http.StatusOK,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprint("Error verifying data"),
		StatusCode: http.StatusForbidden,
	}, nil
}

func main() {
	lambda.Start(handler)
}
