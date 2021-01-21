package db

import (
	"encoding/json"
	"log"
	"os"

	"github.com/arkits/onhub-web/models"
	"github.com/op/go-logging"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var logger = logging.MustGetLogger("domain")

// Db is the Shared Database instance
var Db *gorm.DB

// InitDatabase initializes the database
func InitDatabase() {

	logger.Debug("Setting up the DB...")

	ormLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormLogger.Config{
			LogLevel: gormLogger.Silent, // Log level
		},
	)

	db, err := gorm.Open(sqlite.Open("ohw-data.db"), &gorm.Config{
		Logger: ormLogger,
	})

	if err != nil {
		logger.Errorf("Failed to Setup DB - %v", err)
	}

	Db = db

	Db.AutoMigrate(&models.StoredNetworkMetric{})

}

// PersistNetworkMetrics writes the Network Metrics to the DB with the key passed as metricsDataKey
func PersistNetworkMetrics(networkMetrics models.GetRealTimeMetricsResponse, metricsDataKey string) error {

	var snm models.StoredNetworkMetric

	snm.ID = metricsDataKey

	b, err := json.Marshal(networkMetrics)
	if err != nil {
		logger.Errorf("Failed to marshal networkMetrics to JSON - %v", err)
		return err
	}
	snm.NetworkMetricJSON = string(b)

	Db.Create(&snm)

	return nil
}

// GenerateStoredNetworkMetricsStats generates stats about stored Network Metrics
func GenerateStoredNetworkMetricsStats() models.StoredNetworkMetricsStats {

	var storedNetworkMetricsStats models.StoredNetworkMetricsStats

	var snmCount int
	Db.Raw("SELECT COUNT(*) FROM stored_network_metrics").Scan(&snmCount)
	storedNetworkMetricsStats.Count = snmCount

	var earliestSNM models.StoredNetworkMetric
	Db.First(&earliestSNM)
	storedNetworkMetricsStats.Earliest = earliestSNM.CreatedAt

	var latestSNM models.StoredNetworkMetric
	Db.Last(&latestSNM)
	storedNetworkMetricsStats.Latest = latestSNM.CreatedAt

	return storedNetworkMetricsStats

}

func GenerateNetworkMetricsFromSNM(snm models.StoredNetworkMetric) (models.GetRealTimeMetricsResponse, error) {
	var networkMetrics models.GetRealTimeMetricsResponse

	err := json.Unmarshal([]byte(snm.NetworkMetricJSON), &networkMetrics)
	if err != nil {
		logger.Errorf("Caught Error in GenerateNetworkMetricsFromSNM - %v", err)
		return networkMetrics, err
	}

	return networkMetrics, nil
}
