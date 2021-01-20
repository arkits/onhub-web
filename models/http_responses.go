package models

// GetStationsResponse is the object returned when retriving Stations from the OnHub API
type GetStationsResponse struct {
	Stations []Station `json:"stations,omitempty"`
}

// GetRealTimeMetricsResponse is the object returned when retriving realTimeMetrics from the OnHub API
type GetRealTimeMetricsResponse struct {
	GroupTraffic   GroupTraffic    `json:"groupTraffic,omitempty"`
	StationMetrics []StationMetric `json:"stationMetrics,omitempty"`
}
