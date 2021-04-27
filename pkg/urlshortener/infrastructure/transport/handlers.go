package transport

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(srv *Server) http.Handler {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/shortener/v1").Subrouter()

	subRouter.HandleFunc("/{key}", srv.getRoute).Methods(http.MethodGet)
	subRouter.HandleFunc("/", srv.saveRoute).Methods(http.MethodPost)

	return subRouter
}

func (server Server) getRoute(w http.ResponseWriter, r *http.Request) {

}

func (server Server) saveRoute(w http.ResponseWriter, r *http.Request) {

}
