package domain

import (
	"fmt"
	"sync"
	"time"

	"github.com/VictoriaMetrics/metrics"
	"github.com/arkits/onhub-web/models"
	"github.com/spf13/viper"
)

type MetricsStore struct {
	mu          sync.Mutex
	IsPolling   bool
	MetricsData map[string]models.GetRealTimeMetricsResponse
}

var metricsStore MetricsStore

func init() {

	// Initialize the metricsStore

	initialMetricsData := make(map[string]models.GetRealTimeMetricsResponse)

	metricsStore.mu.Lock()
	metricsStore.IsPolling = false
	metricsStore.MetricsData = initialMetricsData
	metricsStore.mu.Unlock()

}

func BeginPollingNetworkMetrics() {
	if metricsStore.IsPolling {
		logger.Info("metricsStore is already polling")
	} else {
		go PollForNetworkMetrics()
	}
}

func PollForNetworkMetrics() {

	metricsStore.mu.Lock()
	metricsStore.IsPolling = true
	metricsStore.mu.Unlock()

	logger.Info("Starting to poll for NetworkMetrics...")

	for {

		timeStart := time.Now()

		networkMetrics := GetNetworkMetrics()

		metricsStore.mu.Lock()
		metricsDataKey := fmt.Sprintf("%v", timeStart.Unix())
		metricsStore.MetricsData[metricsDataKey] = networkMetrics
		metricsStore.mu.Unlock()

		metrics.GetOrCreateSummary("network_metrics_poll_duration").UpdateDuration(timeStart)

		time.Sleep(viper.GetDuration("network_metrics.poll_rate") * time.Millisecond)
	}

}

// GetStoredNetworkMetrics
func GetStoredNetworkMetrics() map[string]models.GetRealTimeMetricsResponse {
	return metricsStore.MetricsData
}
