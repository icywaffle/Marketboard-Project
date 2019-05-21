package xivapi

import (
	"context"
	"encoding/json"
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
	// Connect to the database
	itemcollection := dbconnect("Recipes")
	pricecollection := dbconnect("Prices")
	// itemresult is the info in the recipeID
	itemresult := database.Ingredientmaterials(itemcollection, recipeID)
	// If the item is not in the database, then we should add it. 0 is an invalid itemID
	if itemresult.ItemID == 0 {
		dbinsertitemrecipe(itemcollection, apirecipeconnect(recipeID), recipeID)
		itemresult = database.Ingredientmaterials(itemcollection, recipeID)
	}

	// We need to calculate the average price on these ingredients.
	ingredients := itemresult.IngredientNames
	ingredientamounts := itemresult.IngredientAmounts
	var mathisprice int
	var matcurrentprice int
	for i := 0; i < len(ingredients); i++ {
		if ingredients[i] != 0 {
			// priceresult is the prices for the itemID
			priceresult := database.Ingredientprices(pricecollection, ingredients[i])
			// If the item is not in the database, then we should add it. 0 is an invalid itemID
			if priceresult.ItemID == 0 {
				dbinsertprice(pricecollection, apipriceconnect(ingredients[i]), ingredients[i])
				priceresult = database.Ingredientprices(pricecollection, ingredients[i])
			}
			history, current := avgprices(ingredients[i], ingredientamounts[i], priceresult)
			mathisprice += history
			matcurrentprice += current
			fmt.Printf("Price: %v , of : %v \n", matcurrentprice, ingredients[i])
		} else {
			continue
		}

	}

	// The find the price of the ingredient itself.
	priceresult := database.Ingredientprices(pricecollection, itemresult.ItemID)
	// TODO : Fix this into the Ingredientprices function instead.
	if priceresult.ItemID == 0 {
		dbinsertprice(pricecollection, apipriceconnect(itemresult.ItemID), itemresult.ItemID)
		priceresult = database.Ingredientprices(pricecollection, itemresult.ItemID)
	}
	mainhisprice, maincurrentprice := avgprices(itemresult.ItemID, 1, priceresult) //There's only 1 ingredient amount, which is the main item.

	fmt.Printf("History of Main Item: %v \n", mainhisprice)
	fmt.Printf("Current of Main Item: %v \n", maincurrentprice)
	fmt.Printf("History of Total Materials: %v \n", mathisprice)
	fmt.Printf("Current of Total Materials: %v \n", matcurrentprice)

	fmt.Printf("Profits from buying Materials: %v \n", mainhisprice-matcurrentprice)

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

func dbinsertitemrecipe(collection *mongo.Collection, byteValue []byte, recipeID int) {

	// Unmarshal the information into the structs
	var recipes database.Recipes
	json.Unmarshal(byteValue, &recipes)

	var amount database.AmountIngredient
	json.Unmarshal(byteValue, &amount)

	var matitemID database.ItemIngredient
	json.Unmarshal(byteValue, &matitemID)

	// Create the slices
	amountslice := []int{amount.AmountIngredient0,
		amount.AmountIngredient1,
		amount.AmountIngredient2,
		amount.AmountIngredient3,
		amount.AmountIngredient4,
		amount.AmountIngredient5,
		amount.AmountIngredient6,
		amount.AmountIngredient7,
		amount.AmountIngredient8,
		amount.AmountIngredient9}

	matitemIDslice := []int{matitemID.ItemIngredient0TargetID,
		matitemID.ItemIngredient1TargetID,
		matitemID.ItemIngredient2TargetID,
		matitemID.ItemIngredient3TargetID,
		matitemID.ItemIngredient4TargetID,
		matitemID.ItemIngredient5TargetID,
		matitemID.ItemIngredient6TargetID,
		matitemID.ItemIngredient7TargetID,
		matitemID.ItemIngredient8TargetID,
		matitemID.ItemIngredient9TargetID}

	database.InsertRecipe(collection, recipes, matitemIDslice, amountslice)
}

func dbinsertprice(collection *mongo.Collection, byteValue []byte, itemID int) {

	var prices database.Prices
	json.Unmarshal(byteValue, &prices)

	database.InsertPrices(collection, prices, itemID)

}

func avgprices(ingredient int, ingredientamount int, matprices *database.PriceResults) (int, int) {

	// Average Price History for the latest 20 entries.
	var hissum int
	var hisN int
	for i := 0; i < len(matprices.Servers.Sargatanas.History) && i < 20; i++ {
		hissum = hissum + matprices.Servers.Sargatanas.History[i].PriceTotal
		hisN = hisN + matprices.Servers.Sargatanas.History[i].Quantity
	}

	soldaverage := hissum / hisN

	// Average Price Listings for the latest 20 entries.

	var listsum int
	var listN int
	for i := 0; i < len(matprices.Servers.Sargatanas.Prices) && i < 20; i++ {
		listsum = listsum + matprices.Servers.Sargatanas.Prices[i].PriceTotal
		listN = listN + matprices.Servers.Sargatanas.Prices[i].Quantity
	}

	currentaverage := listsum / listN

	// Multiply by the ingredient amount.
	averagesoldcost := soldaverage * ingredientamount
	averagecurrentcost := currentaverage * ingredientamount

	return averagesoldcost, averagecurrentcost
}
