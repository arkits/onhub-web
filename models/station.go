package models

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

type GroupTraffic struct {
	TransmitSpeedBps string `json:"transmitSpeedBps,omitempty"`
	ReceiveSpeedBps  string `json:"receiveSpeedBps,omitempty"`
}

type StationMetric struct {
	Station Station `json:"station,omitempty"`
	Traffic Traffic `json:"traffic,omitempty"`
}

type Traffic struct {
	TransmitSpeedBps string `json:"transmitSpeedBps"`
	ReceiveSpeedBps  string `json:"receiveSpeedBps"`
}
