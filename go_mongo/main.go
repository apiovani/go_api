package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var databaseName = "go_mongo"
var collectionName = "teste"

type Task struct {
	ID   int64  `bson:"id"`
	Text string `bson:"text"`
}

func main() {
	// REGISTROS
	// res := insertRecord()
	res := updateRecord()
	// res := listRecords()

	if !res {
		fmt.Println("DEU RUIM")
		return
	}

	fmt.Println("AI SIMMM")
}

func connection() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://go_mongo-mongo/")

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func getCollectionAndCtx() (*mongo.Collection, context.Context) {
	var collection *mongo.Collection
	client := connection()

	collection = client.Database(databaseName).Collection(collectionName)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	return collection, ctx
}

// DATABASES
func createDatabase() bool {
	return false
}

func updateDatabase() bool {
	return false
}

func deleteDatabase() bool {
	return false
}

func listDatabases() bool {
	return false
}

// COLLECTIONS
func createCollection() bool {
	return false
}

func updateCollection() bool {
	return false
}

func deleteCollection() bool {
	return false
}

func listCollections() bool {
	return false
}

// DOCUMENTS
func insertRecord() bool {
	var task Task
	task.ID = 7
	task.Text = "Teste de texto7"

	collection, ctx := getCollectionAndCtx()

	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func updateRecord() bool {

	return false
}

func deleteRecord() bool {

	return false
}

func listRecords() bool {
	var tasks []*Task

	collection, ctx := getCollectionAndCtx()

	// findOptions := options.Find()
	// findOptions.SetLimit(5)

	// cur, err := collection.Find(ctx, bson.D{{}}, findOptions)

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println(err)
		return false
	}

	for cur.Next(ctx) {
		var t Task
		err := cur.Decode(&t)
		if err != nil {
			fmt.Println(err)
			return false
		}

		tasks = append(tasks, &t)
	}

	fmt.Println(tasks)

	return true
}