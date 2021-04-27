package service

import (
	"urlshortener/pkg/urlshortener/app/model"
)

type RedirectService interface {
	AddRedirect(redirect *model.Redirect) error
}

type redirectService struct {
	repository model.RedirectRepository
}

func NewRedirectService(repo model.RedirectRepository) RedirectService {
	return &redirectService{repository: repo}
}

func (rs redirectService) AddRedirect(redirect *model.Redirect) error {
	return rs.repository.AddRedirect(redirect.Key, redirect.Url)
}
