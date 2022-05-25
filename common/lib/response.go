package lib

import (
	//Go packages
	"bytes"
	"encoding/json"
	//Third party packages
	"github.com/aws/aws-lambda-go/events"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// EmptyResponse una respuesta sin body
func EmptyResponse(statusCode int) events.APIGatewayProxyResponse {
	resp := events.APIGatewayProxyResponse{
		StatusCode:      statusCode,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
	}
	return resp
}

// JSONResponse receives a JSON body and a code and returns a Response of type APIGatewayProxyResponse
func JSONResponse(statusCode int, JSONBody []byte) events.APIGatewayProxyResponse {
	var buf bytes.Buffer
	json.HTMLEscape(&buf, JSONBody)
	resp := events.APIGatewayProxyResponse{
		StatusCode:      statusCode,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}
	return resp
}

//JSONError function to build json error
func JSONError(statusCode int, err error) events.APIGatewayProxyResponse {

	body, _ := json.Marshal(map[string]interface{}{
		"message": err.Error(),
	})
	return JSONResponse(statusCode, body)
}

//JSONStringResponse function to response json string
func JSONStringResponse(statusCode int, text string) events.APIGatewayProxyResponse {
	resp := events.APIGatewayProxyResponse{
		StatusCode:      statusCode,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: text,
	}
	return resp
}

//JSONMultipleErrors function to response multiples errors
func JSONMultipleErrors(statusCode int, err []interface{}) events.APIGatewayProxyResponse {

	body, _ := json.Marshal(map[string]interface{}{
		"message": err,
	})

	return JSONResponse(statusCode, body)
}

//JSONSchemaErrorWithExtra function to build a json response with extra params
func JSONSchemaErrorWithExtra(statusCode int, err error, extra interface{}) events.APIGatewayProxyResponse {
	body, _ := json.Marshal(map[string]interface{}{
		"message": err.Error(),
		"extra":   extra,
	})
	return JSONResponse(statusCode, body)
}
