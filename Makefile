.PHONY: build deploy destroy

sqlboilercmd=sqlboiler
lambda-select-path=./cmd/lambda-select
lambda-insert-path=./cmd/lambda-insert

test:
	go test ./...

build: test
	go build -ldflags="-s -w" -o ./bin/lambda-select $(lambda-select-path)
	go build -ldflags="-s -w" -o ./bin/lambda-insert $(lambda-insert-path)

deploy: build
	serverless deploy --verbose

destroy:
	serverless remove --verbose

migrate:
	$(sqlboilercmd) --wipe psql
