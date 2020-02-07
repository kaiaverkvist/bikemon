package resource

import (
	"github.com/kaiaverkvist/bikemon/src/config"
	"github.com/kaiaverkvist/bikemon/src/service"
	"github.com/kaiaverkvist/bikemon/src/util"
	"github.com/yarf-framework/yarf"
)

type BikeshareResource struct {
	yarf.Resource
}

// Get is called on GET http requests.
func (resource *BikeshareResource) Get(ctx *yarf.Context) error {

	// Optional parameter to select individual stations.
	stationId := ctx.Param("id")

	cbs := service.New(config.AppConfig.StationsInformationUrl, config.AppConfig.StationsStatusUrl)

	// Let's create a simple shared error message for the err branches.
	errorMessage := util.GenericMessageSuccessResponse{
		Success: false,
		Message: "Unable to get cityBike",
	}

	// If the StationID is unset, we'll display everything
	// - If not, we'll display only one of them.
	if stationId == "" {
		err, data := cbs.GetCityBikeData()
		if err != nil {
			ctx.RenderJSON(errorMessage)

			return nil
		}

		ctx.RenderJSON(data)
		return nil
	} else {
		err, data := cbs.GetDataById(stationId) // Gets the single entry with the ID specified, and if not
		if err != nil {
			errorMessage.Message = err.Error()

			ctx.RenderJSON(errorMessage)
			return nil
		}

		ctx.RenderJSON(data)
		return nil
	}

	// No errors.
	return nil
}
