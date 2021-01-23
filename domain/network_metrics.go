package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/VictoriaMetrics/metrics"
	"github.com/arkits/onhub-web/db"
	"github.com/arkits/onhub-web/models"
	"github.com/arkits/onhub-web/oauth"
	"github.com/spf13/viper"
)

type NetworkMetricsProperties struct {
	mu        sync.Mutex
	IsPolling bool
}

var networkMetricsProps NetworkMetricsProperties

func init() {

	// Initialize the networkMetricsProps
	networkMetricsProps.mu.Lock()
	networkMetricsProps.IsPolling = false
	networkMetricsProps.mu.Unlock()

}

// BeginPollingNetworkMetrics begin the Polling for Network Metrics
func BeginPollingNetworkMetrics() {
	if networkMetricsProps.IsPolling {
		logger.Info("networkMetricsProps is already polling")
	} else {
		go pollForNetworkMetrics()
	}
}

func pollForNetworkMetrics() {

	networkMetricsProps.mu.Lock()
	networkMetricsProps.IsPolling = true
	networkMetricsProps.mu.Unlock()

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

		tx, _ := strconv.ParseFloat(networkMetrics.GroupTraffic.TransmitSpeedBps, 64)
		metrics.GetOrCreateSummary("network_metrics_tx").Update(tx)

		rx, _ := strconv.ParseFloat(networkMetrics.GroupTraffic.ReceiveSpeedBps, 64)
		metrics.GetOrCreateSummary("network_metrics_rx").Update(rx)

		time.Sleep(viper.GetDuration("network_metrics.poll_rate") * time.Millisecond)
	}

}

// GetStoredNetworkMetrics returns stored Network Metrics based on the params
func GetStoredNetworkMetrics() ([]models.ChartNetworkMetrics, error) {

	var storedNetworkMetrics []models.StoredNetworkMetric
	db.Db.Order("created_at desc").Limit(20).Find(&storedNetworkMetrics)

	var toReturn []models.ChartNetworkMetrics

	for _, storedNetworkMetric := range storedNetworkMetrics {

		var chartNetworkMetrics models.ChartNetworkMetrics

		networkMetric, err := db.GenerateNetworkMetricsFromSNM(storedNetworkMetric)
		if err != nil {
			return toReturn, err
		}

		chartNetworkMetrics.Timestamp = storedNetworkMetric.CreatedAt
		chartNetworkMetrics.NetworkMetrics = networkMetric

		toReturn = append(toReturn, chartNetworkMetrics)
	}

	return toReturn, nil

}

// GenerateNetworkMetricsStatus generates a models.NetworkMetricsStatus
func GenerateNetworkMetricsStatus() models.NetworkMetricsStatus {

	var networkMetricsStatus models.NetworkMetricsStatus

	networkMetricsStatus.IsPooling = networkMetricsProps.IsPolling

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
