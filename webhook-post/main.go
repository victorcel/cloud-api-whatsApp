package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/victorcel/cloud-api-whatsApp/common/lib"
	"log"
	"net/http"
	"os"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	queueURL := os.Getenv("QUEUE_NAME")
	var requestBody interface{}
	err := json.Unmarshal([]byte(request.Body), &requestBody)

	if err != nil {
		log.Fatal(err)
	}

	responseJSON, _ := json.Marshal(requestBody)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sqs.New(sess)

	_, err = svc.SendMessage(&sqs.SendMessageInput{
		MessageGroupId: aws.String("23232332"),
		//MessageAttributes: map[string]*sqs.MessageAttributeValue{
		//	"Title": &sqs.MessageAttributeValue{
		//		DataType:    aws.String("String"),
		//		StringValue: aws.String("The Whistler"),
		//	},
		//	"Author": &sqs.MessageAttributeValue{
		//		DataType:    aws.String("String"),
		//		StringValue: aws.String("John Grisham"),
		//	},
		//	"WeeksOn": &sqs.MessageAttributeValue{
		//		DataType:    aws.String("Number"),
		//		StringValue: aws.String("6"),
		//	},
		//},
		MessageBody: aws.String(string(responseJSON)),
		QueueUrl:    aws.String(queueURL),
	})

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return lib.JSONResponse(http.StatusOK, responseJSON), nil

}

func main() {
	lambda.Start(handler)
}
