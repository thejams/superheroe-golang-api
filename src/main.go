package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"superheroe-api/superheroe-golang-api/src/client"
	"superheroe-api/superheroe-golang-api/src/config"
	"superheroe-api/superheroe-golang-api/src/controller"
	"superheroe-api/superheroe-golang-api/src/httpServer"
	"superheroe-api/superheroe-golang-api/src/repository"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	cfg := config.GetAPIConfig()
	ctx := context.Background()
	repo, mongoClient := repository.NewMongoConnection(ctx)
	defer mongoClient.Disconnect(ctx)

	var http_server httpServer.HTTPServer
	var router = mux.NewRouter().StrictSlash(true)
	var credentials handlers.CORSOption
	var methods handlers.CORSOption
	var origins handlers.CORSOption

	{
		client := client.NewTradeMade(cfg.ClientURI)
		ctrl := controller.NewController(repo, client)
		http_server = httpServer.NewHTTPServer(ctrl, ctx)

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
	}

	fmt.Printf("server runing in port %v", cfg.Port)
	fmt.Println()
	log.Fatal(http.ListenAndServe(cfg.Port, handlers.CORS(credentials, methods, origins)(router)))

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
