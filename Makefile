.PHONY: build deploy destroy

lambda-select-path=./cmd/lambda-select
lambda-insert-path=./cmd/lambda-insert


build:
	go build -ldflags="-s -w" -o ./bin/lambda-select $(lambda-select-path)
	go build -ldflags="-s -w" -o ./bin/lambda-insert $(lambda-insert-path)

deploy: build
	serverless deploy --verbose

destroy:
	serverless remove --verbose
