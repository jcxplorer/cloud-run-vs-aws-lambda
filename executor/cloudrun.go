package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CloudRunExecutor struct {
	URL string

	client *http.Client
}

func NewCloudRunExecutor(url string, concurrency int) *CloudRunExecutor {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			MaxIdleConnsPerHost: concurrency,
			MaxConnsPerHost:     concurrency,
			MaxIdleConns:        concurrency,
		},
	}
	return &CloudRunExecutor{
		URL:    url,
		client: client,
	}
}

func (e *CloudRunExecutor) Execute(_ Job) float64 {
	resp, err := e.client.Get(e.URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		panic(fmt.Errorf("unexpected status code %d", resp.StatusCode))
	}

	dec := json.NewDecoder(resp.Body)
	var respData cloudRunResponse
	if err := dec.Decode(&respData); err != nil {
		panic(err)
	}

	return respData.Result
}

type cloudRunResponse struct {
	Result float64 `json:"result"`
}
