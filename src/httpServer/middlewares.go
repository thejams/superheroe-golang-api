package httpServer

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"superheroe-api/superheroe-golang-api/src/entity"
	"superheroe-api/superheroe-golang-api/src/util"

	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
)

//validateMsgFieldsMiddleware Validate the request object fields
func validateSuperHeroeFieldsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hero := new(entity.Character)
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.WithFields(log.Fields{"package": "httpServer", "method": "validateSuperHeroeFieldsMiddleware"}).Error(err.Error())
			HandleError(w, "Invalid data in request", http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(reqBody, &hero)
		if err != nil {
			log.WithFields(log.Fields{"package": "httpServer", "method": "validateSuperHeroeFieldsMiddleware"}).Error(err.Error())
			HandleError(w, "Invalid data in request", http.StatusBadRequest)
			return
		}

		err = ValidateFields(hero)
		if err != nil {
			log.WithFields(log.Fields{"package": "httpServer", "method": "validateSuperHeroeFieldsMiddleware"}).Error(err.Error())
			HandleCustomError(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), "hero_object", hero)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//ValidateFields Validate the request object fields
func ValidateFields(req interface{}) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return &util.BadRequestError{Message: fmt.Sprintf("Los siguientes campos son requeridos: %v", err.Error())}
	}
	return nil
}
