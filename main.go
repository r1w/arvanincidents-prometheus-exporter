package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"
)

var (
	// Metric to track the total number of incidents
	incidentTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "api_incidents_total",
			Help: "Total number of incidents, with type as a label",
		},
		[]string{"type"},
	)

	// Metric to track detailed information of incidents
	incidentDetails = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "api_incidents_details",
			Help: "Details of incidents including start and end time",
		},
		[]string{"start_time", "end_time", "type", "title"},
	)
)

// APIResponse represents the structure of the API response
type APIResponse struct {
	Meta      Meta       `json:"meta"`
	Incidents []Incident `json:"incidents"`
}

// Meta contains metadata about the API response
type Meta struct {
	TotalCount int `json:"total_count"`
}

// Incident represents an individual incident
type Incident struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	StartsAt string `json:"starts_at"`
	EndsAt   string `json:"ends_at"`
}

// Initialize and register Prometheus metrics
func init() {
	prometheus.MustRegister(incidentTotal)
	prometheus.MustRegister(incidentDetails)
}

// fetchAPIData fetches data from the API and updates Prometheus metrics
func fetchAPIData() {
	log.Println("Fetching API data...")

	url := "https://statuspal.io/api/v2/status_pages/arvancloud/incidents" // Replace with your actual API URL
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching data from API: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return
	}

	var apiResponse APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return
	}

	// Print Meta information to console
	fmt.Printf("Meta: %+v\n", apiResponse.Meta)

	// Print the number of incidents fetched
	fmt.Printf("Number of incidents fetched: %d\n", len(apiResponse.Incidents))

	// Sort incidents by StartAt time
	sortIncidentsByStartTime(apiResponse.Incidents)

	// Print sorted incidents to console
	printIncidents(apiResponse.Incidents)

	// Update Prometheus metrics
	updateMetrics(apiResponse)

	log.Println("API data fetched and Prometheus metrics updated")
}

// sortIncidentsByStartTime sorts incidents by their start time
func sortIncidentsByStartTime(incidents []Incident) {
	sort.SliceStable(incidents, func(i, j int) bool {
		startTimeI, errI := time.Parse(time.RFC3339, incidents[i].StartsAt)
		startTimeJ, errJ := time.Parse(time.RFC3339, incidents[j].StartsAt)
		if errI != nil || errJ != nil {
			if errI != nil {
				return errJ == nil
			}
			return true
		}
		return startTimeI.Before(startTimeJ)
	})
}

// printIncidents prints the list of incidents to the console
func printIncidents(incidents []Incident) {
	fmt.Println("Sorted Incidents:")
	for _, incident := range incidents {
		fmt.Printf("Start Time=%s, End Time=%s, Type=%s, Title=%s\n",
			incident.StartsAt, incident.EndsAt, incident.Type, incident.Title)
	}
}

// updateMetrics updates Prometheus metrics with the latest API data
func updateMetrics(apiResponse APIResponse) {
	incidentTotal.With(prometheus.Labels{"type": "total"}).Set(float64(apiResponse.Meta.TotalCount))
	incidentDetails.Reset() // Clear previous details

	for _, incident := range apiResponse.Incidents {
		incidentTotal.With(prometheus.Labels{"type": incident.Type}).Inc()
		incidentDetails.With(prometheus.Labels{
			"start_time": incident.StartsAt,
			"end_time":   incident.EndsAt,
			"type":       incident.Type,
			"title":      incident.Title,
		}).Set(1) // Use Set(1) to indicate the presence of this detail
	}
}

// startPrometheusServer starts the HTTP server to expose metrics
func startPrometheusServer() {
	log.Println("Starting Prometheus metrics server on :8000")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// periodicDataFetch periodically fetches API data and updates metrics
func periodicDataFetch() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			go fetchAPIData() // Fetch data asynchronously
		}
	}
}

// main starts the Prometheus server and periodically fetches API data
func main() {
	go startPrometheusServer()
	periodicDataFetch()
}
