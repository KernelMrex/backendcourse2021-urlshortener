package service

import (
	"encoding/hex"
	"math/rand"
	"net/url"
	"urlshortener/pkg/urlshortener/app/model"
)

type RedirectService interface {
	AddRedirect(dest *url.URL) (string, error)
}

type redirectService struct {
	repository model.RedirectRepository
}

func NewRedirectService(repo model.RedirectRepository) RedirectService {
	return &redirectService{repository: repo}
}

func (rs *redirectService) AddRedirect(dest *url.URL) (string, error) {
	key, err := generateRandomKeyForRedirect(16)
	if err != nil {
		return "", err
	}

	if err := rs.repository.AddRedirect(&model.Redirect{Key: key, Url: dest}); err != nil {
		return "", err
	}

	return key, nil
}

func generateRandomKeyForRedirect(len uint) (string, error) {
	randBuffer := make([]byte, (len/2)+1)
	if _, err := rand.Read(randBuffer); err != nil {
		return "", err
	}
	hexString := hex.EncodeToString(randBuffer)
	return hexString[:len], nil
}
