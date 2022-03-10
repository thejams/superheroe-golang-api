package httpServer

import "github.com/gorilla/mux"

func (h *HttpServer) initCharacterRouter(r *mux.Router) {
	routes := r.PathPrefix("/character").Subrouter()
	routes.Use(tokenValidatorMiddleware)
	routes.HandleFunc("/", h.GetSuperheroes).Methods("GET")
	routes.HandleFunc("/{id}", h.GetSuperhero).Methods("GET")
	routes.HandleFunc("/{id}", h.DeleteSuperhero).Methods("DELETE")
	routes.HandleFunc("/{id}", h.UpdateSuperhero).Methods("PUT")

	fieldValidatedRoutes := r.PathPrefix("/character").Subrouter()
	fieldValidatedRoutes.Use(tokenValidatorMiddleware)
	fieldValidatedRoutes.Use(fieldsValidatorMiddleware)
	fieldValidatedRoutes.HandleFunc("/", h.AddSuperHero).Methods("POST")
}
