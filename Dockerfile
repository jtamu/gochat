FROM golang:latest

COPY ./src/main.go /go/src/

WORKDIR /go/src/

EXPOSE 80
