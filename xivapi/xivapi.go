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
	baseinfo := finditem(itemcollection, recipeID)
	baseprice := findprices(pricecollection, baseinfo.ItemResultTargetID)

	materialprices := make(map[int][10]int)
	materialingredients := make(map[int][]int)

	pricesarray(itemcollection, pricecollection, baseinfo, materialprices, materialingredients)

	// Find the base item prices
	// Sum these material prices.
	// Subtract from the first item price.
	_, basecurrent := avgprices(baseinfo.ItemResultTargetID, 1, baseprice)
	fmt.Println(basecurrent, materialprices, materialingredients)

}

func pricesarray(itemcollection *mongo.Collection, pricecollection *mongo.Collection, baseinfo *database.Recipes, materialprices map[int][10]int, materialingredients map[int][]int) {
	var pricearray [10]int
	for i := 0; i < len(baseinfo.IngredientNames); i++ {
		// Zero is an invalid material ID
		if baseinfo.IngredientNames[i] != 0 {
			prices := findprices(pricecollection, baseinfo.IngredientNames[i])

			pricearray[i] = prices.Sargatanas.Prices[0].PricePerUnit
		} else {
			continue
		}

	}
	materialprices[baseinfo.ItemResultTargetID] = pricearray
	materialingredients[baseinfo.ItemResultTargetID] = baseinfo.IngredientNames

	// If there's a recipe, we want to go in one more materialprices, and keep appending to it.
	for i := 0; i < len(baseinfo.IngredientRecipes); i++ {
		if len(baseinfo.IngredientRecipes[i]) != 0 {
			matinfo := finditem(itemcollection, baseinfo.IngredientRecipes[i][0])
			pricesarray(itemcollection, pricecollection, matinfo, materialprices, materialingredients)
		}
	}

}

func finditem(itemcollection *mongo.Collection, recipeID int) *database.Recipes {
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

	return itemresult
}
func findprices(pricecollection *mongo.Collection, itemID int) *database.Prices {
	// The find the price of the ingredient itself.
	priceresult := database.Ingredientprices(pricecollection, itemID)
	// TODO : Fix this into the Ingredientprices function instead.
	if priceresult.ItemID == 0 {
		byteValue := apipriceconnect(itemID)
		// TODO : create a json struct that has all these variables.
		prices := database.Jsonprices(byteValue)
		database.InsertPrices(pricecollection, *prices, itemID)

		priceresult = database.Ingredientprices(pricecollection, itemID)
	}

	return priceresult
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

	soldaverage := hissum / 2

	// Average Price Listings for the latest 20 entries.

	var listsum int
	for i := 0; i < len(matprices.Sargatanas.Prices) && i < 2; i++ {
		listsum = listsum + matprices.Sargatanas.Prices[i].PricePerUnit
	}

	currentaverage := listsum / 2

	// Multiply by the ingredient amount.
	averagesoldcost := soldaverage * ingredientamount
	averagecurrentcost := currentaverage * ingredientamount

	return averagesoldcost, averagecurrentcost
}
