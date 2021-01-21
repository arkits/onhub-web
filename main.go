package main

import (
	"fmt"

	"github.com/arkits/onhub-web/domain"
	"github.com/arkits/onhub-web/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	version string
	logger  = logging.MustGetLogger("main")
)

func init() {

	// Setup the Application wide config through Viper
	SetupConfig()

	// Setup Logger
	domain.SetupLogger()

	// Set Gin's Release Mode
	SetGinReleaseMode()

}

func main() {
	port := viper.GetString("server.port")
	serviceName := viper.GetString("server.name")

	r := gin.New()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Add MetricsMiddleware
	r.Use(domain.MetricsMiddleware())

	// Allow all origins - CORS
	r.Use(cors.Default())

	// Expose the Frontend
	r.Use(static.Serve("/", static.LocalFile("./web/build", false)))
	r.Use(static.Serve(fmt.Sprintf("/%s", serviceName), static.LocalFile("./web/build", false)))

	// Expose Version Endpoint
	r.GET(fmt.Sprintf("/%s/api", serviceName), handlers.VersionHandler)

	// Devices Endpoints
	r.GET(fmt.Sprintf("/%s/api/devices", serviceName), handlers.GetAllDevicesHandler)

	// Network Metrics Endpoints
	r.GET(fmt.Sprintf("/%s/api/network-metrics", serviceName), handlers.GetNetworkMetricsHandler)
	r.GET(fmt.Sprintf("/%s/api/network-metrics/status", serviceName), handlers.GetNetworkMetricsStatusHandler)
	r.POST(fmt.Sprintf("/%s/api/network-metrics/start-polling", serviceName), handlers.KickOffNetworkMetricsPolling)

	// Expose Metrics Endpoint
	r.GET(fmt.Sprintf("/%s/api/metrics", serviceName), handlers.MetricsHandler)

	// Run the Web Server
	logger.Infof("Running on http://localhost:%v/%v", port, serviceName)
	r.Run(":" + port)

}

// SetupConfig -  Setup the application config by reading the config file via Viper
func SetupConfig() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf("Error reading config file! - %s", err)
	}

	// If the version is not set, then initialize it to 0.0.1
	if version == "" {
		version = "0.0.1"
	}

	viper.Set("server.version", version)

}

// SetGinReleaseMode set Gin's release mode based on the Config
func SetGinReleaseMode() {

	releaseMode := viper.GetBool("server.release_mode")
	if releaseMode {
		logger.Debugf("Running in ReleaseMode")
		gin.SetMode(gin.ReleaseMode)
	}
}
