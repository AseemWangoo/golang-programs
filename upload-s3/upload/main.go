package main

import (
	"encoding/base64"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ResponseToSend events.APIGatewayProxyResponse

func Handler(request events.APIGatewayProxyRequest) (ResponseToSend, error) {

	// Extract the request body
	bodyRequest := &ImageRequestBody{}
	err := json.Unmarshal([]byte(request.Body), &bodyRequest)

	if err != nil {
		return ResponseToSend{Body: err.Error(), StatusCode: 404}, nil
	}

	decoded, err := base64.StdEncoding.DecodeString(bodyRequest.Body)
	if err != nil {
		return ResponseToSend{Body: err.Error(), StatusCode: 404}, nil
	}

	resp := imageUpload(bodyRequest, decoded)

	response, err := json.Marshal(&resp)
	if err != nil {
		return ResponseToSend{Body: err.Error(), StatusCode: 404}, nil
	}

	respToSend := ResponseToSend{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(response),
	}

	//Returning response with AWS Lambda Proxy Response
	return respToSend, nil
}

func main() {
	lambda.Start(Handler)
}
