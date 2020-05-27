#!/usr/bin/env sh
set -e

rm -rf lambda-package
rm -rf sysbench-lambda.zip

docker rm sysbench-lambda || true
docker build -t sysbench-lambda -f lambda/Dockerfile .
docker create --name sysbench-lambda sysbench-lambda
docker cp sysbench-lambda:/package lambda-package
docker rm sysbench-lambda

cd lambda-package
zip -r ../sysbench-lambda.zip .

cd ..
rm -rf lambda-package