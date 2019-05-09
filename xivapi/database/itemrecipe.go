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

// Pass information from jsonconv to this to input these values into the database.
func MongoInsert(itemname string, recipeid int, itemid int, crafttypeid int, ingredientid []int, ingredientamount []int) {
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

func MongoFind(fieldname string, fieldvalue int) {
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

	//Filters using a map to find the result, finding documents with  "field" : key
	filter := bson.M{fieldname: fieldvalue}
	var result Recipes

	//Actually find some with one filter
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	// We need to call a function to handle all these variables.
	//result.ItemName
	//result.RecipeID
	//result.ItemID
	//result.CraftTypeID
	//result.IngredientName[]
	//result.IngredientAmount[]

}
