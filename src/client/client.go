//Package client provide the posibility to create clients to make http request to externals APIs
package client

type Client interface {
	Get() (interface{}, error)
}
