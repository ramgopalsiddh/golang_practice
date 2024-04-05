package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://ramgopalsiddh:<password>@cluster0.6gq3p7r.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"


// most important part

var collection *mongo.Collection

// connect with mongoDB

func init(){
	// create clientOption/ connection
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	// context.TODO this use when you unclear which context to use [ https://pkg.go.dev/context#TODO ]
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connected successful")

	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("Collection instance/reference is ready")
}