package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	requestProcessed := promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		requestProcessed.Inc()
		fmt.Fprintf(writer, "%s: hello!\n", currentTime())
	})

	http.HandleFunc("/readyz", func(writer http.ResponseWriter, request *http.Request) {
		requestProcessed.Inc()
		fmt.Fprintf(writer, "%s: OK!\n", currentTime())
	})

	http.HandleFunc("/error", func(writer http.ResponseWriter, request *http.Request) {
		requestProcessed.Inc()
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "%s: not working\n", currentTime())
	})

	http.HandleFunc("/bad", func(writer http.ResponseWriter, request *http.Request) {
		requestProcessed.Inc()
		writer.WriteHeader(400)
		fmt.Fprintf(writer, "%s: can't do this\n", currentTime())
	})

	http.HandleFunc("/random/latency", func(writer http.ResponseWriter, request *http.Request) {
		requestProcessed.Inc()

		// between 0 and 2s
		latency := time.Duration(rand.Intn(2000)) * time.Millisecond
		time.Sleep(latency)

		writer.WriteHeader(200)
		fmt.Fprintf(writer, "%s request took %dms\n", currentTime(), latency.Milliseconds())
	})

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

func currentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
