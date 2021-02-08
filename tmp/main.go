package main

import (
	"context"
	"fmt"

	global "github.com/psinthorn/gostore/global/database"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	fmt.Println("Implement data generator")

	// Test connect to MongoDB database
	global.DB.Collection("laptop").InsertOne(context.Background(), bson.M{"name": "Dell"})
}
