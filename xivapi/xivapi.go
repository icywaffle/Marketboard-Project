package xivapi

import (
	"context"
	"fmt"
	"log"
	"time"

	database "./database"
	urlstring "./urlstring"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Current issues.
// We need to remove outliers from the price calculations.
// We have to go into the recipes, and find those too.
func NetItemPrice(recipeID int) {

	// Connect to the database first.
	itemcollection := dbconnect("Recipes")
	pricecollection := dbconnect("Prices")

	// Using the recipeID, find the materials.
	// Find the base item price.
	basemats, baseprice := findinfo(itemcollection, pricecollection, recipeID)

	// Using the materials, use the recipes that are already included.
	// Using those recipes, find the materials.

	tier := 0
	pricepertier := make([]int, 0)
	pricepertier = matinfo(itemcollection, pricecollection, basemats, tier, pricepertier)

	// Find the base item prices
	// Sum these material prices.
	// Subtract from the first item price.
	for i := 0; i < len(pricepertier); i++ {
		_, currentbase := avgprices(basemats.ID, basemats.IngredientAmounts[i], baseprice)
		fmt.Println("Profit", pricepertier[i]-currentbase, "Tier", i)
	}

}

func matinfo(itemcollection *mongo.Collection, pricecollection *mongo.Collection, baseitem *database.Recipes, tier int, pricepertier []int) []int {

	for i := 0; i < len(baseitem.IngredientRecipes); i++ {
		for j := 0; j < len(baseitem.IngredientRecipes[i]); j++ {
			// Zero is an invalid recipeID
			if baseitem.IngredientRecipes[i][j] != 0 {
				matofmats, matprice := findinfo(itemcollection, pricecollection, baseitem.IngredientRecipes[i][j])
				// Append to the price per tier, the average price.
				_, currentcost := avgprices(matofmats.ID, baseitem.IngredientAmounts[i], matprice)
				pricepertier = append(pricepertier, currentcost)

				// Find the other mat recipes.
				// Mark these recipes as a tier and continue adding to that tier if they're later called inside the function.
				tier += 1
				matinfo(itemcollection, pricecollection, matofmats, tier, pricepertier)
			} else {
				continue
			}

		}

	}

	return pricepertier
}

func findinfo(itemcollection *mongo.Collection, pricecollection *mongo.Collection, recipeID int) (*database.Recipes, *database.Prices) {
	// itemresult is the info in the recipeID
	itemresult := database.Ingredientmaterials(itemcollection, recipeID)
	// If the item is not in the database, then we should add it. 0 is an invalid itemID
	if itemresult.ID == 0 {
		byteValue := apirecipeconnect(recipeID)
		// TODO : create a json struct that has all these variables.
		recipes, matIDs, amounts, matrecipes := database.Jsonitemrecipe(byteValue)
		database.InsertRecipe(itemcollection, *recipes, matIDs, amounts, matrecipes)

		itemresult = database.Ingredientmaterials(itemcollection, recipeID)
	}
	// The find the price of the ingredient itself.
	priceresult := database.Ingredientprices(pricecollection, itemresult.ID)
	// TODO : Fix this into the Ingredientprices function instead.
	if priceresult.ItemID == 0 {
		byteValue := apipriceconnect(itemresult.ItemResultTargetID)
		// TODO : create a json struct that has all these variables.
		prices := database.Jsonprices(byteValue)
		database.InsertPrices(pricecollection, *prices, itemresult.ItemResultTargetID)

		priceresult = database.Ingredientprices(pricecollection, itemresult.ItemResultTargetID)
	}

	return itemresult, priceresult
}

func apirecipeconnect(recipeID int) []byte {
	// MAX Rate limit is 20 Req/s -> 0.05s/Req, but safer to use 15req/s -> 0.06s/req
	time.Sleep(100 * time.Millisecond)
	// This ensures that when this function is called, it does not exceed the rate limit.
	// TODO: Use a channel to rate limit instead to allow multiple users to use this.

	websiteurl := urlstring.UrlItemRecipe(recipeID)
	byteValue := urlstring.XiviapiRecipeConnector(websiteurl)
	return byteValue
}

func apipriceconnect(itemID int) []byte {
	// MAX Rate limit is 20 Req/s -> 0.05s/Req, but safer to use 15req/s -> 0.06s/req
	time.Sleep(100 * time.Millisecond)
	// This ensures that when this function is called, it does not exceed the rate limit.
	// TODO: Use a channel to rate limit instead to allow multiple users to use this.

	websiteurl := urlstring.UrlPrices(itemID)
	byteValue := urlstring.XiviapiRecipeConnector(websiteurl)
	return byteValue
}

// Collection names are either "Prices" or "Recipes"
func dbconnect(collectionname string) *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("Marketboard").Collection(collectionname)

	return collection
}

func avgprices(ingredient int, ingredientamount int, matprices *database.Prices) (int, int) {

	// Average Price History for the latest 20 entries.
	var hissum int
	for i := 0; i < len(matprices.Sargatanas.History) && i < 2; i++ {
		hissum = hissum + matprices.Sargatanas.History[i].PricePerUnit
	}

	soldaverage := hissum

	// Average Price Listings for the latest 20 entries.

	var listsum int
	for i := 0; i < len(matprices.Sargatanas.Prices) && i < 2; i++ {
		listsum = listsum + matprices.Sargatanas.Prices[i].PricePerUnit
	}

	currentaverage := listsum

	// Multiply by the ingredient amount.
	averagesoldcost := soldaverage * ingredientamount
	averagecurrentcost := currentaverage * ingredientamount

	return averagesoldcost, averagecurrentcost
}
