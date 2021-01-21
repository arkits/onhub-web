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
	"github.com/gin-gonic/gin"
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

// GetLastStoredNetworkMetrics returns the last stored Network Metric
func GetLastStoredNetworkMetrics() (gin.H, error) {

	var latestSNM models.StoredNetworkMetric
	db.Db.Last(&latestSNM)

	networkMetrics, err := db.GenerateNetworkMetricsFromSNM(latestSNM)
	if err != nil {
		return gin.H{}, err
	}

	return gin.H{
		"created_at":     latestSNM.CreatedAt,
		"id":             latestSNM.ID,
		"network_metric": networkMetrics,
	}, nil

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
