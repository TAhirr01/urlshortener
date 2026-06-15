package main

import (
	"log"

	"github.com/TAhirr01/urlshortener/urlservice/internal/handler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	//Handlers
	urlHandler *handler.URLHandler
}

func NewServer(urlHandler *handler.URLHandler) *Server {
	if urlHandler == nil {
		log.Fatal("Server's userhandler is null")
	}
	return &Server{
		urlHandler: urlHandler,
	}
}

func (s *Server) CreateRautes() *gin.Engine {
	app := gin.Default()
	//Here will be initializing rautes

	app.POST("/shorten", s.urlHandler.ShortenURL)
	app.GET("/:code", s.urlHandler.RedirectURL)
	return app
}
