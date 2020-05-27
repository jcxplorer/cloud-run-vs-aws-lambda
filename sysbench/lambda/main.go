package main

import (
	"os"
	"path/filepath"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jcxplorer/cloud-run-vs-lambda/sysbench"
)

type Event struct {
	Args []string `json:"args"`
}

type Response struct {
	Result float64 `json:"result"`
}

func HandleRequest(event Event) (*Response, error) {
	sysbenchPath := filepath.Join(os.Getenv("LAMBDA_TASK_ROOT"), "sysbench")

	result, err := sysbench.RunCPUTest(sysbenchPath, event.Args...)
	if err != nil {
		return nil, err
	}

	return &Response{Result: result}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
