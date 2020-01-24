package main

import (
	"github.com/kaiaverkvist/bikemon/src/config"
	"github.com/kaiaverkvist/bikemon/src/server"
)

func main() {
	// Keep in mind that this function has to be called before attempting to use config.AppConfig, or else it will hold
	// default values.
	config.LoadConfig()

	// Initialize the http server.
	webservice := server.WebService{}
	webservice.StartWebService()
}
