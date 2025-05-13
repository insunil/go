package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"name"`
	Price float64            `bson:"price"`
}
type CProduct struct {
	Name  string  `bson:"name"`
	Price float64 `bson:"price"`
}
