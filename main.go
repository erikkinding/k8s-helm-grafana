package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "prom_hello_total_requests",
		Help: "The total number of requests to /hello",
	})
)

func hello(w http.ResponseWriter, req *http.Request) {
	opsProcessed.Inc()
	fmt.Fprintf(w, "hello!\n")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":80", nil)
}
