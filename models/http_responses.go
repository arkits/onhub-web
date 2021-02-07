package models

import "time"

// GetStationsResponse is the object returned when retriving Stations from the OnHub API
type GetStationsResponse struct {
	Stations []Station `json:"stations,omitempty"`
}

// GetRealTimeMetricsResponse is the object returned when retriving realTimeMetrics from the OnHub API
type GetRealTimeMetricsResponse struct {
	GroupTraffic   GroupTraffic    `json:"groupTraffic,omitempty"`
	StationMetrics []StationMetric `json:"stationMetrics,omitempty"`
}

// HTTPResponse represents a generic HTTP response
type HTTPResponse struct {
	Message string      `json:"message"`
	Entity  interface{} `json:"entity,omitempty"`
}

// HTTPErrorResponse represents a generic HTTP Error response
type HTTPErrorResponse struct {
	Error       string `json:"error_message"`
	Description string `json:"error_description,omitempty"`
}

// ChartNetworkMetrics represent a RealTimeMetric and it's associated Time
type ChartNetworkMetrics struct {
	Timestamp      time.Time                  `json:"timestamp"`
	NetworkMetrics GetRealTimeMetricsResponse `json:"network_metrics"`
}
