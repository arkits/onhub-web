package models

import "time"

type NetworkMetricsStatus struct {
	IsPooling                 bool                      `json:"is_polling"`
	StoredNetworkMetricsStats StoredNetworkMetricsStats `json:"stored_network_metrics_stats"`
}

type StoredNetworkMetricsStats struct {
	Count    int       `json:"count"`
	Earliest time.Time `json:"earliest"`
	Latest   time.Time `json:"latest"`
}
