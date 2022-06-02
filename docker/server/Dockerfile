FROM golang:1.18-alpine

RUN sed -i 's/https/http/' /etc/apk/repositories
RUN apk update
RUN apk add git && apk add curl && apk add bash

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir -p /server

WORKDIR /server

COPY ./server ./

RUN go build -o tie_scan .

EXPOSE 8081

CMD ./tie_scan