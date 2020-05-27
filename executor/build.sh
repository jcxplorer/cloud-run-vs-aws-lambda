#!/usr/bin/env sh
set -e

docker rm cloud-run-vs-lambda-executor || true
docker build -t cloud-run-vs-lambda-executor  .
docker create --name cloud-run-vs-lambda-executor cloud-run-vs-lambda-executor
docker cp cloud-run-vs-lambda-executor:/package/executor cloud-run-vs-lambda-executor
docker rm cloud-run-vs-lambda-executor