package main

import (
	"fmt"
	"strconv"

	xivapi "./xivapi"
	database "./xivapi/database"
)

// Uses the Web Appending Function to create the url to request.
func search() {
	// Prompt for Search Input
	fmt.Printf("XIVAPI Search:")
	fmt.Printf("ID:")
	var usersearchinput string
	fmt.Scan(&usersearchinput)

	// Pass the input to the Web Appending Function
	urlxivapisearch := xivapi.UrlSearch(usersearchinput)

	// Then use the xivapi to search for results
	xivapi.SearchItem(urlxivapisearch)
}

func main() {
	// Allows user to select what they want to do.
	for {
		var input int
		var userchoice string
		var userID int
		fmt.Printf("Input Integer: Search(1), Find Recipe(2), Find Item(3), Get  Item Prices(4):")
		fmt.Scanln(&input)

		switch input {
		case 1:
			search()
		case 2:
			fmt.Printf("Recipe ID:")
			fmt.Scan(&userID)
			userchoice = "recipe"
			//Gets the item recipe, and puts it into the database.
			xivapi.Getitem(xivapi.UrlItemRecipe(userchoice, strconv.Itoa(userID)), userchoice)
		case 3:
			fmt.Printf("Item ID:")
			fmt.Scan(&userID)
			userchoice = "item"
			xivapi.Getitem(xivapi.UrlItemRecipe(userchoice, strconv.Itoa(userID)), userchoice)
		case 4:
			database.MongoFind("recipeid", 33180)
		case 5:
			xivapi.GetItemPrices(xivapi.UrlPrices(strconv.Itoa(14146)), 14146)
		default:
			fmt.Println("Invalid Case Selected.")
			continue
		}
	}
}
