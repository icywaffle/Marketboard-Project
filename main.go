package main

import (
	"bufio"
	"fmt"
	"os"

	database "./database"
	xivapi "./xivapi"
)

func search() {
	// Prompt for Search Input
	fmt.Printf("XIVAPI Search:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	usersearchinput := scanner.Text()
	//Takes user input, convert it into the url, and searches through the json.
	xivapi.SearchItem(xivapi.UrlSearch(usersearchinput))
	// Output: Array of objects that have the itemID and item icon, and item/recipe type.

	// Items are meant to be information for
	/*
			  "GameContentLinks": {
		        "Recipe": {
		            "ItemResult": [
		                33180
		            ]
		        }
			},*/
	// To be able to find Item Recipes. If null, then it's a base item.
}
func chooseRecipe() {
	//After finding the results, and user chooses put the item ID into here.
	fmt.Printf("RecipeID:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	useriteminput := scanner.Text()
	//And identify whether this id is a recipe or item id
	userchoiceinput := "recipe"
	// Convert input into a url, then output the url string
	itemurl := xivapi.UrlRecipe(userchoiceinput, useriteminput)
	//We need to use the GET request on the url, and put JSON data into the database.
	xivapi.Get(itemurl, userchoiceinput)
}
func chooseItem() {
	//After finding the results, and user chooses put the item ID into here.
	fmt.Printf("ItemID:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	useriteminput := scanner.Text()
	userchoiceinput := "item"
	// Convert input into a url, then output the url string
	itemurl := xivapi.UrlRecipe(userchoiceinput, useriteminput)
	//We need to use the GET request on the url, and put JSON data into the database.
	xivapi.Get(itemurl, userchoiceinput)
}

func main() {
	//We need to get the URL
	for {
		var input int
		fmt.Printf("Input Case 1,2,3:")
		n, err := fmt.Scanln(&input)
		// Force choose a positive number
		if n < 1 || err != nil {
			fmt.Println("invalid input")
			os.Exit(2)
		}
		// If you need to re-search, use case 1.
		// If you have the right itemID and Recipe, move onto case 2.
		switch input {
		case 1:
			search()
		case 2:
			chooseRecipe()
		case 3:
			chooseItem()
		case 4:
			os.Exit(1)
		case 5:
			database.MongoHandler()
		default:
			fmt.Println("Invalid Case Selected.")
			continue
		}
	}
}
