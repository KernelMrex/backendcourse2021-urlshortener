package repository

import (
	"github.com/jmoiron/sqlx"
	"urlshortener/pkg/urlshortener/app/model"
)

func NewRedirectRepository(db *sqlx.DB) model.RedirectRepository {
	return &redirectRepository{
		db: db,
	}
}

type redirectRepository struct {
	db *sqlx.DB
}

func (r *redirectRepository) AddRedirect(redirect *model.Redirect) error {
	_, err := r.db.Exec("INSERT INTO redirects(`key`, `destination`) VALUES (?, ?)", redirect.Key, redirect.Url.String())
	return err
}
