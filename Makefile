.PHONY: build local deploy

build:
	sam build

local:
	sam local start-api

deploy:
	sam build && sam deploy