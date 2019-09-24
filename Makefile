build:
	sudo docker build . --tag metric-boy

server:
	sudo docker run -p 8080:8080 -it metric-boy bash

serve:
	go run main.go

start-cluster:
	sudo docker-compose up -d

restart-cluster:
	sudo docker-compose down && sudo docker-compose up -d
