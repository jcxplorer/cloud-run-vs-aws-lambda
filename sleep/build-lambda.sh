#!/usr/bin/env sh
set -e

rm -rf lambda-package
rm -rf sleep-lambda.zip

docker rm sleep-lambda || true
docker build -t sleep-lambda -f lambda/Dockerfile .
docker create --name sleep-lambda sleep-lambda
docker cp sleep-lambda:/package lambda-package
docker rm sleep-lambda

cd lambda-package
zip -r ../sleep-lambda.zip .

cd ..
rm -rf lambda-package