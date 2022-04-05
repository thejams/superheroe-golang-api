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
	client := client.NewTradeMade(cfg.ClientURI)
	controller := controller.NewController(conn, client)

	http_server := httpServer.NewHTTPServer(ctx, cfg, controller)

	fmt.Printf("server runing in port %v \n", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, handlers.CORS(http_server.Credentials, http_server.Methods, http_server.Origins)(http_server.Router)))
}
