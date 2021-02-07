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

// NetworkMetricsProperties captures the various state about the Network Metrics collection feature
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

		go collectAndPersistNetworkMetrics(timeStart)

		time.Sleep(viper.GetDuration("network_metrics.poll_rate_ms") * time.Millisecond)
	}

}

func collectAndPersistNetworkMetrics(timeStart time.Time) {

	// Collect Network Metrics from API
	networkMetrics, err := GetNetworkMetrics()

	// Log the error and fast-fail
	if err != nil {
		logger.Errorf("Caught Error in GetNetworkMetrics  - %v", err)
		return
	}

	// Generate a key for storage
	metricsDataKey := fmt.Sprintf("%v", timeStart.Unix())

	// Persist to DB
	err = db.PersistNetworkMetrics(networkMetrics, metricsDataKey)
	if err != nil {
		logger.Errorf("Failed to persist networkMetrics - %v", err)
		return
	}

	// Export to Prometheus
	go exportNetworkMetricsToPrometheus(networkMetrics, timeStart)

}

// GetStoredNetworkMetrics returns stored Network Metrics based on the params
func GetStoredNetworkMetrics(limit int, skip int) ([]models.ChartNetworkMetrics, error) {

	var storedNetworkMetrics []models.StoredNetworkMetric

	db.Db.Order("created_at desc").Limit(limit).Find(&storedNetworkMetrics)

	var toReturn []models.ChartNetworkMetrics

	for _, storedNetworkMetric := range storedNetworkMetrics {

		var chartNetworkMetrics models.ChartNetworkMetrics

		// Convert NetworkMetrics from StoredNetworkMetrics
		networkMetrics, err := db.GenerateNetworkMetricsFromSNM(storedNetworkMetric)
		if err != nil {
			return toReturn, err
		}

		// Assign the timestamp
		chartNetworkMetrics.Timestamp = storedNetworkMetric.CreatedAt

		// Filter out un-needed StationMetrics
		var filteredStationMetrics []models.StationMetric
		for _, stationMetric := range networkMetrics.StationMetrics {

			// We only care about Connected Stations
			if isStationConnected(stationMetric.Station) {
				filteredStationMetrics = append(filteredStationMetrics, stationMetric)
			}

		}

		// Assing the networkMetrics
		networkMetrics.StationMetrics = filteredStationMetrics
		chartNetworkMetrics.NetworkMetrics = networkMetrics

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
func GetNetworkMetrics() (models.GetRealTimeMetricsResponse, error) {

	// to return
	var getRealTimeMetricsResponse models.GetRealTimeMetricsResponse

	token := oauth.GetToken()

	systemID := viper.GetString("system_id")

	requestURL := FOYER_BASE_URL + "/groups/" + systemID + "/realtimeMetrics?prettyPrint=false"

	request, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		logger.Errorf("Caught Error in network_metrics.GetNetworkMetrics - err=%s ", err)
		return getRealTimeMetricsResponse, err
	}
	request.Header.Add("Content-Type", "application/json; charset=utf-8")
	request.Header.Add("Authorization", "Bearer "+token)

	response, err := httpClient.Do(request)
	if err != nil {
		logger.Errorf("Caught Error in network_metrics.GetNetworkMetrics - err=%s ", err)
		return getRealTimeMetricsResponse, err
	}

	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(&getRealTimeMetricsResponse)

	return getRealTimeMetricsResponse, nil
}

func exportNetworkMetricsToPrometheus(networkMetrics models.GetRealTimeMetricsResponse, timeStart time.Time) {

	numberOfConnectedStations := 0.0

	metrics.GetOrCreateSummary("network_metrics_poll_duration").UpdateDuration(timeStart)

	tx, _ := strconv.ParseFloat(networkMetrics.GroupTraffic.TransmitSpeedBps, 64)
	metrics.GetOrCreateSummary("network_metrics_tx").Update(tx)

	rx, _ := strconv.ParseFloat(networkMetrics.GroupTraffic.ReceiveSpeedBps, 64)
	metrics.GetOrCreateSummary("network_metrics_rx").Update(rx)

	for _, stationMetrics := range networkMetrics.StationMetrics {

		// Only export metrics for connected stations
		if isStationConnected(stationMetrics.Station) {

			// Station Rx Metrics
			txMetricName := fmt.Sprintf(`station_network_metrics_tx{friendly_name="%v"}`,
				stationMetrics.Station.FriendlyName,
			)
			tx, _ := strconv.ParseFloat(stationMetrics.Traffic.TransmitSpeedBps, 64)
			metrics.GetOrCreateSummary(txMetricName).Update(tx)

			// Station Tx Metrics
			rxMetricName := fmt.Sprintf(`station_network_metrics_rx{friendly_name="%v"}`,
				stationMetrics.Station.FriendlyName,
			)
			rx, _ := strconv.ParseFloat(stationMetrics.Traffic.ReceiveSpeedBps, 64)
			metrics.GetOrCreateSummary(rxMetricName).Update(rx)

			// Count number of connected stations
			numberOfConnectedStations = numberOfConnectedStations + 1

		}

	}

	metrics.GetOrCreateSummary("connected_stations").Update(numberOfConnectedStations)

}

// isStationConnected is a helper function to determine whether a Station is connected to the network
// There are multiple indicators to reach this conclusion, however checking the length of IPAddresses
// is the most accurate.
func isStationConnected(station models.Station) bool {

	if len(station.IPAddresses) >= 1 {
		return true
	}

	return false
}
