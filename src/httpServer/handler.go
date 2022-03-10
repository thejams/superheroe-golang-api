package httpServer

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"superheroe-api/superheroe-golang-api/src/entity"
)

// Health verify if the api is up and running
func (h *HttpServer) Health(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(entity.Message{MSG: "status up"})
}

// GetHttpRequest connects with the TradeMade controller
func (h *HttpServer) GetHttpRequest(w http.ResponseWriter, _ *http.Request) {
	httpGetResponse, err := h.ctrl.GetHttpRequest()
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "GetHttpRequest"}).Error(err.Error())
		HandleCustomError(w, err)
		return
	}
	log.WithFields(log.Fields{"package": "httpServer", "method": "GetHttpRequest"}).Info("ok")
	json.NewEncoder(w).Encode(httpGetResponse)
}
