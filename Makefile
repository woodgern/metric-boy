build:
	docker build . --tag metric-boy

server: build
	docker run -p 8080:8080 -it metric-boy bash

up:
	docker-compose up -d

restart:
	docker-compose down && docker-compose up -d
