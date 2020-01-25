package server

import (
	"fmt"
	"github.com/kaiaverkvist/bikemon/src/config"
	"github.com/kaiaverkvist/bikemon/src/resource"
	"github.com/kaiaverkvist/bikemon/src/static"
	"log"

	"github.com/yarf-framework/yarf"
)

type WebService struct{}

// StartWebService starts the actual HTTP server.
func (ws WebService) StartWebService() {

	// YARF gets initialized and is essentially a router.
	// https://github.com/yarf-framework/yarf/
	y := yarf.New()

	// Set up the routes.
	ws.addRoutes(y)

	// Start the actual web service.
	// This checks the config to determine whether we should use HTTP or HTTPS.
	if config.AppConfig.Http.UseHttp {
		ws.start(y)
	} else if config.AppConfig.Http.UseSSL {
		ws.startTLS(y)
	} else {
		// If neither HTTP or HTTPS is enabled we'll just notify the user.
		log.Println("❌ | No server enabled.")
		log.Println("❌ | Enable either HTTPS or HTTP to run the server.")
	}
}

// Starts the server in regular HTTP mode.
func (ws WebService) start(y *yarf.Yarf) {
	httpPort := fmt.Sprint(":", config.AppConfig.Http.Port)

	log.Println("❌", "Using unsecured HTTP.")
	log.Println("❌", "DO NOT DEPLOY IN PRODUCTION WITHOUT PROPER SSL CONFIGURATION!")

	log.Println("Running on port:", httpPort)
	y.Start(httpPort)
}

// Starts the server in TLS (SSL) mode.
func (ws WebService) startTLS(y *yarf.Yarf) {
	httpsPort := fmt.Sprint(":", config.AppConfig.Http.SSLPort)

	log.Println("✔️", "Using TLS over HTTPS.")
	log.Println("✔️", "OK.")

	log.Println("Running on port", httpsPort)
	y.StartTLS(httpsPort, config.AppConfig.Http.CertFile, config.AppConfig.Http.Keyfile)
}

// This is where routes are defined.
func (ws WebService) addRoutes(y *yarf.Yarf) {

	// Add the page routes.
	y.Add("/", new(resource.IndexResource))

	// Add a static file server which will serve files over /public.
	static.NewFileServer(y, "public", "/public")
}
