#!/usr/bin/env bash

REGION="us-east-1"
LAMBDA_ARN="ADD YOUR LAMBDA FUNCTION ARN HERE"
LAMBDA_NAME="cold-start-lambda"

for mem in $(seq 128 64 3008); do
  aws lambda update-function-configuration --function-name "$LAMBDA_NAME" --memory-size "$mem" --region "$REGION"
  cloud-run-vs-lambda-executor -count 100 -report-duration=false -lambda-arn "$LAMBDA_ARN" -lambda-payload '{"args": ["--threads=8"]}' > "$mem.txt"
done