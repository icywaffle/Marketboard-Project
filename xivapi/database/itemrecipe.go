package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Structure of the basic Item Recipe Schema.
type Recipes struct {
	ItemName         string
	RecipeID         int
	ItemID           int
	CraftTypeID      int
	IngredientName   []int
	IngredientAmount []int
}

func MongoHandler(itemname string, recipeid int, itemid int, crafttypeid int, ingredientid []int, ingredientamount []int) {
	/*
		for {
			var input int
			fmt.Printf("Mongo Case 1,2:")
			n, err := fmt.Scanln(&input)
			// Force choose a positive number
			if n < 1 || err != nil {
				fmt.Println("invalid input")
				os.Exit(2)
			}

			switch input {
			case 1:
				mongoInsert(itemname, recipeid, itemid, ingredientid, ingredientamount)
			case 2:
				mongoFind()
			default:
				fmt.Println("Invalid Case Selected.")
				continue
			}
		}
	*/
	//For now, just insert into database.

	mongoInsert(itemname, recipeid, itemid, crafttypeid, ingredientid, ingredientamount)
}

// Pass information from jsonconv to this to input these values into the database.
func mongoInsert(itemname string, recipeid int, itemid int, crafttypeid int, ingredientid []int, ingredientamount []int) {
	//Sets the Client Options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") //There are many client options available.
	//Connect to the MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	//This is the colleciton that we're accessing, from our database Marketboard, and from the collecion, Items.
	collection := client.Database("Marketboard").Collection("Recipes")

	//This is an example item that should be inserted into the existing document
	Itemexample := Recipes{itemname, recipeid, itemid, crafttypeid, ingredientid, ingredientamount}

	// This should insert the Itemexample into the document.
	insertResult, err := collection.InsertOne(context.TODO(), Itemexample)
	if err != nil {
		log.Fatal(err)
	}

	// Insertresult.InsertedID shows the objectID that we inserted this with.
	fmt.Println("Inserted Item into Database: ", insertResult.InsertedID)

}

func mongoFind() {
	//Sets the Client Options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") //There are many client options available.
	//Connect to the MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	//This is the colleciton that we're accessing, from our database Marketboard, and from the collecion, Items.
	collection := client.Database("Marketboard").Collection("Items")

	//Filters using a map to find the result, finding documents with  "field" : key
	filter := bson.M{"recipeid": 123123123}
	var result Recipes

	//Actually find some with one filter
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

}
