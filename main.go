package main

import (
	"bufio"
	"fmt"
	"os"

	xivapi "./xivapi"
)

func search() {
	fmt.Printf("XIVAPI Search:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	usersearchinput := scanner.Text()
	//Searches the item, and put the json into an array.
	xivapi.SearchItem(xivapi.UrlSearch(usersearchinput))
}
func choose() {
	//After finding the results, and user chooses put the item ID into here.
	fmt.Printf("ItemID:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	useriteminput := scanner.Text()
	fmt.Printf("Item/Recipe:")
	scanner.Scan()
	userchoiceinput := scanner.Text()
	testitem := xivapi.UrlRecipe(userchoiceinput, useriteminput)

	//We need to use the GET request, and put JSON data into the database.
	xivapi.GetRecipe(testitem)
}

func main() {
	//We need to get the URL
	for {
		var input int
		n, err := fmt.Scanln(&input)
		// Force choose a positive number
		if n < 1 || err != nil {
			fmt.Println("invalid input")
			return
		}
		// If you need to re-search, use case 1.
		// If you have the right itemID and Recipe, move onto case 2.
		switch input {
		case 1:
			search()
		case 2:
			choose()
		default:
			os.Exit(2000) // ERROR 2000 : End program with no input.
		}
	}

	// Current Status: Able to search, run, and obtain the Amount of ingredients, and ingredients of items
}
