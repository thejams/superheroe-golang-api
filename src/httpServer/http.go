package httpServer

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"superheroe-api/superheroe-golang-api/src/controller"
	"superheroe-api/superheroe-golang-api/src/entity"
	"superheroe-api/superheroe-golang-api/src/util"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type HTTPServer interface {
	Health(res http.ResponseWriter, _ *http.Request)
	GetSuperheroes(res http.ResponseWriter, _ *http.Request)
	AddSuperHero(res http.ResponseWriter, req *http.Request)
	GetSuperhero(res http.ResponseWriter, req *http.Request)
	UpdateSuperhero(res http.ResponseWriter, req *http.Request)
	DeleteSuperhero(res http.ResponseWriter, req *http.Request)
}

type httpServer struct {
	ctrl controller.Controller
	ctx  context.Context
}

//NewHTTPServer initialice a new http server
func NewHTTPServer(ctrl controller.Controller, ctx context.Context) HTTPServer {
	log.SetFormatter(&log.JSONFormatter{})
	return &httpServer{
		ctrl: ctrl,
		ctx:  ctx,
	}
}

// Health verify if the api is up and running
func (h *httpServer) Health(res http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(res).Encode(entity.Message{MSG: "status up"})
}

// GetSuperheroes provides all the superheroes
func (h *httpServer) GetSuperheroes(res http.ResponseWriter, _ *http.Request) {
	superheroList, err := h.ctrl.GetAll(h.ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "GetSuperheroes"}).Error(err.Error())
		HandleCustomError(res, err)
		return
	}
	log.WithFields(log.Fields{"package": "httpServer", "method": "GetSuperheroes"}).Info("ok")
	json.NewEncoder(res).Encode(superheroList)
}

// AddSuperHero let us add a new super hero
func (h *httpServer) AddSuperHero(res http.ResponseWriter, req *http.Request) {
	var newHero entity.Superhero
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "AddSuperHero"}).Error(err.Error())
		HandleError(res, "Invalid data in request", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &newHero)
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "AddSuperHero"}).Error(err.Error())
		HandleError(res, "Invalid data in request", http.StatusBadRequest)
		return
	}

	err = ValidateFields(newHero)
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "AddSuperHero"}).Error(err.Error())
		HandleCustomError(res, err)
		return
	}

	_, err = h.ctrl.Add(&newHero, h.ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "AddSuperHero"}).Error(err.Error())
		HandleCustomError(res, err)
		return
	}

	log.WithFields(log.Fields{"package": "httpServer", "method": "AddSuperHero"}).Info("ok")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(newHero)
}

// GetSuperhero return a single super hero
func (h *httpServer) GetSuperhero(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	hero, err := h.ctrl.GetByID(vars["id"], h.ctx)

	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "GetSuperhero"}).Error(err.Error())
		HandleCustomError(res, err)
		return
	}

	log.WithFields(log.Fields{"package": "httpServer", "method": "GetSuperhero"}).Info("ok")
	json.NewEncoder(res).Encode(hero)
}

// UpdateSuperhero updates a super hero information
func (h *httpServer) UpdateSuperhero(res http.ResponseWriter, req *http.Request) {
	var updatedHero entity.Superhero
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "GetSuperhero"}).Error(err.Error())
		HandleError(res, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &updatedHero)
	resp, err := h.ctrl.Edit(&updatedHero)
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "GetSuperhero"}).Error(err.Error())
		HandleCustomError(res, err)
		return
	}

	log.WithFields(log.Fields{"package": "httpServer", "method": "GetSuperhero"}).Info("ok")
	json.NewEncoder(res).Encode(resp)
}

// DeleteSuperhero deletes a super hero
func (h *httpServer) DeleteSuperhero(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	resp, err := h.ctrl.Delete(vars["id"])
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "DeleteSuperhero"}).Error(err.Error())
		HandleCustomError(res, err)
		return
	}

	json.NewEncoder(res).Encode(resp)
	log.WithFields(log.Fields{"package": "httpServer", "method": "DeleteSuperhero"}).Info("ok")
}

// HandleError handle the custom errors to be returned to the user
func HandleCustomError(res http.ResponseWriter, customErr error) {
	status, err := util.DecodeError(customErr)
	res.WriteHeader(status)
	json.NewEncoder(res).Encode(err)
}

// HandleError handle the errors to be returned to the user
func HandleError(res http.ResponseWriter, err string, httpCode int) {
	res.WriteHeader(httpCode)
	json.NewEncoder(res).Encode(err)
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
