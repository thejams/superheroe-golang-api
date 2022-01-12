package httpServer

import (
	"context"
	"net/http"
	"superheroe-api/superheroe-golang-api/src/client"
	"superheroe-api/superheroe-golang-api/src/controller"
	"superheroe-api/superheroe-golang-api/src/entity"
	"superheroe-api/superheroe-golang-api/src/repository"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type HttpServer struct {
	ctrl        controller.Controller
	ctx         context.Context
	Credentials handlers.CORSOption
	Methods     handlers.CORSOption
	Origins     handlers.CORSOption
	Router      *mux.Router
}

//NewHTTPServer initialice a new http server
func NewHTTPServer(ctx context.Context, cfg *entity.APPConfig, repo repository.Repository) *HttpServer {
	log.SetFormatter(&log.JSONFormatter{})
	var credentials handlers.CORSOption
	var methods handlers.CORSOption
	var origins handlers.CORSOption

	router := mux.NewRouter().StrictSlash(true)
	http_server := new(HttpServer)

	{
		client := client.NewTradeMade(cfg.ClientURI)
		ctrl := controller.NewController(repo, client)

		http_server.ctrl = ctrl
		http_server.ctx = ctx

		router.Use(commonMiddleware)
		http_server.initRouter(router)

		credentials = handlers.AllowCredentials()
		methods = handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
		origins = handlers.AllowedMethods([]string{"*"})

		http_server.Credentials = credentials
		http_server.Methods = methods
		http_server.Origins = origins
		http_server.Router = router
	}

	return http_server
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
