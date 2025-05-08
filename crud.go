package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	Id    int     `bson:"id`
	Name  string  `bson:"name`
	Price float64 `bson:"price"`
}

var collection *mongo.Collection

func init() {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading env file")
	}
	clientoption := options.Client().ApplyURI(os.Getenv("connectionstring"))
	client, _ := mongo.Connect(context, clientoption)

	collection = client.Database("product").Collection("productdetails")

}
func insert(p Product) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := collection.InsertOne(context, p)
	if err != nil {
		fmt.Println("error occured:", err)
	}
	fmt.Println(res.InsertedID)

}
func findAll() {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, _ := collection.Find(context, bson.M{}) //

	defer cursor.Close(context)

	for cursor.Next(context) {
		var p Product
		err := cursor.Decode(&p)
		if err != nil {
			fmt.Println("error occured:", err)
		} else {
			fmt.Println("Product:", p)
		}
	}

}
func findById(id int) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var p Product
	err := collection.FindOne(context, bson.M{"id": id}).Decode(&p)
	if err != nil {
		fmt.Println("error occured:", err)
	} else {
		fmt.Println("Product:", p)
	}
}
func updateById(id int, name string, price float64) {
	context, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	filter := bson.M{"id": id}

	update := bson.M{"$set": bson.M{"name": name, "price": price}}
	result, err := collection.UpdateOne(context, filter, update)
	if err != nil {
		fmt.Println("error occured:", err)
	} else {
		fmt.Println("Matched count:", result.MatchedCount)
		fmt.Println("Modified count:", result.ModifiedCount)
	}
}
func delete(id int) {
	context, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	res, errr := collection.DeleteOne(context, bson.M{"id": id})
	if errr != nil {
		fmt.Println("error occured:", errr.Error())
	} else {
		fmt.Println(res.DeletedCount)
	}
}
func main() {
	for {
		fmt.Println("Please choose correct option")
		var choice int
		fmt.Println("1.insert")
		fmt.Println("2.viewAll")
		fmt.Println("3.viewSpecific")
		fmt.Println("4.update")
		fmt.Println("5.delete")
		fmt.Println("6.exit")

		fmt.Scan(&choice)
		switch choice {
		case 1:
			var id, price = 0, 0.0
			fmt.Println("Enter id,name and price")
			fmt.Scan(&id)
			reader := bufio.NewReader(os.Stdin)
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Scan(&price)
			p := Product{Id: id, Name: name, Price: price}
			insert(p)
		case 2:
			findAll()
		case 3:
			eid := 0
			fmt.Println("Enter id")
			fmt.Scan(&eid)
			findById(eid)
		case 4:
			id := 0
			fmt.Println("Enter id")
			fmt.Scan(&id)
			fmt.Println("Enter name")

			reader := bufio.NewReader(os.Stdin)
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Println("Enter price")
			price := 0.0
			fmt.Scan(&price)
			updateById(id, name, price)
		case 5:
			id := 0
			fmt.Println("Enter id")
			fmt.Scan(&id)
			delete(id)

		case 6:
			fmt.Println("Exiting...")
			os.Exit(0)
		}
	}
	
	
}
