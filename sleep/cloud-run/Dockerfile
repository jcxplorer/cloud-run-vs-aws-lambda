FROM golang:1.14.2 AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN go build -o /package/sleep .

FROM scratch
COPY --from=builder /package/sleep /sleep
ENTRYPOINT ["/sleep"]
EXPOSE 8080
