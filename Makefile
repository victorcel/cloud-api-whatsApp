.PHONY: build local deploy

build:
	sam build

local:
	sam build && sam local start-api --env-vars env.json

deploy:
	sam build --use-container --container-env-var-file env.json && sam deploy