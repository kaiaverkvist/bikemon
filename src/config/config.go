package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Config is the main configuration struct, holding sub-structs for each type of configuration.
type Config struct {
	StationsInformationUrl string
	StationsStatusUrl      string

	AppIdentifier string
	Http          Http
	ViewConfig    ViewConfig
}

// ViewConfig represents the configuration used to render views.
type ViewConfig struct {
	BaseURI   string
	Folder    string
	Extension string
	BaseFile  string
	Caching   bool
}

// Http represents settings and variables related to the http server.
type Http struct {
	Port int

	UseHttp  bool
	UseSSL   bool
	SSLPort  int
	CertFile string
	Keyfile  string
}

var (
	AppConfig Config // AppConfig is a global variable used to access the config struct.
)

// LoadConfig does as the name suggests - loads the config.
// If the config.toml file is not found, it will create one using the default config defined in getDefaultConfig().
func LoadConfig() Config {
	configLocation := "./config.toml"
	var conf Config

	if _, err := os.Stat(configLocation); err != nil {
		if os.IsNotExist(err) {
			var buffer bytes.Buffer

			// Create a new encoder.
			encoder := toml.NewEncoder(&buffer)

			// Set the indent level to none.
			encoder.Indent = ""

			// Run the encoding with default config.
			err = encoder.Encode(GetDefaultConfig())

			// Write to filesystem.
			err = ioutil.WriteFile(configLocation, buffer.Bytes(), 777)

			log.Println("Creating config: " + configLocation)
		}
	}

	// Finally deal with errors.
	if _, err := toml.DecodeFile(configLocation, &conf); err != nil {
		fmt.Println("Unable to load config: ", err)
	}

	// Set the instance config.
	AppConfig = conf

	return conf
}

// GetDefaultConfig populates the default configuration struct.
func GetDefaultConfig() Config {

	// Default http settings.
	http := Http{8080, true, false, 443, "tls/server.crt", "tls/server.key"}

	// Default view settings.
	// Defines where a view is rendered from and so on.
	viewConfig := ViewConfig{"/", "templates", "html", "base", true}

	// Place everything into the config struct.
	config := Config{"https://gbfs.urbansharing.com/oslobysykkel.no/station_information.json", "https://gbfs.urbansharing.com/oslobysykkel.no/station_status.json",
		"bikemon-github", http, viewConfig}

	return config
}
