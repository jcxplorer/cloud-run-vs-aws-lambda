package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Executor interface {
	Execute(job Job) float64
}

type Job struct{}

type Result struct {
	Value    float64
	Duration time.Duration
}

func main() {
	var (
		concurrency    = flag.Int("c", 10, "number of requests to execute concurrently")
		number         = flag.Int("n", 100, "number of requests to execute")
		ramp           = flag.Duration("ramp", 0, "time between new concurrent workers being spawned")
		reportValue    = flag.Bool("report-value", true, "report (print) result value")
		reportDuration = flag.Bool("report-duration", true, "report (print) result value")
		cloudRunURL    = flag.String("cloud-run-url", "", "URL of Cloud Run service")
		lambdaARN      = flag.String("lambda-arn", "", "ARN of Lambda function")
		lambdaPayload  = flag.String("lambda-payload", "{}", "Payload for Lambda function")
	)
	flag.Parse()

	jobs := make(chan Job)
	results := make(chan Result)

	var wg sync.WaitGroup
	wg.Add(*number)
	go func() {
		wg.Wait()
		close(results)
	}()

	var executor Executor
	if *cloudRunURL != "" {
		executor = NewCloudRunExecutor(*cloudRunURL, *concurrency)
	}
	if *lambdaARN != "" {
		lambdaExecutor, err := NewLambdaExecutor(*lambdaARN, *lambdaPayload)
		if err != nil {
			panic(err)
		}
		executor = lambdaExecutor
	}
	if executor == nil {
		fmt.Fprintln(os.Stderr, "Either -lambda-arn or -cloud-run-url must be specified")
		os.Exit(1)
	}

	go func() {
		for i := 0; i < *concurrency; i++ {
			go func() {
				go worker(executor, jobs, results, &wg)
			}()
			if *ramp > 0 {
				time.Sleep(*ramp)
			}
		}
	}()

	go func() {
		for i := 0; i < *number; i++ {
			jobs <- Job{}
		}
		close(jobs)
	}()

	w := csv.NewWriter(os.Stdout)
	for execution := range results {
		record := []string{}

		if *reportDuration {
			record = append(record, strconv.FormatInt(int64(execution.Duration/time.Millisecond), 10))
		}
		if *reportValue {
			record = append(record, strconv.FormatFloat(execution.Value, 'f', -1, 64))
		}

		_ = w.Write(record)
	}
	w.Flush()
}

func worker(executor Executor, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	for job := range jobs {
		t := time.Now()
		result := executor.Execute(job)
		duration := time.Since(t)
		results <- Result{
			Value:    result,
			Duration: duration,
		}
		wg.Done()
	}
}
