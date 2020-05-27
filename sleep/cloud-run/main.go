package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		milliseconds, err := strconv.Atoi(r.URL.Query().Get("ms"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		time.Sleep(time.Duration(milliseconds) * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"result":%d}`, milliseconds)
	})
	http.ListenAndServe(":8080", nil)
}
