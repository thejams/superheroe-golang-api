package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	"superheroe-api/superheroe-golang-api/src/client"
	"superheroe-api/superheroe-golang-api/src/config"
	"superheroe-api/superheroe-golang-api/src/controller"
	"superheroe-api/superheroe-golang-api/src/httpServer"
	"superheroe-api/superheroe-golang-api/src/repository"
)

func main() {
	cfg := config.GetAPIConfig()
	ctx := context.Background()
	repo, err := repository.NewMongoConnection(ctx, cfg)
	if err != nil {
		panic(err)
	}

	client := client.NewTradeMade(cfg.ClientURI)
	controller := controller.NewController(repo, client)
	/* defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}() */
	defer repository.DisconnectDB(ctx)

	http_server := httpServer.NewHTTPServer(ctx, cfg, controller)

	fmt.Printf("server runing in port %v \n", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, handlers.CORS(http_server.Credentials, http_server.Methods, http_server.Origins)(http_server.Router)))
}
