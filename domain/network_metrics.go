package domain

import (
	"fmt"
	"sync"
	"time"

	"github.com/VictoriaMetrics/metrics"
	"github.com/arkits/onhub-web/models"
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

		MetricsStore.mu.Lock()
		metricsDataKey := fmt.Sprintf("%v", timeStart.Unix())
		MetricsStore.MetricsData[metricsDataKey] = networkMetrics
		MetricsStore.mu.Unlock()

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

	networkMetricsStatus.StoredNetworkMetricsStats.Count = 0

	networkMetricsStatus.StoredNetworkMetricsStats.Earliest = time.Now()
	networkMetricsStatus.StoredNetworkMetricsStats.Latest = time.Now()

	return networkMetricsStatus

}
