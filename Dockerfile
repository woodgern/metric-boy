FROM golang:1.13.0-alpine

RUN apk add --update make
RUN apk add --update git

WORKDIR /go/src/metric-boy

ADD . /go/src/metric-boy

ENTRYPOINT ["sh", "./docker-entrypoint.sh"]