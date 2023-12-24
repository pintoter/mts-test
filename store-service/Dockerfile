FROM golang:1.21.1-alpine

WORKDIR /usr/src/app/

ENV CGO_ENABLED=0

COPY ./store-service ./

RUN apk add --no-cache make && go mod download

ENTRYPOINT go build -o ./.bin/app ./cmd/store-service/main.go && ./.bin/app