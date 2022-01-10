package httpServer

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"superheroe-api/superheroe-golang-api/src/client"
	"superheroe-api/superheroe-golang-api/src/controller"
	"superheroe-api/superheroe-golang-api/src/entity"
	"superheroe-api/superheroe-golang-api/src/repository"
	"superheroe-api/superheroe-golang-api/src/util"

	"github.com/go-playground/validator"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type HttpServer struct {
	ctrl        controller.Controller
	ctx         context.Context
	Credentials handlers.CORSOption
	Methods     handlers.CORSOption
	Origins     handlers.CORSOption
	Router      *mux.Router
}

//NewHTTPServer initialice a new http server
func NewHTTPServer(ctx context.Context, cfg *entity.APPConfig, repo repository.Repository) *HttpServer {
	log.SetFormatter(&log.JSONFormatter{})
	var router = mux.NewRouter().StrictSlash(true)
	var credentials handlers.CORSOption
	var methods handlers.CORSOption
	var origins handlers.CORSOption

	http_server := new(HttpServer)

	{
		client := client.NewTradeMade(cfg.ClientURI)
		ctrl := controller.NewController(repo, client)

		http_server.ctrl = ctrl
		http_server.ctx = ctx

		router.Use(commonMiddleware)
		router.HandleFunc("/health", http_server.Health)
		router.HandleFunc("/superhero", http_server.AddSuperHero).Methods("POST")
		router.HandleFunc("/superhero", http_server.GetSuperheroes).Methods("GET")
		router.HandleFunc("/superhero/{id}", http_server.GetSuperhero).Methods("GET")
		router.HandleFunc("/superhero/{id}", http_server.DeleteSuperhero).Methods("DELETE")
		router.HandleFunc("/superhero/{id}", http_server.UpdateSuperhero).Methods("PUT")
		router.HandleFunc("/client/get", http_server.GetHttpRequest).Methods("GET")

		credentials = handlers.AllowCredentials()
		methods = handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
		origins = handlers.AllowedMethods([]string{"*"})

		http_server.Credentials = credentials
		http_server.Methods = methods
		http_server.Origins = origins
		http_server.Router = router
	}

	return http_server
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

// Health verify if the api is up and running
func (h *HttpServer) Health(res http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(res).Encode(entity.Message{MSG: "status up"})
}

// GetHttpRequest provides all the superheroes
func (h *HttpServer) GetHttpRequest(res http.ResponseWriter, _ *http.Request) {
	httpGetResponse, err := h.ctrl.GetHttpRequest()
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "GetHttpRequest"}).Error(err.Error())
		HandleCustomError(res, err)
		return
	}
	log.WithFields(log.Fields{"package": "httpServer", "method": "GetHttpRequest"}).Info("ok")
	json.NewEncoder(res).Encode(httpGetResponse)
}

// GetSuperheroes provides all the superheroes
func (h *HttpServer) GetSuperheroes(res http.ResponseWriter, _ *http.Request) {
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
func (h *HttpServer) AddSuperHero(res http.ResponseWriter, req *http.Request) {
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
func (h *HttpServer) GetSuperhero(res http.ResponseWriter, req *http.Request) {
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
func (h *HttpServer) UpdateSuperhero(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	var updatedHero entity.Superhero
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "GetSuperhero"}).Error(err.Error())
		HandleError(res, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &updatedHero)
	resp, err := h.ctrl.Edit(vars["id"], &updatedHero, h.ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "GetSuperhero"}).Error(err.Error())
		HandleCustomError(res, err)
		return
	}

	log.WithFields(log.Fields{"package": "httpServer", "method": "GetSuperhero"}).Info("ok")
	json.NewEncoder(res).Encode(resp)
}

// DeleteSuperhero deletes a super hero
func (h *HttpServer) DeleteSuperhero(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	resp, err := h.ctrl.Delete(vars["id"], h.ctx)
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
