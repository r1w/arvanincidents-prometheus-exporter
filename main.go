package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "net/http/pprof" // Import pprof for profiling
)

var (
	incidentTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "api_incidents_total",
			Help: "Total number of incidents",
		},
		[]string{"type"},
	)
)

type APIResponse struct {
	Meta      Meta       `json:"meta"`
	Incidents []Incident `json:"incidents"`
}

type Meta struct {
	TotalCount int `json:"total_count"`
}

type Incident struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	StartsAt string `json:"starts_at"`
	EndsAt   string `json:"ends_at"`
	// Omitting ServiceID field from struct
}

func init() {
	// Register the custom metrics with Prometheus's default registry.
	prometheus.MustRegister(incidentTotal)
}

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

	// Sort incidents by StartsAt time in ascending order
	sort.SliceStable(apiResponse.Incidents, func(i, j int) bool {
		// Parse the start times
		startTimeI, errI := time.Parse(time.RFC3339, apiResponse.Incidents[i].StartsAt)
		startTimeJ, errJ := time.Parse(time.RFC3339, apiResponse.Incidents[j].StartsAt)
		if errI != nil || errJ != nil {
			// Handle parse errors; consider entries with parse errors as earlier
			if errI != nil {
				return errJ == nil
			}
			return true // errJ != nil
		}
		return startTimeI.Before(startTimeJ) // Ascending order
	})

	// Print the sorted incidents to the console without IDs and with the new format
	fmt.Println("API Response:")
	fmt.Printf("Meta: %+v\n", apiResponse.Meta)
	fmt.Println("Incidents:")
	for _, incident := range apiResponse.Incidents {
		// Print incident details without ID and in the format: Start Time, End Time, Type, Title
		fmt.Printf("Start Time=%s, End Time=%s, Type=%s, Title=%s\n",
			incident.StartsAt, incident.EndsAt, incident.Type, incident.Title)
	}

	// Update the Prometheus metrics
	incidentTotal.With(prometheus.Labels{"type": "total"}).Set(float64(apiResponse.Meta.TotalCount))

	for _, incident := range apiResponse.Incidents {
		incidentTotal.With(prometheus.Labels{"type": incident.Type}).Inc()
	}

	log.Println("API data fetched and Prometheus metrics updated")
}

func main() {
	// Start the Prometheus metrics server
	http.Handle("/metrics", promhttp.Handler())

	// Start the pprof server for profiling
	go func() {
		log.Println("Starting pprof server on :6060")
		log.Fatal(http.ListenAndServe(":6060", nil))
	}()

	// Start the data fetching goroutine
	go func() {
		for {
			fetchAPIData()
			time.Sleep(5 * time.Minute) // Fetch data every 5 minutes
		}
	}()

	// Start the main HTTP server
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
