.PHONY: build

lambda-select-path=./cmd/lambda-select
lambda-insert-path=./cmd/lambda-insert


build:
	go build -o $(lambda-select-path)/lambda-select-bin $(lambda-select-path)
	go build -o $(lambda-insert-path)/lambda-insert-bin $(lambda-insert-path)
