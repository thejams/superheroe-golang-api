package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"superheroe-api/superheroe-golang-api/src/config"
	"superheroe-api/superheroe-golang-api/src/httpServer"
	"superheroe-api/superheroe-golang-api/src/repository"

	"github.com/gorilla/handlers"
)

func main() {
	cfg := config.GetAPIConfig()
	ctx := context.Background()
	repo, mongoClient := repository.NewMongoConnection(ctx)
	defer mongoClient.Disconnect(ctx)

	http_server := httpServer.NewHTTPServer(ctx, cfg, repo)

	fmt.Printf("server runing in port %v \n", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, handlers.CORS(http_server.Credentials, http_server.Methods, http_server.Origins)(http_server.Router)))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
