package main

import (
	"github.com/gin-gonic/gin"
)

type service struct {
	//DB connection

	//collection

	//repo

	//logger

	//redis

	//Services

	//handlers

}

func (s *service) Start() *gin.Engine {
	//DB connection

	s.DependencyInjection()

	server := NewServer()

	app := server.CreateRautes()

	return app
}

func (s *service) DependencyInjection() {

	//Url Collection

	//Repository initializtion

	//Service initialization

	//Handler initialization

}
