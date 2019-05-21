package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecipeResults struct {
	ItemName          string `bson:"itemname"`
	ItemID            int    `bson:"itemid"`
	RecipeID          int    `bson:"recipeid"`
	CraftTypeID       int    `bson:"crafttypeid"`
	IngredientNames   []int  `bson:"ingredientname"`
	IngredientAmounts []int  `bson:"ingredientamount"`
}

type PriceResults struct {
	ItemID  int `bson:"itemid"`
	Servers struct {
		Sargatanas struct {
			History []struct {
				Added        int  `bson:"added"` // Time is in Unix epoch time
				IsHQ         bool `bson:"ishq"`
				PricePerUnit int  `bson:"priceperunit"`
				PriceTotal   int  `bson:"pricetotal"`
				PurchaseDate int  `bson:"purchasedate"`
				Quantity     int  `bson:"quantity"`
			} `bson:"history"`
			Prices []struct {
				Added        int  `bson:"added"`
				IsHQ         bool `bson:"ishq"`
				PricePerUnit int  `bson:"priceperunit"`
				PriceTotal   int  `bson:"pricetotal"`
				Quantity     int  `bson:"quantity"`
			} `bson:"prices"`
		} `bson:"sargatanas"`
	} `bson:"servers"`
}

// Calls ingredient amounts and item IDs, and returns the results
func Ingredientmaterials(collection *mongo.Collection, recipeID int) *RecipeResults {
	filter := bson.M{"recipeid": recipeID}
	var result RecipeResults
	collection.FindOne(context.TODO(), filter).Decode(&result)

	return &result
}

// Call the prices from the database, and return the sold average and the current average
func Ingredientprices(collection *mongo.Collection, itemID int) *PriceResults {
	filter := bson.M{"itemid": itemID}
	var result PriceResults
	collection.FindOne(context.TODO(), filter).Decode(&result)

	return &result
}

// Pass information from jsonconv to this to input these values into the database.
func InsertRecipe(collection *mongo.Collection, recipes Recipes, ingredientid []int, ingredientamount []int) {

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

func InsertPrices(collection *mongo.Collection, prices Prices, userID int) {
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
