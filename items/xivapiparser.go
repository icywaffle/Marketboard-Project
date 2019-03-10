package main

import (
	"fmt"
	// We need to import encoding/json package
	// in order to parse with structs
	"encoding/json"

	"io/ioutil"

	// In order to open files, we need to use os package.
	"os"

	//In order to convert int to str
	"strconv"
)

// We need to put the byteValue into an array
// An ItemRecipe will look like
type ItemRecipe struct {
	//Name of the Item
	Name string `json:"Name"`
	//Item ID of the Item
	ItemResultTargetID int `json:"ItemResultTargetID"`
	//Recipe of the Item
	ID int `json:"ID"`
	//Url of What you're looking at
	Url string `json:"Url"`
	//The 9 Ingredient Objects
	// Why they didn't put it as an array of 9 objects... I don't know.
	ItemIngredient0 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient0"`
	ItemIngredient1 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient1"`
	ItemIngredient2 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient2"`
	ItemIngredient3 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient3"`
	ItemIngredient4 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient4"`
	ItemIngredient5 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient5"`
	ItemIngredient6 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient6"`
	ItemIngredient7 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient7"`
	ItemIngredient8 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient8"`
	ItemIngredient9 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient9"`
	//The 9 Amount Ingredients
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

func main() {

	// Open the json file
	jsonFile, err := os.Open("Highmythriteingot.json")

	// Os.Open will give two values, the jsonFile, and the error.
	// If it returns an error, it will print it.
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Success. Opened json.")

	// We need to keep the jsonFile Open.
	defer jsonFile.Close()

	// This should read into our memory.
	// We need to convert our byte array
	// using ioutil.ReadAll
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Temporary Initialization. We will pass a struct through the function later.
	var item ItemRecipe

	//We now just unmarshal the byteArray into the struct
	json.Unmarshal(byteValue, &item)

	//Iterate every user array element,
	// print out Item Name, Recipe ID, URL, and The Ingredients

	fmt.Println("Item:" + item.Name)
	// Changes int to a string. strconv.Itoa.
	fmt.Println("Recipe ID: " + strconv.Itoa(item.ID))
	fmt.Println("URL: " + item.Url)
	/*
		fmt.Println("First Ingredient: "+ strconv.Itoa(item.ItemIngredient0.ID))
		fmt.Println("First Ingredient: "+ strconv.Itoa(item.ItemIngredient1.ID))
		fmt.Println("First Ingredient: "+ strconv.Itoa(item.ItemIngredient2.ID))
		fmt.Println("First Ingredient: "+ strconv.Itoa(item.ItemIngredient3.ID))
		fmt.Println("First Ingredient: "+ strconv.Itoa(item.ItemIngredient4.ID))
		fmt.Println("First Ingredient: "+ strconv.Itoa(item.ItemIngredient5.ID))
		fmt.Println("First Ingredient: "+ strconv.Itoa(item.ItemIngredient6.ID))
		fmt.Println("First Ingredient: "+ strconv.Itoa(item.ItemIngredient7.ID))
		fmt.Println("First Ingredient: "+ strconv.Itoa(item.ItemIngredient8.ID))
		fmt.Println("First Ingredient: "+ strconv.Itoa(item.ItemIngredient9.ID))
		fmt.Println("Ingred0 Amount: "+ strconv.Itoa(item.AmountIngredient0))
		fmt.Println("Ingred1 Amount: "+ strconv.Itoa(item.AmountIngredient1))
		fmt.Println("Ingred2 Amount: "+ strconv.Itoa(item.AmountIngredient2))
		fmt.Println("Ingred3 Amount: "+ strconv.Itoa(item.AmountIngredient3))
		fmt.Println("Ingred4 Amount: "+ strconv.Itoa(item.AmountIngredient4))
		fmt.Println("Ingred5 Amount: "+ strconv.Itoa(item.AmountIngredient5))
		fmt.Println("Ingred6 Amount: "+ strconv.Itoa(item.AmountIngredient6))
		fmt.Println("Ingred7 Amount: "+ strconv.Itoa(item.AmountIngredient7))
		fmt.Println("Ingred8 Amount: "+ strconv.Itoa(item.AmountIngredient8))
		fmt.Println("Ingred9 Amount: "+ strconv.Itoa(item.AmountIngredient9))
	*/
	//TODO: Put this information into the database!
}
