package query

import (
	"errors"
	"net/url"
)

var (
	ErrRedirectNotFound = errors.New("redirect was not found")
)

type RedirectView struct {
	Url *url.URL
}

type RedirectQueryService interface {
	GetRedirectByKey(key string) (*RedirectView, error)
}
