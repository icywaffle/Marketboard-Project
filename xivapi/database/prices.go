package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PriceResults struct {
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

func Profits(itemID int) (int, int) {

	// Connect to the database and fill out the struct

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("Marketboard").Collection("Prices")
	filter := bson.M{"itemid": itemID}

	var result PriceResults
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	// Now manipulate the struct data here

	// Average Price History for the latest 20 entries.
	var hissum int
	var hisN int
	for i := 0; i < len(result.Servers.Sargatanas.History) && i < 20; i++ {
		hissum = hissum + result.Servers.Sargatanas.History[i].PriceTotal
		hisN = hisN + result.Servers.Sargatanas.History[i].Quantity
	}

	soldaverage := hissum / hisN

	// Average Price Listings for the latest 20 entries.

	var listsum int
	var listN int
	for i := 0; i < len(result.Servers.Sargatanas.Prices) && i < 20; i++ {
		listsum = listsum + result.Servers.Sargatanas.Prices[i].PriceTotal
		listN = listN + result.Servers.Sargatanas.Prices[i].Quantity
	}

	currentaverage := listsum / listN

	return soldaverage, currentaverage

}

func TotalProfits(itemID int) {
	// We need a function that takes the information from the recipes database, and pair it with the prices.
	// Connect to the database and fill out the struct

	// Find an item with an item ID.
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("Marketboard").Collection("Recipes")
	filter := bson.M{"itemid": itemID}

	var result RecipeResults
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	// Find the ingredients.
	ingredients := result.IngredientNames

	var Phistcraft int
	var Pcraft int
	for i := 0; i < len(ingredients); i++ {
		// Find the prices of those ingredients.
		if ingredients[i] != 0 {
			ingrsoldcost, ingrcurrentcost := Profits(ingredients[i])
			// Multiply by the ingredient amount.
			totalsoldcost := ingrsoldcost * result.IngredientAmounts[i]
			totalcurrentcost := ingrcurrentcost * result.IngredientAmounts[i]
			// Then add them all to find the sum, for it's cost.
			Phistcraft += totalsoldcost
			Pcraft += totalcurrentcost
		} else {
			break
		}
	}

	// The find the price of the ingredient itself.
	mainprice, _ := Profits(itemID)
	fmt.Printf("Price of Material: %v \n", mainprice)

	// Find the difference Pitem-Pcraft
	difference := mainprice - Pcraft
	fmt.Printf("Difference: %v \n", difference)
}
