FROM iflavoursbv/go-gin-alpine:latest

RUN mkdir -p /usr/src/app

RUN apk add --update make

WORKDIR /usr/src/app

ADD . /usr/src/app

ENTRYPOINT ["go", "run", "main.go"]
