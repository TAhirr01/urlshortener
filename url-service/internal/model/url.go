package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type URL struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OriginalURL string             `bson:"original_url" json:"original_url" binding:"required"`
	ShortCode   string             `bson:"short_code" json:"short_code"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	Clicks      int64              `bson:"clicks" json:"clicks"`
}
