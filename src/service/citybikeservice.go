package service

import (
	"github.com/kaiaverkvist/bikemon/src/config"
	"github.com/kaiaverkvist/bikemon/src/model"
	"io/ioutil"
	"net/http"
)

// This holds our CityBikeService that we use to perform operations with.
type CityBikeService struct {
	stationsInformationRestUrl string
	stationsStatusRestUrl string
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
	req, err := http.Get(cbs.stationsInformationRestUrl)

	// Set the identifier header as requested by the API creators.
	// https://oslobysykkel.no/apne-data/sanntid
	req.Header.Set("Client-Identifier", config.AppConfig.AppIdentifier)

	// This reads the request into a string we can use later.
	defer req.Body.Close()
	respBody, err := ioutil.ReadAll(req.Body)
	rawStationInformation := string(respBody)

	req, err = http.Get(cbs.stationsStatusRestUrl)

	// Set the identifier header as requested by the API creators.
	// https://oslobysykkel.no/apne-data/sanntid
	req.Header.Set("Client-Identifier", config.AppConfig.AppIdentifier)

	// Reads into a string we will use later.
	defer req.Body.Close()
	respBody, err = ioutil.ReadAll(req.Body)
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