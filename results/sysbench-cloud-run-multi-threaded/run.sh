#!/usr/bin/env bash

URL_1CPU="ADD YOUR CLOUD RUN 1 CPU SERVICE URL HERE"
URL_2CPU="ADD YOUR CLOUD RUN 2 CPU SERVICE URL HERE"

cloud-run-vs-lambda-executor -count 100 -report-duration=false -cloud-run-url "$URL_1CPU?threads=8" > 1cpu.txt
cloud-run-vs-lambda-executor -count 100 -report-duration=false -cloud-run-url "$URL_2CPU?threads=8" > 2cpu.txt