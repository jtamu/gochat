FROM golang:latest

COPY ./src/gochat /go/src/gochat

RUN apt-get -y update \
	&& apt-get -y install postgresql-client \
	&& go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/gochat
RUN dep ensure

EXPOSE 80

ENTRYPOINT ["go", "run"]
CMD ["main.go"]
