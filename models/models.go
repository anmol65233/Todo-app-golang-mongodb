package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ToDoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty"`
	Status bool               `json:"status,omitempty"`
}

const connectionString = "mongodb://localhost:27017"

// Database Name
const dbName = "gotodo"

// Collection name
const collName = "todolist"

// collection object/instance
var collection *mongo.Collection

func init() {

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(" ✅ Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println(" 🥂 Collection instance created!")
}

func GetAllTask() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

func CreateTask(task ToDoList) string {
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
	str, ok := insertResult.InsertedID.(string)
	if ok {
		fmt.Println(ok)
	}
	return str
}

func GetById(id string) primitive.M {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}
	var result bson.M
	data := collection.FindOne(context.Background(), bson.M{"_id": objectId})
	e := data.Decode(&result)
	if e != nil {

		return result
	}
	//data.Close()
	result["result"] = true
	return result
}

func UpdateStatus(id string, status bool) int64 {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"status": status}}
	fmt.Println(objectId)
	data, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return 0
	}

	return data.ModifiedCount
}

func DeleteOne(id string) int64 {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}
	filter := bson.M{"_id": objectId}
	fmt.Println(objectId)
	data, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return 0
	}

	return data.DeletedCount
}
