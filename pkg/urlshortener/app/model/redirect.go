package model

import "net/url"

type Redirect struct {
	Key string
	Url *url.URL
}

type RedirectRepository interface {
	AddRedirect(key string, url *url.URL) error
}
