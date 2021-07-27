package main

// import of all dependencies
// github.com/gorilla/mux gives us a minimal server and router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"superheroe-api/superheroe-golang-api/controller"
	"superheroe-api/superheroe-golang-api/entity"
	"superheroe-api/superheroe-golang-api/repository"
	"superheroe-api/superheroe-golang-api/util"

	"github.com/gorilla/mux"
)

var ctrl controller.Controller

func health(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(entity.Message{MSG: "status up"})
}

func getSuperheroes(res http.ResponseWriter, req *http.Request) {
	superheroList, _ := ctrl.GetAll()
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(superheroList)
}

func addSuperHeroe(res http.ResponseWriter, req *http.Request) {
	var newHeroe entity.Superheroe
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		util.HandleError(res, "Invalid data in request", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &newHeroe)
	if err != nil {
		util.HandleError(res, "Invalid data in request", http.StatusBadRequest)
		return
	}

	_, err = ctrl.Add(&newHeroe)
	if err != nil {
		util.HandleCustomError(res, err)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(newHeroe)
}

func getSuperheroe(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	hero, err := ctrl.GetByID(vars["id"])

	if err != nil {
		util.HandleCustomError(res, err)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(hero)
}

func updateSuperheroe(res http.ResponseWriter, req *http.Request) {
	var updatedHero entity.Superheroe
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		util.HandleError(res, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &updatedHero)
	resp, err := ctrl.Edit(&updatedHero)
	if err != nil {
		util.HandleCustomError(res, err)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(resp)

}

func deleteSuperheroe(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	resp, err := ctrl.Delete(vars["id"])
	if err != nil {
		util.HandleCustomError(res, err)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(resp)
}

func main() {
	repo := repository.NewRepository()
	ctrl = controller.NewController(repo)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", health)
	router.HandleFunc("/superhero", addSuperHeroe).Methods("POST")
	router.HandleFunc("/superhero", getSuperheroes).Methods("GET")
	router.HandleFunc("/superhero/{id}", getSuperheroe).Methods("GET")
	router.HandleFunc("/superhero/{id}", deleteSuperheroe).Methods("DELETE")
	router.HandleFunc("/superhero/{id}", updateSuperheroe).Methods("PUT")

	fmt.Println(http.ListenAndServe(":5000", router))
}
