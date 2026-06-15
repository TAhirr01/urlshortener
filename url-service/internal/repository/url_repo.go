package repository

import (
	"context"
	"log"
	"time"

	"github.com/TAhirr01/urlshortener/urlservice/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlRepository interface {
	InsertShortURL(url *model.URL) error
	GetUrl(shortUrl string) (*model.URL, error)
}

type MongoUrlRepo struct {
	c *mongo.Collection
}

func NewmongoUrlRepo(c *mongo.Collection) *MongoUrlRepo {
	if c == nil {
		log.Fatal("mongo repo's mongo collection is null")
	}
	return &MongoUrlRepo{c: c}
}

func (m *MongoUrlRepo) InsertShortURL(url *model.URL) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := m.c.InsertOne(ctx, url)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoUrlRepo) GetUrl(shortUrl string) (*model.URL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url := new(model.URL)
	err := m.c.FindOne(ctx, bson.M{"short_code": shortUrl}).Decode(url)
	if err != nil {
		return nil, nil
	}
	log.Print("url repository", url)

	return url, nil
}
