package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

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

	db := repository.NewMongoDB()

	err := db.Conn(ctx, cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close(ctx)

	client := client.NewTradeMadeClient(cfg.TradeMadeClientURI)

	controller := controller.NewController(db, client)
	http_server := httpServer.NewHTTPServer(ctx, controller)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         cfg.Port,
		Handler:      handlers.CORS(http_server.Credentials, http_server.Methods, http_server.Origins)(http_server.Router),
	}

	fmt.Printf("server runing in port %v \n", cfg.Port)
	log.Fatal(srv.ListenAndServe())
}
