package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
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

// updateMetrics updates Prometheus metrics with the latest API data
func UpdateMetrics(apiResponse APIResponse) {
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
