package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func init() {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading env file")
	}

	clientOption := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context, clientOption)
	if err != nil {
		panic(err)
	}
	db = client.Database("crud")
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
			p := CProduct{Name: name, Price: price}
			insert(p)
		case 2:
			findAll()
		case 3:
			eid := ""
			fmt.Println("Enter id")
			fmt.Scan(&eid)

			objectid, err := primitive.ObjectIDFromHex(eid)
			if err != nil {
				log.Fatal("Invalid object id")
			}

			findById(objectid)
		case 4:
			eid := ""
			fmt.Println("Enter id")
			fmt.Scan(&eid)

			objectid, err := primitive.ObjectIDFromHex(eid)
			if err != nil {
				log.Fatal("Invalid object id")
			}

			fmt.Println("Enter name")

			reader := bufio.NewReader(os.Stdin)
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Println("Enter price")
			price := 0.0
			fmt.Scan(&price)
			updateById(objectid, name, price)
		case 5:
			eid := ""
			fmt.Println("Enter id")
			fmt.Scan(&eid)
			objectid, err := primitive.ObjectIDFromHex(eid)
			if err != nil {
				log.Fatal("Invalid object id")
			}

			delete(objectid)

		case 6:
			fmt.Println("Exiting...")
			os.Exit(0)
		}
	}

}
