package transport

import (
	"urlshortener/pkg/urlshortener/app/query"
	"urlshortener/pkg/urlshortener/app/service"
)

type Server struct {
	redirectQueryService query.RedirectQueryService
	redirectService      service.RedirectService
}

func NewServer(redirectService service.RedirectService) *Server {
	return &Server{
		redirectQueryService: nil,
		redirectService:      redirectService,
	}
}
