package main

import (
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Milliseconds int `json:"ms"`
}

type Response struct {
	Result int `json:"result"`
}

func HandleRequest(event Event) (*Response, error) {
	time.Sleep(time.Duration(event.Milliseconds) * time.Millisecond)
	return &Response{Result: event.Milliseconds}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
