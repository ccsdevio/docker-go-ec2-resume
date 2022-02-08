# syntax=docker/dockerfile:1

FROM golang:1.17

LABEL maintainer="Chris Scott <chris@ccsdev.io>"

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8081:8081

CMD ["./main"]