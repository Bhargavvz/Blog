package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Slug      string             `bson:"slug" json:"slug"`
	Summary   string             `bson:"summary" json:"summary"`
	Content   string             `bson:"content" json:"content"`
	Image     *Image             `bson:"image,omitempty" json:"image,omitempty"`
	Tags      []string           `bson:"tags" json:"tags"`
	Author    string             `bson:"author" json:"author"`
	ReadTime  int                `bson:"read_time" json:"readTime"` // in minutes
	Published bool               `bson:"published" json:"published"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updatedAt"`
}

type Image struct {
	Data      []byte    `bson:"data" json:"-"` // Binary image data
	Filename  string    `bson:"filename" json:"filename"`
	MimeType  string    `bson:"mime_type" json:"mimeType"`
	CreatedAt time.Time `bson:"created_at" json:"createdAt"`
}
