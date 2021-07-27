package main

import (
	"log"
	"net/http"
	"superheroe-api/superheroe-golang-api/controller"
	"superheroe-api/superheroe-golang-api/httpServer"
	"superheroe-api/superheroe-golang-api/repository"

	"github.com/gorilla/mux"
)

func main() {
	repo := repository.NewRepository()
	ctrl := controller.NewController(repo)
	http_server := httpServer.NewHTTPServer(ctrl)

	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)
	router.HandleFunc("/health", http_server.Health)
	router.HandleFunc("/superhero", http_server.AddSuperHero).Methods("POST")
	router.HandleFunc("/superhero", http_server.GetSuperheroes).Methods("GET")
	router.HandleFunc("/superhero/{id}", http_server.GetSuperhero).Methods("GET")
	router.HandleFunc("/superhero/{id}", http_server.DeleteSuperhero).Methods("DELETE")
	router.HandleFunc("/superhero/{id}", http_server.UpdateSuperhero).Methods("PUT")

	log.Fatal(http.ListenAndServe(":5000", router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
