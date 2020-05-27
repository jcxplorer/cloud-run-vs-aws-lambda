package main

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type LambdaExecutor struct {
	ARN     string
	Payload string
	lambda  *lambda.Lambda
}

func NewLambdaExecutor(arn, payload string) (*LambdaExecutor, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	svc := lambda.New(sess)
	return &LambdaExecutor{
		ARN:     arn,
		Payload: payload,
		lambda:  svc,
	}, nil
}

func (e *LambdaExecutor) Execute(_ Job) float64 {
	out, err := e.lambda.Invoke(&lambda.InvokeInput{
		FunctionName: aws.String(e.ARN),
		Payload:      []byte(e.Payload),
	})
	if err != nil {
		panic(err)
	}
	var outData lambdaOutput
	if err := json.Unmarshal(out.Payload, &outData); err != nil {
		panic(err)
	}
	return outData.Result
}

type lambdaOutput struct {
	Result float64 `json:"result"`
}
