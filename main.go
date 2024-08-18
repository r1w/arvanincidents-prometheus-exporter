package main

import (
	"cloud_server_status/exporter"
	"cloud_server_status/server"
)

// main starts the Prometheus server and periodically fetches API data
func main() {
	go server.StartPrometheusServer()
	exporter.PeriodicDataFetch()
}
