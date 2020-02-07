package resource

import (
	"github.com/kaiaverkvist/bikemon/src/config"
	"github.com/kaiaverkvist/bikemon/src/service"
	"github.com/kaiaverkvist/bikemon/src/view"
	"github.com/yarf-framework/yarf"
	"log"
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
	v := view.New(c, config.AppConfig.ViewConfig)
	v.Name = "index"

	// Initializes a new CityBikeService instance and configures it to use the urls from the config.
	cbs := service.New(config.AppConfig.StationsInformationUrl, config.AppConfig.StationsStatusUrl)
	err, cityBikeData := cbs.GetCityBikeData()

	v.Variables["stations"] = cityBikeData

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
