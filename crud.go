package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"

)
type User struct {
    Id int `bson:"id"`
    Category string `bson:"category"`
   
}
var collection *mongo.Collection

func init() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI("mongodb+srv://sunil:sunil@cluster0.bzpjx.mongodb.net/")
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    collection = client.Database("eproductdb").Collection("tblcategories")
}
func getAllUsers() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var user User
        if err := cursor.Decode(&user); err != nil {
            log.Fatal(err)
        }
        fmt.Println("User:", user)
    }
}


func main() {
    //user := User{Name: "Alice", Email: "alice@example.com", Age: 25}

   // insertUser(user)
    //getUserByName("Alice")
   // updateUserEmail("Alice", "newalice@example.com")
       getAllUsers()
   // deleteUserByName("Alice")
}

