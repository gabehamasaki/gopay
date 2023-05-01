.PHONY: default run build test docs clean

# Variables
APP_NAME=gopay

# Tasks 
default: run

run:
	@go run ./cmd/server/main.go
run-with-docs:
	@swag init
	@go run ./cmd/server/main.go
build:
	@go build -o ./dist/$(APP_NAME) ./cmd/server/main.go
run-build:
	@make build
	@./dist/$(APP_NAME)
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -rf ./dist
	@rm -rf ./docs