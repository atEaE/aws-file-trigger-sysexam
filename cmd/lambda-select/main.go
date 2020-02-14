package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

// SelectHandler : Lambdaを介して、RDSのデータをセレクトして、フロント返却します。
func SelectHandler(ctx context.Context) (*Response, error) {
	res := &Response{
		ID:      "R00001",
		Author:  "atEaE",
		Message: "Hello World",
	}
	return res, nil
}

func main() {
	lambda.Start(SelectHandler)
}
