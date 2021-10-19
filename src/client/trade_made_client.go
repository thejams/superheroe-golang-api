//Package client provide the posibility to create clients to make http request to externals APIs
package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"superheroe-api/superheroe-golang-api/src/entity"

	log "github.com/sirupsen/logrus"
)

type tradeMade struct{}

var (
	currencies = "EURUSD,GBPUSD"
	api_key    = "7N5lZrYSYrdbDwlP85iT"
	url        = "https://marketdata.tradermade.com/api/v1/live?currency=" + currencies + "&api_key=" + api_key
)

//NewTradeMade initialice a new trade made controller
func NewTradeMade() Client {
	log.SetFormatter(&log.JSONFormatter{})
	return &tradeMade{}
}

//Get makes an http get request to the public TradeMade API
func (c *tradeMade) Get() (interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		log.WithFields(log.Fields{"package": "client", "client": "TradeMade", "method": "Get"}).Error(err.Error())
		return nil, err
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
