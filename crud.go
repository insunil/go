package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func insert(p CProduct) {
	collection := db.Collection("product")

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(context, p)
	if err != nil {
		fmt.Println("error occurred:", err)
	}
	fmt.Println(res.InsertedID)

}

func findAll() {
	collection := db.Collection("product")

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, _ := collection.Find(context, bson.M{}) //

	defer cursor.Close(context)

	for cursor.Next(context) {
		var p Product
		err := cursor.Decode(&p)
		if err != nil {
			fmt.Println("error occurred:", err)
		} else {
			fmt.Println("Product:", p)
		}
	}

}
func findById(id primitive.ObjectID) {
	collection := db.Collection("product")

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var p Product

	err := collection.FindOne(context, bson.M{"_id": id}).Decode(&p)
	if err != nil {
		fmt.Println("error occurred:", err)
	} else {
		fmt.Println("Product:", p)
	}
}

func updateById(id primitive.ObjectID, name string, price float64) {
	collection := db.Collection("product")

	context, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"name": name, "price": price}}
	result, err := collection.UpdateOne(context, filter, update)
	if err != nil {
		fmt.Println("error occurred:", err)
	} else {
		fmt.Println("Matched count:", result.MatchedCount)
		fmt.Println("Modified count:", result.ModifiedCount)
	}
}

func delete(id primitive.ObjectID) {
	collection := db.Collection("product")

	context, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	res, err := collection.DeleteOne(context, bson.M{"_id": id})
	if err != nil {
		fmt.Println("error occurred:", err.Error())
	} else {
		fmt.Println(res.DeletedCount)
	}
}
