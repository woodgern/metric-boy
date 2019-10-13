.DEFAULT_GOAL:=help
.PHONY: help

help: ## Show all the available make commands
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

build: ## build the docker image
	docker build . --tag metric-boy

server: build ## run the service in the forground
	docker run -p 8080:8080 -it metric-boy bash

up: build ## run the service in the background
	docker-compose up -d

restart: ## resetart the service
	docker-compose down && docker-compose up -d

rebuild: ## restart the service with rebuild
	docker-compose down && docker build && docker-compose up -d
