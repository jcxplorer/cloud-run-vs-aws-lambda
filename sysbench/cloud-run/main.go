package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jcxplorer/cloud-run-vs-lambda/sysbench"
)

type response struct {
	Result float64 `json:"result"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var args []string
		for k, v := range r.URL.Query() {
			args = append(args, fmt.Sprintf("--%s=%s", k, v[0]))
		}

		result, err := sysbench.RunCPUTest("sysbench", args...)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		enc := json.NewEncoder(w)
		enc.Encode(&response{Result: result})
	})
	http.ListenAndServe(":8080", nil)
}
