package httpServer

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"superheroe-api/superheroe-golang-api/src/util"
)

func (h *HttpServer) initRouter(r *mux.Router) {
	r.HandleFunc("/health", h.Health)
	r.HandleFunc("/client/get", h.GetHttpRequest).Methods("GET")

	h.initAuthRouter(r)
	h.initCharacterRouter(r)
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
