//Package client provide the posibility to create clients to make http request to externals APIs
package client

import "superheroe-api/superheroe-golang-api/src/config"

type Client interface {
	InitClient(*config.APPConfig)
	Get() (interface{}, error)
}
