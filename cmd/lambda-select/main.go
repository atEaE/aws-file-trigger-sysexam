package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// APIResponse is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type APIResponse events.APIGatewayProxyResponse

// SelectHandler : Lambdaを介して、RDSのデータをセレクトして、フロント返却します。
func SelectHandler(ctx context.Context) (APIResponse, error) {
	var buf bytes.Buffer

	entity := &Response{
		ID:      "R00001",
		Author:  "atEaE",
		Message: "Hello World",
	}
	jbuf, err := json.Marshal(entity)

	if err != nil {
		return APIResponse{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, jbuf)

	res := APIResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":         "application/json",
			"X-SysExam-Func-Reply": "select-handler",
		},
	}
	return res, nil
}

func main() {
	lambda.Start(SelectHandler)
}
