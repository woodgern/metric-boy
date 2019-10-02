build:
	docker build . --tag metric-boy

server:
	docker run -p 8080:8080 -it metric-boy bash

serve:
	go run main.go

start-cluster:
	docker-compose up -d

restart-cluster:
	docker-compose down && sudo docker-compose up -d
