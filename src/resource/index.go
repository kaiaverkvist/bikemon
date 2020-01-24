package resource

import (
	"github.com/kaiaverkvist/bikemon/src/view"
	"github.com/yarf-framework/yarf"
	"log"
)

type IndexResource struct {
	yarf.Resource
}

func (resource *IndexResource) Get(c *yarf.Context) error {

	// Initializes a new view instance based on the current context.
	// Also selects "index.html" as the template, and leaves an empty variable map-array (v.Variables) that can be used
	// for displaying content.
	v := view.New(c)
	v.Name = "index"



	// Renders the template and checks for errors.
	err := v.Render()
	if err != nil {
		log.Println("Render error: ", err)

		// Display a simple error message, while leaving out the details since it gets logged anyways.
		c.Render("There was an error displaying the page. Please try again.")
	}

	return nil
}
