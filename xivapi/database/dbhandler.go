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

type RecipeResults struct {
	ItemName          string `bson:"itemname"`
	ItemID            int    `bson:"itemid"`
	RecipeID          int    `bson:"recipeid"`
	CraftTypeID       int    `bson:"crafttypeid"`
	IngredientNames   []int  `bson:"ingredientname"`
	IngredientAmounts []int  `bson:"ingredientamount"`
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
		primitive.E{Key: "itemid", Value: recipes.ItemResultTargetID},
		primitive.E{Key: "recipeid", Value: recipes.ID},
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

func MongoInsertPrices(prices Prices, userID int) {
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
		// For here, we need to write this code for each individual server.
		primitive.E{Key: "itemid", Value: userID},
		primitive.E{Key: "servers", Value: prices},
	}

	// This should insert the Itemexample into the document.
	insertResult, err := collection.InsertOne(context.TODO(), Itemexample)
	if err != nil {
		log.Fatal(err)
	}

	// Insertresult.InsertedID shows the objectID that we inserted this with.
	fmt.Println("Inserted Item into Database: ", insertResult.InsertedID)
}

func MongoFind(collectionname string, fieldname string, fieldvalue int) bool {
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
	collection := client.Database("Marketboard").Collection(collectionname)

	//Filters using a map to find the result, finding documents with  "field" : key
	filter := bson.M{fieldname: fieldvalue}

	// TODO: we can create this struct in a different file, and pass it through the function, so that we can manipulate this later.
	var result RecipeResults

	//Actually find some with one filter
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		// This is where we find nothing.
		return false
	}

	return true

}
