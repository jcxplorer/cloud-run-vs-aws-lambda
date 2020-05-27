#!/usr/bin/env bash

SERVICE_NAME="sleep"
IMAGE="ADD YOUR IMAGE HERE"
PROJECT="ADD YOUR GCP PROJECT ID HERE"
REGION="us-central1"
URL="ADD YOUR CLOUD RUN SERVICE URL HERE"
REVISION_PREFIX="A" # Change to another value each time you run it, if running multiple times.

for i in {1..20}; do
  gcloud run deploy "$SERVICE_NAME" --image "$IMAGE" --set-env-vars "REVISION=$REVISION_PREFIX$i" --platform managed --project "$PROJECT" --region "$REGION"
  cloud-run-vs-lambda-executor -c 10 -n 10 -report-value=false -cloud-run-url "$URL?ms=5000" >> cold.txt
  cloud-run-vs-lambda-executor -c 10 -n 10 -report-value=false -cloud-run-url "$URL?ms=5000" >> warm.txt
done