.PHONY: test build local delete deploy

## tests
test:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out
build:
	sam build

local:
	sam build && sam local start-api --env-vars env.json

delete:
	sam delete
deploy:
ifeq ("$(token)","")
	@echo "Escribir el token=xxxxxx"
else
	sam build && sam deploy  --parameter-overrides VerifyToken=$(token)
endif


