package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"superheroe-api/superheroe-golang-api/src/controller"
	"superheroe-api/superheroe-golang-api/src/httpServer"
	"superheroe-api/superheroe-golang-api/src/repository"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if len(strings.TrimSpace(port)) == 0 {
		port = ":5000"
	}
	ctx := context.Background()
	repo, mongoClient := repository.NewMongoConnection(ctx)
	defer mongoClient.Disconnect(ctx)
	ctrl := controller.NewController(repo)
	http_server := httpServer.NewHTTPServer(ctrl, ctx)

	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)
	router.HandleFunc("/health", http_server.Health)
	router.HandleFunc("/superhero", http_server.AddSuperHero).Methods("POST")
	router.HandleFunc("/superhero", http_server.GetSuperheroes).Methods("GET")
	router.HandleFunc("/superhero/{id}", http_server.GetSuperhero).Methods("GET")
	router.HandleFunc("/superhero/{id}", http_server.DeleteSuperhero).Methods("DELETE")
	router.HandleFunc("/superhero/{id}", http_server.UpdateSuperhero).Methods("PUT")

	fmt.Printf("server runing in port %v", port)
	fmt.Println()
	log.Fatal(http.ListenAndServe(port, router))

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
