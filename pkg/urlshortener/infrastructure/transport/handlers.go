package transport

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
)

const (
	saveRedirectRoute = "/"
	getRedirectRoute  = "/{key}"
)

func NewRouter(srv *Server) http.Handler {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/shortener/v1").Subrouter()

	subRouter.HandleFunc(getRedirectRoute, srv.getRedirect).Methods(http.MethodGet)
	subRouter.HandleFunc(saveRedirectRoute, srv.saveRedirect).Methods(http.MethodPost)

	return subRouter
}

type errorResponse struct {
	Msg string `json:"msg"`
}

type saveRedirectResponse struct {
	ShortenUrl string `json:"shorten_url"`
}

func (server *Server) getRedirect(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (server *Server) saveRedirect(w http.ResponseWriter, r *http.Request) {
	rawDestUrl, ok := r.URL.Query()["dest"]
	if !ok || len(rawDestUrl) != 1 {
		if err := writeJSONResponse(w, errorResponse{Msg: "Invalid param 'dest'"}); err != nil {
			log.Error(errors.Wrap(err, "could not write JSON response"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	destUrl := rawDestUrl[0]

	parsedDestUrl, err := url.Parse(destUrl)
	if err != nil {
		if err := writeJSONResponse(w, errorResponse{Msg: "Could not parse provided url"}); err != nil {
			log.Error(errors.Wrap(err, "could not write JSON response"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	key, err := server.redirectService.AddRedirect(parsedDestUrl)
	if err != nil {
		log.Error(errors.Wrap(err, "could not add redirect"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := writeJSONResponse(w, saveRedirectResponse{
		ShortenUrl: fmt.Sprintf("%s://%s/%s", getHttpSchema(r), r.Host, key),
	}); err != nil {
		log.Error(errors.Wrap(err, "could not write JSON response"))
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func writeJSONResponse(w io.Writer, data interface{}) error {
	return errors.Wrap(json.NewEncoder(w).Encode(data), "could not encode or write response")
}

func getHttpSchema(r *http.Request) string {
	if r.TLS == nil {
		return "http"
	}
	return "https"
}
