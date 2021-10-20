//Package client provide the posibility to create clients to make http request to externals APIs
package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"superheroe-api/superheroe-golang-api/src/entity"

	log "github.com/sirupsen/logrus"
)

type tradeMadeStruct struct {
	currencies string
	api_key    string
	url        string
}

//NewTradeMade initialice a new trade made controller
func NewTradeMade() Client {
	api_key := os.Getenv("API_KEY")
	currencies := os.Getenv("CURRENCIES")
	if len(strings.TrimSpace(currencies)) == 0 {
		currencies = "EURUSD,GBPUSD"
	}
	if len(strings.TrimSpace(api_key)) == 0 {
		api_key = "12345"
	}

	log.SetFormatter(&log.JSONFormatter{})
	return &tradeMadeStruct{
		currencies: currencies,
		api_key:    api_key,
		url:        "https://marketdata.tradermade.com/api/v1/live?currency=" + currencies + "&api_key=" + api_key,
	}
}

//Get makes an http get request to the public TradeMade API
func (c *tradeMadeStruct) Get() (interface{}, error) {
	res, err := http.Get(c.url)
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
