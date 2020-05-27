FROM golang:1.14.2 AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN go build -o /package/sleep-lambda ./lambda
