package main

import (
	"log"
	"os"

	"github.com/TAhirr01/cliflags"
	"github.com/TAhirr01/confmaker"
	"github.com/TAhirr01/urlshortener/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	path := cliflags.RegisterFlag("ENV")
	cliflags.ParseFlags()
	godotenv.Load(path.Value)

	if err := confmaker.Load(os.Getenv("FILE_LOCATION"), &config.Conf); err != nil {
		log.Fatal("File location is empty or something else", err)
	}
	s := new(service)
	r := s.Start()
	log.Println("Server running on port 8080...")
	r.Run(":8080")
}
