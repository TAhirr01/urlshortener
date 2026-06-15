package main

import (
	"log"
	"os"

	"github.com/TAhirr01/cliflags"
	"github.com/TAhirr01/confmaker"
	"github.com/TAhirr01/urlshortener/urlservice/internal/config"
	"github.com/joho/godotenv"
)
const (
	FILE_LOCATION string="FILE_LOCATION"
	ENV string="ENV"
)


func main() {
	path := cliflags.RegisterFlag(ENV)
	cliflags.ParseFlags()
	if path.Value!=""{
		if _,err:=os.Stat(path.Value);err==nil{
			godotenv.Load(path.Value)
		}
	}
	yamlPath:=os.Getenv(FILE_LOCATION)
	if yamlPath==""{
		log.Fatal(FILE_LOCATION,": enviroment variable is missing from system memory")
	}
	if err := confmaker.Load(yamlPath, &config.Conf); err != nil {
		log.Fatal("Failed to load YAML configuration", err)
	}
	s := new(service)
	r := s.Start()
	log.Println("Server running on port 8080...")
	r.Run(":8080")
}
