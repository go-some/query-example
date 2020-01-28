package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-some/crawler"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	id, pw := os.Getenv("DBID"), os.Getenv("DBPW")
	addrTemplate := os.Getenv("DBADDR")
	mongoDBAddr := fmt.Sprintf(addrTemplate, id, pw)

	clientOptions := options.Client().ApplyURI(mongoDBAddr)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("news")

	findOptions := options.Find()
	findOptions.SetLimit(5)

	var results []*crawler.News

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var news crawler.News
		err := cur.Decode(&news)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &news)
	}

	for _, doc := range results {
		fmt.Println("")
		fmt.Println("title:", doc.Title)
		fmt.Println("url:", doc.Url)
		fmt.Println("origin:", doc.Origin)
		fmt.Println("")
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
