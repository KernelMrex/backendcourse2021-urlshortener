package query

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"net/url"
	"urlshortener/pkg/urlshortener/app/query"
)

type redirectQueryService struct {
	db *sqlx.DB
}

func NewRedirectQueryService(db *sqlx.DB) query.RedirectQueryService {
	return &redirectQueryService{db: db}
}

func (r *redirectQueryService) GetRedirectByKey(key string) (*query.RedirectView, error) {
	row := r.db.QueryRow("SELECT destination from redirects WHERE key=?", key)

	var rawRedirectURL string
	if err := row.Scan(&rawRedirectURL); err != nil {
		return nil, query.ErrRedirectNotFound
	}

	redirectUrl, err := url.Parse(rawRedirectURL)
	if err != nil {
		return nil, errors.Wrap(err, "could not parse url")
	}

	return &query.RedirectView{Url: redirectUrl}, nil
}
