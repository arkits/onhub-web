package models

import "time"

// Station represents a device connected to the WiFI network
// Note: this is a representation of an internal model used by the OnHub API
type Station struct {
	ID                    string   `json:"id"`
	FriendlyName          string   `json:"friendlyName"`
	ConnectionType        string   `json:"connectionType"`
	ApID                  string   `json:"apId"`
	IPAddresses           []string `json:"ipAddresses"`
	WirelessBand          string   `json:"wirelessBand"`
	Connected             bool     `json:"connected"`
	AutomaticFriendlyName string   `json:"automaticFriendlyName"`
	IPAddress             string   `json:"ipAddress"`
}

// GroupTraffic represents the complete the cumulative traffic on the network
// Note: this is a representation of an internal model used by the OnHub API
type GroupTraffic struct {
	TransmitSpeedBps string `json:"transmitSpeedBps,omitempty"`
	ReceiveSpeedBps  string `json:"receiveSpeedBps,omitempty"`
}

// StationMetric is wraps the Station and it's Traffic into one struct
// Note: this is a representation of an internal model used by the OnHub API
type StationMetric struct {
	Station Station `json:"station,omitempty"`
	Traffic Traffic `json:"traffic,omitempty"`
}

// Traffic represents the individual station's traffic metric
// Note: this is a representation of an internal model used by the OnHub API
type Traffic struct {
	TransmitSpeedBps string `json:"transmitSpeedBps"`
	ReceiveSpeedBps  string `json:"receiveSpeedBps"`
}

// NetworkMetricsStatus represents the status of the Network Metrics feature
type NetworkMetricsStatus struct {
	IsPooling                 bool                      `json:"is_polling"`
	StoredNetworkMetricsStats StoredNetworkMetricsStats `json:"stored_network_metrics_stats"`
}

// StoredNetworkMetricsStats encapsulates stats regarding the Network Metrics feature
type StoredNetworkMetricsStats struct {
	Count    int       `json:"count"`
	Earliest time.Time `json:"earliest"`
	Latest   time.Time `json:"latest"`
}
