package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

////// Recipes Struct////
type Recipes struct {
	Name               string `json:"Name"`
	ItemResultTargetID int    `json:"ItemResultTargetID"` // This is the Item ID
	ID                 int    `json:"ID"`                 // This is the recipeID
	CraftTypeTargetID  int    `json:"CraftTypeTargetID"`
}

/////Price Struct//////
type Prices struct {
	Sargatanas struct {
		ItemID  int `json:"ItemID"`
		History []struct {
			Added        int  `json:"Added"` // Time is in Unix epoch time
			IsHQ         bool `json:"IsHQ"`
			PricePerUnit int  `json:"PricePerUnit"`
			PriceTotal   int  `json:"PriceTotal"`
			PurchaseDate int  `json:"PurchaseDate"`
			Quantity     int  `json:"Quantity"`
		} `json:"History"`
		Prices []struct {
			Added        int  `json:"Added"`
			IsHQ         bool `json:"IsHQ"`
			PricePerUnit int  `json:"PricePerUnit"`
			PriceTotal   int  `json:"PriceTotal"`
			PurchaseDate int  `json:"PurchaseDate"`
			Quantity     int  `json:"Quantity"`
		} `json:"Prices"`
	} `json:"Sargatanas"`
}

// Pass information from jsonconv to this to input these values into the database.
func MongoInsertRecipe(recipes Recipes, ingredientid []int, ingredientamount []int) {
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
	Itemexample := bson.D{
		primitive.E{Key: "itemname", Value: recipes.Name},
		primitive.E{Key: "recipeid", Value: recipes.ItemResultTargetID},
		primitive.E{Key: "itemid", Value: recipes.ID},
		primitive.E{Key: "crafttypeid", Value: recipes.CraftTypeTargetID},
		primitive.E{Key: "ingredientname", Value: ingredientid},
		primitive.E{Key: "ingredientamount", Value: ingredientamount},
	}

	// This should insert the Itemexample into the document.
	insertResult, err := collection.InsertOne(context.TODO(), Itemexample)
	if err != nil {
		log.Fatal(err)
	}

	// Insertresult.InsertedID shows the objectID that we inserted this with.
	fmt.Println("Inserted Item into Database: ", insertResult.InsertedID)

}

func MongoInsertPrices(prices Prices) {
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
	collection := client.Database("Marketboard").Collection("Prices")

	Itemexample := bson.D{
		primitive.E{Key: "server", Value: prices.Sargatanas.ItemID},
		primitive.E{Key: "history", Value: prices.Sargatanas.History},
		primitive.E{Key: "prices", Value: prices.Sargatanas.Prices},
	}

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
