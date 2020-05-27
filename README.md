# Cloud Run vs. Lambda Benchmark Suite

This repository contains the code and raw result data to accompany the article Google Cloud Run vs. AWS Lambda: Performance Benchmarks (Part 2).

You can find the result data in the `results/` directory. Below you'll find instructions on running your own benchmarks using the same code.

## Building Lambda Functions and Docker Images

You need Docker installed for this.

There are two Lambda functions and Cloud Run containers that you need to build. There is no automation included here to create the Lambda functions or Cloud Run services.

### sysbench

```
$ cd sysbench
$ ./build-lambda.sh
$ docker build -t <tag> -f cloud-run/Dockerfile .
```

Replace `<tag>` with the image tag you want for Cloud Run. A file `sysbench-lambda.zip` is created and contains the Lambda function you can use to create the Lambda function.

### sleep

```
$ cd sleep
$ ./build-lambda.sh
$ docker build -t <tag> -f cloud-run/Dockerfile cloud-run
```

## Running Benchmarks

This uses Docker to build the executable that you should copy to a Linux instance on the same cloud provider (and region) where you are executing the benchmarks.


```
$ cd executor
$ ./build.sh
```

If the command executed successfully, you'll have a `cloud-run-vs-lambda-executor` static binary for Linux in the `executor` directory.

Each result directory includes a `run.sh` script to reproduce the benchmark. Each script assumes that `cloud-run-vs-lambda-executor` can be found from your `$PATH`. Check and edit the contents of each run script.