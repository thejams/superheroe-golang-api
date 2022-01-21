package httpServer

import (
	"encoding/json"
	"net/http"
	"superheroe-api/superheroe-golang-api/src/entity"
	"superheroe-api/superheroe-golang-api/src/util"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func (h *HttpServer) initRouter(r *mux.Router) {
	r.HandleFunc("/health", h.Health)
	r.HandleFunc("/client/get", h.GetHttpRequest).Methods("GET")

	r.HandleFunc("/character", h.GetSuperheroes).Methods("GET")
	r.HandleFunc("/character/{id}", h.GetSuperhero).Methods("GET")
	r.HandleFunc("/character/{id}", h.DeleteSuperhero).Methods("DELETE")
	r.HandleFunc("/character/{id}", h.UpdateSuperhero).Methods("PUT")

	heroRouter := r.PathPrefix("/character").Subrouter()
	heroRouter.Use(validateSuperHeroeFieldsMiddleware)
	heroRouter.HandleFunc("/", h.AddSuperHero).Methods("POST")
}

// Health verify if the api is up and running
func (h *HttpServer) Health(res http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(res).Encode(entity.Message{MSG: "status up"})
}

// GetHttpRequest provides all the superheroes
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

// HandleError handle the custom errors to be returned to the user
func HandleCustomError(w http.ResponseWriter, customErr error) {
	status, err := util.DecodeError(customErr)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

// HandleError handle the errors to be returned to the user
func HandleError(w http.ResponseWriter, err string, httpCode int) {
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(err)
}
