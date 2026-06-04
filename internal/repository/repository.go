package repository

import "github.com/TAhirr01/urlshortener/internal/model"

type Repository interface {
	InsertShortURL(url *model.URL) error
	GetUrl(shortUrl string) (*model.URL, error)
}
