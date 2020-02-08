package service

import (
	"errors"
	"github.com/kaiaverkvist/bikemon/src/config"
	"github.com/kaiaverkvist/bikemon/src/model"
	"io/ioutil"
	"net/http"
)

// This holds our CityBikeService that we use to perform operations with.
type CityBikeService struct {
	stationsInformationRestUrl string
	stationsStatusRestUrl      string
}

// New creates a new instance of the CityBikeService.
func New(stationsInformationUrl string, stationsStatusUrl string) *CityBikeService {

	// Set our arguments.
	cityBikeService := &CityBikeService{
		stationsInformationRestUrl: stationsInformationUrl,
		stationsStatusRestUrl:      stationsStatusUrl,
	}

	return cityBikeService
}

// GetCityBikeData externally requests for data
func (cbs *CityBikeService) GetCityBikeData() (error, *model.CompositeResponseData) {

	// Requests information from the stationInformation service.
	response, err := http.Get(cbs.stationsInformationRestUrl)

	// Set the identifier header as requested by the API creators.
	// https://oslobysykkel.no/apne-data/sanntid
	response.Header.Set("Client-Identifier", config.AppConfig.AppIdentifier)

	// This reads the request into a string we can use later.
	defer response.Body.Close()
	respBody, err := ioutil.ReadAll(response.Body)
	rawStationInformation := string(respBody)

	response, err = http.Get(cbs.stationsStatusRestUrl)

	// Set the identifier header as requested by the API creators.
	// https://oslobysykkel.no/apne-data/sanntid
	response.Header.Set("Client-Identifier", config.AppConfig.AppIdentifier)

	// Reads into a string we will use later.
	defer response.Body.Close()
	respBody, err = ioutil.ReadAll(response.Body)
	rawStationStatus := string(respBody)

	// We can now create a compositeData model and unmarshal our data into it.
	var compositeData *model.CompositeResponseData
	err, compositeData = model.UnmarshalIntoCompositeStationData(rawStationInformation, rawStationStatus)

	// Takes the raw strings and unmarshals it into the StationResponse and the StationStatusResponse structs.
	if err != nil {

		// The error will traverse functions into the actual error handler which lives in the Get function.
		return err, nil
	}

	// This should only be hit if we have no errors.
	// Returns *model.CompositeResponseData.
	return nil, compositeData
}

func (cbs *CityBikeService) GetDataById(id string) (error, *model.CompositeResponseData) {

	// Get the base city bike data and catch any errors from the http calls.
	err, data := cbs.GetCityBikeData()
	if err != nil {
		return err, nil
	}

	// We create a new blank struct of the same type and return this instead of modifying the
	// pointer we get from cbs.GetCityBikeData.
	newData := model.CompositeResponseData{
		LastUpdated: data.LastUpdated,
		TTL:         data.TTL,
		Data: struct {
    		Stations []model.Station
		}{},
	}

	// Loop over the stations and add the one with the right ID to the new struct.
	for _, station := range data.Data.Stations {

		// If the station ID matches the ID string, we can remove
		if station.StationID == id {
			newData.Data.Stations = append(newData.Data.Stations, station)
		}
	}

	// Return an error if we don't find a matching station.
	if len(newData.Data.Stations) == 0 {
		return errors.New("unable to retrieve any matching stations"), nil
	}

	return nil, &newData
}
