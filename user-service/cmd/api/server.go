package main

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	//Handlers

}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) CreateRautes() *gin.Engine {
	app := gin.Default()
	//Here will be initializing rautes

	return app
}
