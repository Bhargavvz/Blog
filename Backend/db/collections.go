package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	UsersCollection   *mongo.Collection
	PostsCollection   *mongo.Collection
	ContactCollection *mongo.Collection
)

func InitCollections(client *mongo.Client) {
	UsersCollection = client.Database("blog").Collection("users")
	PostsCollection = client.Database("blog").Collection("posts")
	ContactCollection = client.Database("blog").Collection("contacts")
}
