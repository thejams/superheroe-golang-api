package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"

	"superheroe-api/superheroe-golang-api/src/config"
	"superheroe-api/superheroe-golang-api/src/controller"
	"superheroe-api/superheroe-golang-api/src/factory"
	"superheroe-api/superheroe-golang-api/src/httpServer"
)

func main() {
	cfg := config.GetAPIConfig()
	ctx := context.Background()

	conn := factory.DBFactory(1)
	if conn == nil {
		panic("DB engine not found")
	}

	err := conn.Conn(ctx, cfg)
	if err != nil {
		panic(err)
	}

	defer conn.Close(ctx)

	client := factory.ClientFactory(1)
	if client == nil {
		panic("Client not found")
	}
	client.InitClient(cfg)

	controller := controller.NewController(conn, client)
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
