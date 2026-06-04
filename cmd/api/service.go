package main

import (
	"fmt"
	"log"

	"github.com/TAhirr01/urlshortener/internal/config"
	"github.com/TAhirr01/urlshortener/internal/handler"
	"github.com/TAhirr01/urlshortener/internal/repository"
	urlService "github.com/TAhirr01/urlshortener/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type service struct {
	//DB connection
	client *mongo.Client
	//collection
	urlCollection *mongo.Collection
	//repo
	repo repository.Repository

	//logger

	//redis

	//Services
	urlService *urlService.Service
	//handlers
	urlHandler *handler.URLHandler
}

func (s *service) Start() *gin.Engine {
	//MongoDB connection
	log.Print("Opening MongoDB connection")
	client, err := config.ConnectDB(config.Conf.Application.MongoDB.URI)
	fmt.Println("URI:", config.Conf.Application.MongoDB.URI)
	if err != nil {
		log.Fatal("Cannon establish mongoDB connection:", err)
	}
	s.client = client
	s.DependencyInjection()

	server := NewServer(s.urlHandler)

	app := server.CreateRautes()

	return app
}

func (s *service) DependencyInjection() {

	//Url Collection
	s.urlCollection = config.GetCollection(s.client, "urls")

	//Repository initializtion
	s.repo = repository.NewMongoRepo(s.urlCollection)

	//Service initialization
	s.urlService = urlService.NewService(s.repo)

	//Handler initialization
	s.urlHandler = handler.NewURLHandler(s.urlService)

}
