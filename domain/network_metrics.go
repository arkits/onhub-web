package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/VictoriaMetrics/metrics"
	"github.com/arkits/onhub-web/db"
	"github.com/arkits/onhub-web/models"
	"github.com/arkits/onhub-web/oauth"
	"github.com/spf13/viper"
)

type InMemoryMetricsStore struct {
	mu          sync.Mutex
	IsPolling   bool
	MetricsData map[string]models.GetRealTimeMetricsResponse
}

var MetricsStore InMemoryMetricsStore

func init() {

	// Initialize the MetricsStore

	initialMetricsData := make(map[string]models.GetRealTimeMetricsResponse)

	MetricsStore.mu.Lock()
	MetricsStore.IsPolling = false
	MetricsStore.MetricsData = initialMetricsData
	MetricsStore.mu.Unlock()

}

// BeginPollingNetworkMetrics begin the Polling for Network Metrics
func BeginPollingNetworkMetrics() {
	if MetricsStore.IsPolling {
		logger.Info("MetricsStore is already polling")
	} else {
		go pollForNetworkMetrics()
	}
}

func pollForNetworkMetrics() {

	MetricsStore.mu.Lock()
	MetricsStore.IsPolling = true
	MetricsStore.mu.Unlock()

	logger.Info("Starting to poll for NetworkMetrics...")

	for {

		timeStart := time.Now()

		networkMetrics := GetNetworkMetrics()

		metricsDataKey := fmt.Sprintf("%v", timeStart.Unix())

		err := db.PersistNetworkMetrics(networkMetrics, metricsDataKey)
		if err != nil {
			logger.Errorf("Failed to persist networkMetrics - %v", err)
		}

		metrics.GetOrCreateSummary("network_metrics_poll_duration").UpdateDuration(timeStart)

		time.Sleep(viper.GetDuration("network_metrics.poll_rate") * time.Millisecond)
	}

}

// GetStoredNetworkMetrics - Temp
func GetStoredNetworkMetrics() map[string]models.GetRealTimeMetricsResponse {

	return MetricsStore.MetricsData

}

// GenerateNetworkMetricsStatus generates a models.NetworkMetricsStatus
func GenerateNetworkMetricsStatus() models.NetworkMetricsStatus {

	var networkMetricsStatus models.NetworkMetricsStatus

	networkMetricsStatus.IsPooling = MetricsStore.IsPolling

	networkMetricsStatus.StoredNetworkMetricsStats = db.GenerateStoredNetworkMetricsStats()

	return networkMetricsStatus

}

// GetNetworkMetrics returns all devices that are, or used to be connected to the network.
func GetNetworkMetrics() models.GetRealTimeMetricsResponse {

	token := oauth.GetToken()

	systemID := viper.GetString("system_id")

	requestURL := FOYER_BASE_URL + "/groups/" + systemID + "/realtimeMetrics?prettyPrint=false"

	request, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		logger.Fatal(err)
	}
	request.Header.Add("Content-Type", "application/json; charset=utf-8")
	request.Header.Add("Authorization", "Bearer "+token)

	response, err := httpClient.Do(request)
	if err != nil {
		logger.Fatal(err)
	}

	defer response.Body.Close()

	var getRealTimeMetricsResponse models.GetRealTimeMetricsResponse
	json.NewDecoder(response.Body).Decode(&getRealTimeMetricsResponse)

	return getRealTimeMetricsResponse
}
