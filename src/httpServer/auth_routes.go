package httpServer

import "github.com/gorilla/mux"

func (h *HttpServer) initAuthRouter(r *mux.Router) {
	r.HandleFunc("/signin", h.Signup).Methods("POST")

	protectedRoutes := r.PathPrefix("/auth").Subrouter()
	protectedRoutes.Use(tokenValidatorMiddleware)
	protectedRoutes.HandleFunc("/refresh", h.RefreshToken).Methods("GET")
}
