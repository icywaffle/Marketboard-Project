package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Calls ingredient amounts and item IDs, and returns the results
func Ingredientmaterials(collection *mongo.Collection, recipeID int) *Recipes {
	recipefilter := bson.M{"RecipeID": recipeID}
	var reciperesult Recipes
	collection.FindOne(context.TODO(), recipefilter).Decode(&reciperesult)

	return &reciperesult
}

// Looks through the database, and finds if there's an itemID in the prices.
func Ingredientprices(collection *mongo.Collection, itemID int) *Prices {
	pricefilter := bson.M{"ItemID": itemID}
	var priceresult Prices
	collection.FindOne(context.TODO(), pricefilter).Decode(&priceresult)

	return &priceresult
}

// Pass information from jsonconv to this to input these values into the database.
func InsertRecipe(collection *mongo.Collection, recipes Recipes, ingredientid []int, ingredientamount []int, ingredientrecipes [][]int) {

	//This is an example item that should be inserted into the existing document
	Itemexample := bson.D{
		primitive.E{Key: "Name", Value: recipes.Name},
		primitive.E{Key: "ItemID", Value: recipes.ItemResultTargetID},
		primitive.E{Key: "RecipeID", Value: recipes.ID},
		primitive.E{Key: "CraftTypeTargetID", Value: recipes.CraftTypeTargetID},
		primitive.E{Key: "IngredientName", Value: ingredientid},
		primitive.E{Key: "IngredientAmount", Value: ingredientamount},
		primitive.E{Key: "IngredientRecipes", Value: ingredientrecipes},
	}

	// This should insert the Itemexample into the document.
	insertResult, err := collection.InsertOne(context.TODO(), Itemexample)
	if err != nil {
		log.Fatal(err)
	}

	// Insertresult.InsertedID shows the objectID that we inserted this with.
	fmt.Println("Inserted Item into Database: ", insertResult.InsertedID)

}

func InsertPrices(collection *mongo.Collection, prices Prices, itemID int) {
	Itemexample := bson.D{
		// For here, we need to write this code for each individual server.
		primitive.E{Key: "ItemID", Value: itemID},
		primitive.E{Key: "Sargatanas", Value: prices.Sargatanas},
	}

	// This should insert the Itemexample into the document.
	insertResult, err := collection.InsertOne(context.TODO(), Itemexample)
	if err != nil {
		log.Fatal(err)
	}

	// Insertresult.InsertedID shows the objectID that we inserted this with.
	fmt.Println("Inserted Item into Database: ", insertResult.InsertedID)
}
