package server

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

// startPrometheusServer starts the HTTP server to expose metrics
func StartPrometheusServer() {
	log.Println("Starting Prometheus metrics server on :8001")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe("0.0.0.0:8001", nil))
}
