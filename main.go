package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go func() {
		debugHttp := http.NewServeMux()
		debugHttp.HandleFunc("/health", healthHandler)
		debugHttp.Handle("/metrics", promhttp.Handler())
		log.Println("Starting debug server on :8082")
		log.Fatal(http.ListenAndServe(":8082", debugHttp))
	}()

	http.HandleFunc("/", rootHandler)
	log.Println("Starting server on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func rootHandler(w http.ResponseWriter, _ *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		requestDuration.WithLabelValues("/").Observe(duration)
		requestsTotal.WithLabelValues("/").Inc()
	}()

	fmt.Fprintln(w, "Hello world!")
}
