FROM golang:latest

COPY ./src/gochat /go/src/gochat

RUN apt-get -y update \
	&& apt-get -y install postgresql-client \
	&& go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/gochat
RUN dep ensure

RUN echo "alias pq='psql -U \$POSTGRES_USER -h \$POSTGRES_HOST -p \$POSTGRES_PORT -d \$POSTGRES_DB'" >> /root/.bashrc

EXPOSE 80

ENTRYPOINT ["go", "run"]
CMD ["main.go"]
