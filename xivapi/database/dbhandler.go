package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

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
