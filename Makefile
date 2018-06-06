build-lambda:
		@echo "go build for lambda"
		@cd cmd/ && GOOS=linux GOARCH=amd64 go build -o call lambda.go
		@cd cmd/ && zip handler.zip call

build-cli:
		@echo "go build for cli"
		@GOOS=linux GOARCH=amd64 go build -o ./bin/cli ./cmd/cli.go
		@zip ./bin/cli.zip ./bin/cli
