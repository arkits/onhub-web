package db

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/arkits/onhub-web/models"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var logger = logging.MustGetLogger("domain")

// Db is the Shared Database instance
var Db *gorm.DB

// InitDatabase initializes the database
func InitDatabase() error {

	logger.Debug("Setting up the DB...")

	ormLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormLogger.Config{
			LogLevel: gormLogger.Silent, // Log level
		},
	)

	switch dbType := viper.GetString("db.type"); dbType {
	case "postgres":

		dbHost := viper.GetString("db.host")
		dbUsername := viper.GetString("db.username")
		dbPassword := viper.GetString("db.password")
		dbName := viper.GetString("db.name")
		dbPort := viper.GetString("db.port")

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUsername, dbPassword, dbName, dbPort)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: ormLogger,
		})

		if err != nil {
			logger.Errorf("Failed to Setup Postgres DB - %v", err)
			return err
		}

		Db = db

	case "sqlite":

		db, err := gorm.Open(sqlite.Open("ohw-data.db"), &gorm.Config{
			Logger: ormLogger,
		})

		if err != nil {
			logger.Errorf("Failed to Setup Sqlite DB - %v", err)
			return err
		}

		Db = db

	default:
		return fmt.Errorf("Invalid db.type - %v", dbType)
	}

	Db.AutoMigrate(&models.StoredNetworkMetric{})

	return nil

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
