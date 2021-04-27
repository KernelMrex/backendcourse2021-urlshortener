package query

import (
	"errors"
	"urlshortener/pkg/urlshortener/app/model"
)

var (
	ErrRedirectNotFound = errors.New("redirect was not found")
)

type RedirectQueryService interface {
	GetRedirectByKey(key string) (*model.Redirect, error)
}
