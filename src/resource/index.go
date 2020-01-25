package resource

import (
	"encoding/json"
	"github.com/kaiaverkvist/bikemon/src/config"
	"github.com/kaiaverkvist/bikemon/src/model"
	"github.com/kaiaverkvist/bikemon/src/view"
	"github.com/yarf-framework/yarf"
	"io/ioutil"
	"log"
	"net/http"
)

// IndexResource defines a yarf.Resource which handles HTTP requests.
// The routing defined in server.go calls this resource.
type IndexResource struct {
	yarf.Resource
}

// Get is a handler used by the yarf framework for GET.
func (resource *IndexResource) Get(c *yarf.Context) error {

	// Initializes a new view instance based on the current context.
	// Also selects "index.html" as the template, and leaves an empty variable map-array (v.Variables) that can be used
	// for displaying content.
	v := view.New(c)
	v.Name = "index"

	// Performs a GET request, and attempts to unite the two required requests for the page display.
	err, compositeResponseData := sendRequest()

	v.Variables["stations"] = compositeResponseData

	// Renders the template and checks for errors.
	err = v.Render()
	if err != nil {
		log.Println("Render error: ", err)

		// Display a simple error message, while leaving out the details since it gets logged anyways.
		c.Render("There was an error displaying the page. Please try again.")
		return nil
	}

	return nil
}

// unmarshalStationStatusRequest takes a raw JSON and converts it into the StationResponse model.
func unmarshalRequest(rawStationInformation string, rawStationStatus string) (error, *model.CompositeResponseData) {
	var response model.CompositeResponseData

	err := json.Unmarshal([]byte(rawStationInformation), &response)
	err = json.Unmarshal([]byte(rawStationStatus), &response)

	// Handle unmarshaling errors
	if err != nil {
		return err, nil
	}

	// No error, so response should be ok.
	return nil, &response
}

// sendRequest
func sendRequest() (error, *model.CompositeResponseData) {
	req, err := http.Get("https://gbfs.urbansharing.com/oslobysykkel.no/station_information.json")

	// Set the identifier header as requested by the API creators.
	// https://oslobysykkel.no/apne-data/sanntid
	// This block of code basically sets the header, reads the body content into a string.
	req.Header.Set("Client-Identifier", config.AppConfig.AppIdentifier)
	defer req.Body.Close()
	respBody, err := ioutil.ReadAll(req.Body)
	rawStationInformation := string(respBody)

	req, err = http.Get("https://gbfs.urbansharing.com/oslobysykkel.no/station_status.json")

	// Set the identifier header as requested by the API creators.
	// https://oslobysykkel.no/apne-data/sanntid
	// This block of code basically sets the header, reads the body content into a string.
	req.Header.Set("Client-Identifier", config.AppConfig.AppIdentifier)
	defer req.Body.Close()
	respBody, err = ioutil.ReadAll(req.Body)
	rawStationStatus := string(respBody)

	// We can now create a compositeData model and unmarshal our data into it.
	var compositeData *model.CompositeResponseData
	err, compositeData = unmarshalRequest(rawStationInformation, rawStationStatus)

	// Takes the raw strings and unmarshals it into the StationResponse and the StationStatusResponse structs.
	if err != nil {
		log.Println(err.Error())
		// The error will traverse functions into the actual error handler which lives in the Get function.
		return err, nil
	}


	// No errors? Return the responses.
	return nil, compositeData
}
