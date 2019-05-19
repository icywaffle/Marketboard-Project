package xivapi

import (
	"encoding/json"

	database "./database"
)

func dbinsertitemrecipe(byteValue []byte, recipeID int) {
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

	database.MongoInsertRecipe(recipes, amountslice, matitemIDslice)
}

func dbinsertprice(byteValue []byte, itemID int) {

	var prices database.Prices
	json.Unmarshal(byteValue, &prices)

}
