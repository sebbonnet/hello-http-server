package main

import (
	"fmt"
	"net/http"
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
	http.ListenAndServe(":8080", nil)
}
