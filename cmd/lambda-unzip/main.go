package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type S3Object struct {
	Id     string
	region string
	bucket string
	key    string
}

func main() {
	lambda.Start(UnzipHandler)
}
