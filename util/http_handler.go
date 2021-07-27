package util

import (
	"encoding/json"
	"net/http"
)

func HandleCustomError(res http.ResponseWriter, customErr error) {
	status, err := DecodeError(customErr)
	res.WriteHeader(status)
	json.NewEncoder(res).Encode(err)
}

func HandleError(res http.ResponseWriter, err string, httpCode int) {
	res.WriteHeader(httpCode)
	json.NewEncoder(res).Encode(err)
}
