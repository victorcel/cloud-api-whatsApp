package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"reflect"
	"testing"
)

func Test_deserialize(t *testing.T) {
	mode := "subscribe"
	verifyToken := "ewewjdeodoewpd"
	challenge := "32983u23u"
	type args struct {
		queryParameters map[string]string
	}
	tests := []struct {
		name string
		args args
		want Webhook
	}{
		{
			name: "success - deserialize",
			args: args{
				map[string]string{
					"hub.mode":         mode,
					"hub.verify_token": verifyToken,
					"hub.challenge":    challenge,
				},
			},
			want: Webhook{
				HubMode:        mode,
				HubVerifyToken: verifyToken,
				HubChallenge:   challenge,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deserialize(tt.args.queryParameters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handler(t *testing.T) {
	token := "12345"
	challenge := "errerre"
	t.Setenv("VERIFY_TOKEN", token)
	type args struct {
		request events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{

		{
			name: "Error validate - handler",
			args: args{
				request: events.APIGatewayProxyRequest{
					QueryStringParameters: map[string]string{
						"hub.mode":          "subscribe",
						"hub.verify_token1": "3232",
						"hub.challenge":     "ewewew",
					},
				},
			},
			want: events.APIGatewayProxyResponse{
				Body:       fmt.Sprint("validate from"),
				StatusCode: http.StatusBadRequest,
			},
			wantErr: false,
		},
		{
			name: "Error verifying data - handler",
			args: args{
				request: events.APIGatewayProxyRequest{
					QueryStringParameters: map[string]string{
						"hub.mode":         "subscribe",
						"hub.verify_token": "3232",
						"hub.challenge":    "ewewew",
					},
				},
			},
			want: events.APIGatewayProxyResponse{
				Body:       fmt.Sprint("Error verifying data"),
				StatusCode: http.StatusForbidden,
			},
			wantErr: false,
		},
		{
			name: "successful - handler",
			args: args{
				request: events.APIGatewayProxyRequest{
					QueryStringParameters: map[string]string{
						"hub.mode":         "subscribe",
						"hub.verify_token": token,
						"hub.challenge":    challenge,
					},
				},
			},
			want: events.APIGatewayProxyResponse{
				Body:       challenge,
				StatusCode: http.StatusOK,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handler(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handler() got = %v, want %v", got, tt.want)
			}
		})
	}
}
