package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"superheroe-api/superheroe-golang-api/src/client"
	"superheroe-api/superheroe-golang-api/src/controller"
	"superheroe-api/superheroe-golang-api/src/httpServer"
	"superheroe-api/superheroe-golang-api/src/repository"

	"github.com/gorilla/handlers"
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

	api_key := os.Getenv("API_KEY")
	currencies := os.Getenv("CURRENCIES")
	if len(strings.TrimSpace(currencies)) == 0 {
		currencies = "EURUSD,GBPUSD"
	}
	if len(strings.TrimSpace(api_key)) == 0 {
		api_key = "12345"
	}
	url := "https://marketdata.tradermade.com/api/v1/live?currency=" + currencies + "&api_key=" + api_key
	client := client.NewTradeMade(url)

	ctrl := controller.NewController(repo, client)
	http_server := httpServer.NewHTTPServer(ctrl, ctx)

	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)
	router.HandleFunc("/health", http_server.Health)
	router.HandleFunc("/superhero", http_server.AddSuperHero).Methods("POST")
	router.HandleFunc("/superhero", http_server.GetSuperheroes).Methods("GET")
	router.HandleFunc("/superhero/{id}", http_server.GetSuperhero).Methods("GET")
	router.HandleFunc("/superhero/{id}", http_server.DeleteSuperhero).Methods("DELETE")
	router.HandleFunc("/superhero/{id}", http_server.UpdateSuperhero).Methods("PUT")
	router.HandleFunc("/client/get", http_server.GetHttpRequest).Methods("GET")

	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedMethods([]string{"*"})

	fmt.Printf("server runing in port %v", port)
	fmt.Println()
	log.Fatal(http.ListenAndServe(port, handlers.CORS(credentials, methods, origins)(router)))

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
