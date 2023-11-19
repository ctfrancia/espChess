help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  run-tests       to run tests"
	@echo "  run-dev         to run server in development mode"
	@echo "  run-sit         to run server in sit mode"
	@echo "  run-uat         to run server in uat mode"
	@echo "  run-prod        to run server in production mode"

run-tests:
	@echo "Running server..."
	@go run cmd/*.go

run-dev:
	@echo "Running server..."
	@go run cmd/*.go -env=dev

run-sit:
	@echo "Running server..."
	@go run cmd/*.go

run-uat:
	@echo "Running server..."
	@go run cmd/*.go

run-prod:
	@echo "Running server..."
	@go run cmd/*.go
