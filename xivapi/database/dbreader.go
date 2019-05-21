package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
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
