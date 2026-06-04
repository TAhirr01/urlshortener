package service

import (
	"log"
	"time"

	"github.com/TAhirr01/urlshortener/internal/model"
	"github.com/TAhirr01/urlshortener/internal/repository"
	"github.com/TAhirr01/urlshortener/pkg"
)

type Service struct {
	r repository.Repository
}

func NewService(r repository.Repository) *Service {
	if r == nil {
		log.Fatal("Service's repository is null")
	}
	return &Service{r: r}
}

func (s *Service) CreateUrl(orgUrl string) (*model.URL, error) {
	shortCode := pkg.GenerateShortCode(6)
	url:=new(model.URL)
	url.OriginalURL = orgUrl
	url.ShortCode = shortCode
	url.CreatedAt = time.Now()

	if err := s.r.InsertShortURL(url); err != nil {
		return nil, err
	}
	return url, nil
}

func (s *Service) GetUrl(shortUrl string) (*model.URL, error) {
	url, err := s.r.GetUrl(shortUrl)
	log.Print("service url:",url.OriginalURL)
	if err != nil {
		return nil, err
	}
	return url, nil
}
