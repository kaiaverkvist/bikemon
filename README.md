# bikemon

A simple web app to display city bike stations with stats.

## How to run
To run the web app, use the following command:
```
go run main.go
```
This should create the configuration file, and start the server.
If any errors are found, those will also be displayed.
Note: The SSL warning is expected in development environments.

### Setting up SSL
In order to use SSL, enable and configure the relevant settings in config.toml

*Note: valid certificate files required.*

```
UseSSL = true
SSLPort = 443
CertFile = "tls/server.crt"
Keyfile = "tls/server.key"
```

### Tests

Use:
```
go test -v ./...
```
This runs the tests in verbose mode, printing out results.

## Technical improvements / concerns
This is a list of things I view as either necessary in a production
scenario, or rather big improvements to the UX.

* Caching the responses from the request to prevent flooding the API.
* Automatic page refresh would make a nice improvement to the UX.
* Map view using OpenStreetView.
* I would probably use a frontend framework with a simple static server
rather than Server Side Rendering (SSR) if this was a production deployment.
This is because the server currently handles all the requests, rather than letting
the clients handle their own data requests. **A good use case for this project would be
a monitor in an office continuesly displaying the status of a station nearby.**
