package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func broken(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(500)
	fmt.Fprintf(w, "not working\n")
}

func bad(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(400)
	fmt.Fprintf(w, "can't do this\n")
}

func randomLatency(w http.ResponseWriter, req *http.Request) {
	// between 0 and 1.5s
	latency := time.Duration(rand.Intn(1500)) * time.Millisecond
	time.Sleep(latency)

	w.WriteHeader(200)
	fmt.Fprintf(w, "request took %dms\n", latency.Milliseconds())
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/error", broken)
	http.HandleFunc("/bad", bad)
	http.HandleFunc("/random/latency", randomLatency)
	http.ListenAndServe(":8080", nil)
}
