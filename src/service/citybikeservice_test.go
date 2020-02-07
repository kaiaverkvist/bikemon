package service

import (
	"github.com/kaiaverkvist/bikemon/src/config"
	"testing"
)

func TestCityBikeService_GetCityBikeData(t *testing.T) {
	// Let's use the default config.
	config := config.GetDefaultConfig()

	// Instantiate the city bike service and performs the data requests.
	// Also checks for errors related to the http stack.
	service := CityBikeService{config.StationsInformationUrl, config.StationsStatusUrl}
	err, data := service.GetCityBikeData()
	if err != nil {
		t.Error("Unable to get city bike data. Error: ", err)
	}

	// We now check whether LastUpdated is 0.
	// If it is 0, we treat the data as invalid.
	lastUpdated := data.LastUpdated
	if lastUpdated == 0 {
		t.Error("Invalid city bike data, LastUpdated field is nil:", lastUpdated)
	}
}

func TestCityBikeService_GetDataById(t *testing.T) {
	// Let's use the default config.
	config := config.GetDefaultConfig()

	// Instantiate the city bike service and performs the data requests.
	// Also checks for errors related to the http stack.
	service := CityBikeService{config.StationsInformationUrl, config.StationsStatusUrl}
	err, data := service.GetDataById("1009")
	if err != nil {
		t.Error("Unable to get city bike data for test id.", err)
	}

	// We now check whether LastUpdated is 0.
	// If it is 0, we treat the data as invalid.
	lastUpdated := data.LastUpdated
	if lastUpdated == 0 {
		t.Error("Invalid city bike data, LastUpdated field is nil:", lastUpdated)
	}

	if len(data.Data.Stations) != 1 {
		t.Error("Incorrect data filtering. Should only provide a single entry in the Stations struct.", len(data.Data.Stations))
	}
}