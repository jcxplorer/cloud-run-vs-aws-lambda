#!/usr/bin/env bash

REGION="us-east-1"
LAMBDA_ARN="ADD YOUR LAMBDA FUNCTION ARN HERE"
LAMBDA_NAME="cold-start-lambda"
REVISION_PREFIX="A" # Change to another value each time you run it, if running multiple times.

for i in {1..20}; do
  aws lambda update-function-configuration --function-name "$LAMBDA_NAME" --environment "Variables={REVISION=$REVISION_PREFIX$i}" --region "$REGION"
  cloud-run-vs-lambda-executor -c 10 -n 10 -report-value=false -lambda-arn "$LAMBDA_ARN" >> cold.txt
  cloud-run-vs-lambda-executor -c 10 -n 10 -report-value=false -lambda-arn "$LAMBDA_ARN" >> warm.txt
done