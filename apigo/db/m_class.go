package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Db_mongo struct {
	url        string
	collection *mongo.Collection
}

func (db *Db_mongo) Db_start() {
	db.url = "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(db.url))
	if err != nil {
		fmt.Print(err)
	}
	ctx, cte := context.WithTimeout(context.Background(), 10*time.Second)
	defer cte()
	err = client.Connect(ctx)
	collection := client.Database("WEB").Collection("webidk")

	db.collection = collection
}
func (db Db_mongo) Db_InsertOne(Insert map[string]string) {

	_, err := db.collection.InsertOne(context.TODO(), Insert)
	fmt.Print(err)

}
