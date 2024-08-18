package exporter

import (
	"cloud_server_status/metrics"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"
)

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

	var apiResponse metrics.APIResponse
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
	SortIncidentsByStartTime(apiResponse.Incidents)

	// Print sorted incidents to console
	PrintIncidents(apiResponse.Incidents)

	// Update Prometheus metrics
	metrics.UpdateMetrics(apiResponse)

	log.Println("API data fetched and Prometheus metrics updated")
}

// periodicDataFetch periodically fetches API data and updates metrics
func PeriodicDataFetch() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			go fetchAPIData() // Fetch data asynchronously
		}
	}
}

// sortIncidentsByStartTime sorts incidents by their start time
func SortIncidentsByStartTime(incidents []metrics.Incident) {
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
func PrintIncidents(incidents []metrics.Incident) {
	fmt.Println("Sorted Incidents:")
	for _, incident := range incidents {
		fmt.Printf("Start Time=%s, End Time=%s, Type=%s, Title=%s\n",
			incident.StartsAt, incident.EndsAt, incident.Type, incident.Title)
	}
}
