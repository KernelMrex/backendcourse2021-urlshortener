package model

import "net/url"

type Redirect struct {
	Key string
	Url *url.URL
}

type RedirectRepository interface {
	AddRedirect(redirect *Redirect) error
}
