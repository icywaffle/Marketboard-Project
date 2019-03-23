// Converts Recipe Pages of json, to arrays.

package xivapi

import (
	"encoding/json" // Passes the byteValue to our struct.
	"fmt"           // Println etc.
	"io/ioutil"

	// Converts jsonFile into a byteValue, which is our byte array.
	"reflect"

	// Opens files and store it into jsonFile, in our memory
	"log"
	"net/http"
	"strconv" // Converts ints to strings etc.
)

type Recipe struct {
	Name               string `json:"Name"`
	ItemResultTargetID int    `json:"ItemResultTargetID"`
	ID                 int    `json:"ID"`
	Url                string `json:"Url"`
}

type AmountIngredient struct {
	//The outer values
	AmountIngredient0 int `json:"AmountIngredient0"`
	AmountIngredient1 int `json:"AmountIngredient1"`
	AmountIngredient2 int `json:"AmountIngredient2"`
	AmountIngredient3 int `json:"AmountIngredient3"`
	AmountIngredient4 int `json:"AmountIngredient4"`
	AmountIngredient5 int `json:"AmountIngredient5"`
	AmountIngredient6 int `json:"AmountIngredient6"`
	AmountIngredient7 int `json:"AmountIngredient7"`
	AmountIngredient8 int `json:"AmountIngredient8"`
	AmountIngredient9 int `json:"AmountIngredient9"`
}

// Outer Container
type IngredientRecipe struct {
	/*
		ItemIngredientRecipe0 []struct {
			CraftType struct {
				ID   int    `json:"ID"`
				Name string `json:"Name"`
			} `json:"CraftType"`
			ItemResult struct {
				ID int `json:"ID"`
			} `json:"ItemResult"`
		} `json:"ItemIngredientRecipe0"`
		This is how you obtain information from a json array of objects with properties*/
	ItemIngredientRecipe0 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe0"`
	ItemIngredientRecipe1 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe1"`
	ItemIngredientRecipe2 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe2"`
	ItemIngredientRecipe3 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe3"`
	ItemIngredientRecipe4 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe4"`
	ItemIngredientRecipe5 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe5"`
	ItemIngredientRecipe6 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe6"`
	ItemIngredientRecipe7 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe7"`
	ItemIngredientRecipe8 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe8"`
	ItemIngredientRecipe9 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe9"`
}

//Pass a struct item to ItemRecipe.
func GetRecipe(itemjson string) {
	// TODO: We can split the URL using categories, to get smaller payloads of JSON.
	// ABOUT TODO: You want to find an optimal amount of splitting, or just having one big payload (or one reduced payload would be ideal).
	//What this does, is open the file, and read it
	jsonFile, err := http.Get(itemjson)
	if err != nil {
		log.Fatalln(err)
	}
	// Takes the jsonFile.Body, and put it into memory as byteValue array.
	byteValue, err := ioutil.ReadAll(jsonFile.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Body.Close()

	// Recipe Struct information
	var item Recipe
	json.Unmarshal(byteValue, &item)
	//Print out our data to check.
	fmt.Println("Item:" + item.Name)
	fmt.Println("Recipe ID: " + strconv.Itoa(item.ID))
	fmt.Println("URL: " + item.Url)

	// Amount of Ingredients Information
	var amount AmountIngredient
	json.Unmarshal(byteValue, &amount)
	// We need to change the ugly AmountIngredient struct into an array.
	r_amount := reflect.ValueOf(amount)
	n_amount := r_amount.NumField()
	AmountIngredients := make([]string, n_amount)
	for i := 0; i < n_amount; i++ {
		AmountIngredients[i] = fmt.Sprintf(`%v`, r_amount.Field(i))
	}
	fmt.Println(AmountIngredients) // Output element = AmountIngredient

	//Ingredient Recipe Information
	var ingredients IngredientRecipe
	json.Unmarshal(byteValue, &ingredients)
	r_ingred := reflect.ValueOf(ingredients)
	n_ingred := r_ingred.NumField()
	IngredientRecipes := make([]string, n_ingred)
	for i := 0; i < n_ingred; i++ {
		IngredientRecipes[i] = fmt.Sprintf(`%v`, r_ingred.Field(i))
	}
	fmt.Println(IngredientRecipes) // Output element = [{Recipe for ingredient} {other recipe for ingredient}]

	//TODO: If every element in Ingredient Recipes=null, then that means it's a base item.
}
