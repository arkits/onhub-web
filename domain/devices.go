package domain

import (
	"encoding/json"
	"net/http"

	"github.com/arkits/onhub-web/models"
	"github.com/arkits/onhub-web/oauth"
	"github.com/spf13/viper"
)

var httpClient = &http.Client{}

const (
	FOYER_BASE_URL = "https://googlehomefoyer-pa.googleapis.com/v2"
)

// GetAllDevices returns all devices that are, or used to be connected to the network.
func GetAllDevices() []models.Station {

	token := oauth.GetToken()

	systemID := viper.GetString("system_id")

	requestURL := FOYER_BASE_URL + "/groups/" + systemID + "/stations?prettyPrint=false"

	request, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		logger.Errorf("Caught Error in devices.GetAllDevices - err=%s ", err)
	}
	request.Header.Add("Content-Type", "application/json; charset=utf-8")
	request.Header.Add("Authorization", "Bearer "+token)

	response, err := httpClient.Do(request)
	if err != nil {
		logger.Errorf("Caught Error in devices.GetAllDevices - err=%s ", err)
	}

	defer response.Body.Close()

	var getStationsResponse models.GetStationsResponse
	json.NewDecoder(response.Body).Decode(&getStationsResponse)

	return getStationsResponse.Stations
}
