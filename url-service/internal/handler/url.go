package handler

import (
	"log"
	"net/http"

	"github.com/TAhirr01/urlshortener/urlservice/internal/service"
	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	s *service.Service
}

func NewURLHandler(s *service.Service) *URLHandler {
	if s == nil {
		log.Fatal("Service instance in null")
	}
	return &URLHandler{s: s}
}

// POST /shorten
func (uc *URLHandler) ShortenURL(c *gin.Context) {
	var uri struct {
		OrgURL string `json:"original_url"`
	}

	if err := c.ShouldBindJSON(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	url, err := uc.s.CreateUrl(uri.OrgURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannont shorten url"})
		return
	}
	// Return the shortened URL details
	c.JSON(http.StatusOK, gin.H{
		"original_url": url.OriginalURL,
		"short_url":    "http://localhost:8080/" + url.ShortCode,
	})
}

// GET /:code
func (uc *URLHandler) RedirectURL(c *gin.Context) {
	code := c.Param("code")

	url, err := uc.s.GetUrl(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Url not found"})
		return
	}
	log.Print("original url", url.OriginalURL)
	// Redirect to the original destination
	c.Redirect(http.StatusFound, url.OriginalURL)
}
