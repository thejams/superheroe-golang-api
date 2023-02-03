//Package client provide the posibility to create clients to make http request to externals APIs
package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"superheroe-api/superheroe-golang-api/src/entity"
)

type tradeMadeClient struct {
	url    string
	client http.Client
}

var once sync.Once

// NewTradeMadeClient returns a new trade made client
func NewTradeMadeClient(uri string) Client {
	log.SetFormatter(&log.JSONFormatter{})

	c := &tradeMadeClient{}

	once.Do(func() {
		c.url = uri //cfg.TradeMadeClientURI
		c.client = http.Client{
			Timeout: 5 * time.Second,
		}
	})
	return c
}

// SetURL sets the client url
func (c *tradeMadeClient) SetURL(s string) {
	c.url = s
}

//Get makes an http get request to the public TradeMade API
func (c *tradeMadeClient) Get() (interface{}, error) {
	req, err := http.NewRequest(http.MethodGet, c.url, nil)
	if err != nil {
		log.WithFields(log.Fields{"package": "client", "client": "TradeMade", "method": "Get"}).Error(err.Error())
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		log.WithFields(log.Fields{"package": "client", "client": "TradeMade", "method": "Get"}).Error(err.Error())
		return nil, err
	}
	if res.StatusCode == 500 {
		log.WithFields(log.Fields{"package": "client", "client": "TradeMade", "method": "Get"}).Error("TradeMade client error 500")
		return nil, fmt.Errorf("TradeMade client error 500")
	}
	if res.StatusCode == 404 {
		log.WithFields(log.Fields{"package": "client", "client": "TradeMade", "method": "Get"}).Error("TradeMade client error 404")
		return nil, fmt.Errorf("TradeMade client error 404")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.WithFields(log.Fields{"package": "client", "client": "TradeMade", "method": "Get"}).Error(err.Error())
		return nil, err
	}

	tradeMadeObj := entity.TradeMade{}
	err = json.Unmarshal(body, &tradeMadeObj)
	if err != nil {
		log.WithFields(log.Fields{"package": "client", "client": "TradeMade", "method": "Get"}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{"package": "client", "client": "TradeMade", "method": "Get"}).Info("ok")
	return tradeMadeObj, nil
}
