package resource

import (
	"encoding/json"
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
	err, stationData := sendStationRequest()

	v.Variables["stationData"] = stationData

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

func unmarshalRequest(rawJson string) (error, model.StationResponse) {
	var response model.StationResponse
	err := json.Unmarshal([]byte(rawJson), &response)

	// Handle unmarshaling errors
	if err != nil {
		return err, response
	}

	// No error, so response should be ok.
	return nil, response
}

func sendStationRequest() (error, model.StationResponse) {
	response, err := http.Get("https://gbfs.urbansharing.com/oslobysykkel.no/station_information.json")

	if err != nil {
		log.Println(err.Error())
	}

	defer response.Body.Close()

	// Reads a normal string containing the json string.
	bodyStr, err := ioutil.ReadAll(response.Body)
	rawJSON := string(bodyStr)

	// Takes the raw string and unmarshals it into the StationResponse struct.
	err, modelResponse := unmarshalRequest(rawJSON)

	if err != nil {
		log.Println(err.Error())
	}

	return err, modelResponse
}