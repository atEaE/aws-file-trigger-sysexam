package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var invokeCount = 0
var myObjects []*s3.Object

func init() {
	svc := s3.New(session.New())
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String("example_bucket"),
	}
	result, _ := svc.ListObjectsV2(input)
	myObjects = result.Contents
}

func InsertHandler(ctx context.Context) (*Response, error) {
	res := &Response{
		ID:      "R00001",
		Author:  "atEaE",
		Message: "Hello World",
	}
	return res, nil
}

func main() {
	lambda.Start(InsertHandler)
}
