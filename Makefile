.PHONY: build local deploy prueba

build:
	sam build

local:
	sam build && sam local start-api --env-vars env.json

deploy:
ifeq ("$(token)","")
	@echo "Escribir el token=xxxxxx"
else
	sam build && sam deploy  --parameter-overrides VerifyToken=$(token)
endif


