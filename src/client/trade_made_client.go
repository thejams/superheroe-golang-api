//Package client provide the posibility to create clients to make http request to externals APIs
package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"

	"superheroe-api/superheroe-golang-api/src/config"
	"superheroe-api/superheroe-golang-api/src/entity"
)

type TradeMadeClient struct {
	url string
}

var once sync.Once

//InitClient initialice a new trade made controller
func (c *TradeMadeClient) InitClient(cfg *config.APPConfig) {
	log.SetFormatter(&log.JSONFormatter{})

	once.Do(func() {
		c.url = cfg.TradeMadeClientURI
	})
}

// SetURL sets the client url
func (c *TradeMadeClient) SetURL(s string) {
	c.url = s
}

//Get makes an http get request to the public TradeMade API
func (c *TradeMadeClient) Get() (interface{}, error) {
	res, err := http.Get(c.url)
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
