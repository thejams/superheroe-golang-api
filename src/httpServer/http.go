package httpServer

import (
	"context"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"superheroe-api/superheroe-golang-api/src/controller"
)

type HttpServer struct {
	ctrl        *controller.Controller
	ctx         context.Context
	Credentials handlers.CORSOption
	Methods     handlers.CORSOption
	Origins     handlers.CORSOption
	Router      *mux.Router
}

//NewHTTPServer initialice a new http server
func NewHTTPServer(ctx context.Context, ctrl *controller.Controller) *HttpServer {
	log.SetFormatter(&log.JSONFormatter{})
	http_server := new(HttpServer)

	{
		router := mux.NewRouter().StrictSlash(true)
		router.Use(commonMiddleware)
		http_server.initRouter(router)

		credentials := handlers.AllowCredentials()
		methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
		origins := handlers.AllowedOrigins([]string{"*"})

		http_server.Credentials = credentials
		http_server.Methods = methods
		http_server.Origins = origins
		http_server.Router = router
		http_server.ctrl = ctrl
		http_server.ctx = ctx
	}

	return http_server
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
