package httpServer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"superheroe-api/superheroe-golang-api/src/entity"
)

// GetSuperheroes provides all the superheroes
func (h *HttpServer) GetSuperheroes(w http.ResponseWriter, _ *http.Request) {
	superheroList, err := h.ctrl.GetAll(h.ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "GetSuperheroes"}).Error(err.Error())
		HandleCustomError(w, err)
		return
	}

	log.WithFields(log.Fields{"package": "httpServer", "method": "GetSuperheroes"}).Info("ok")
	json.NewEncoder(w).Encode(superheroList)
}

// AddSuperHero let us add a new super hero
func (h *HttpServer) AddSuperHero(w http.ResponseWriter, r *http.Request) {
	// extract body from http.Request context
	ctx := r.Context().Value("hero_object")
	newHero, ok := ctx.(*entity.Character)
	if !ok {
		log.WithFields(log.Fields{"package": "httpServer", "method": "AddSuperHero"}).Error("missing body structure")
		HandleError(w, "Invalid data in request", http.StatusBadRequest)
		return
	}

	_, err := h.ctrl.Add(newHero, h.ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "AddSuperHero"}).Error(err.Error())
		HandleCustomError(w, err)
		return
	}

	log.WithFields(log.Fields{"package": "httpServer", "method": "AddSuperHero"}).Info("ok")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newHero)
}

// GetSuperhero return a single super hero
func (h *HttpServer) GetSuperhero(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hero, err := h.ctrl.GetByID(vars["id"], h.ctx)

	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "GetSuperhero"}).Error(err.Error())
		HandleCustomError(w, err)
		return
	}

	log.WithFields(log.Fields{"package": "httpServer", "method": "GetSuperhero"}).Info("ok")
	json.NewEncoder(w).Encode(hero)
}

// UpdateSuperhero updates a super hero information
func (h *HttpServer) UpdateSuperhero(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var updatedHero entity.Character
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "GetSuperhero"}).Error(err.Error())
		HandleError(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &updatedHero)
	resp, err := h.ctrl.Edit(vars["id"], &updatedHero, h.ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "GetSuperhero"}).Error(err.Error())
		HandleCustomError(w, err)
		return
	}

	log.WithFields(log.Fields{"package": "httpServer", "method": "GetSuperhero"}).Info("ok")
	json.NewEncoder(w).Encode(resp)
}

// DeleteSuperhero deletes a super hero
func (h *HttpServer) DeleteSuperhero(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	resp, err := h.ctrl.Delete(vars["id"], h.ctx)
	if err != nil {
		log.WithFields(log.Fields{"package": "httpServer", "method": "DeleteSuperhero"}).Error(err.Error())
		HandleCustomError(w, err)
		return
	}

	json.NewEncoder(w).Encode(resp)
	log.WithFields(log.Fields{"package": "httpServer", "method": "DeleteSuperhero"}).Info("ok")
}
