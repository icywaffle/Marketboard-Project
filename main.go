package main

import (
	"fmt"
	"strconv"

	database "./database"
	xivapi "./xivapi"
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
func choose(userchoiceinput string) string {
	//After finding the results, and user chooses put the item ID into here.
	fmt.Printf("ID:")
	var userID int
	fmt.Scan(&userID)
	fmt.Println(xivapi.UrlItemRecipe(userchoiceinput, strconv.Itoa(userID)))
	return xivapi.UrlItemRecipe(userchoiceinput, strconv.Itoa(userID))
}

func getprices() string {
	fmt.Printf("ID:")
	var userID int
	fmt.Scan(&userID)
	return xivapi.UrlPrices(strconv.Itoa(userID))

}

// Checks the database first before creating the url to request.
func mongoHandler() {
	//Ask user for the itemID, and check the database if it exists.
	fmt.Printf("ID:")
	var userID int
	fmt.Scan(&userID)
	//If it exists, then this function should automatically use the information
	//Else if it does not exist, then the MongoFind function should automatically call the WebUrl Requests.
	// And mongoinsert into the database.
	database.MongoFind("itemid", userID)

}

func main() {
	// Allows user to select what they want to do.
	for {
		var input int
		var userchoice string
		fmt.Printf("Input Integer: Search(1), Find Recipe(2), Find Item(3), MongoDatabase(4), getprices(5):")
		fmt.Scanln(&input)

		switch input {
		case 1:
			search()
		case 2:
			userchoice = "recipe"
			xivapi.Getitem(choose(userchoice), userchoice)
		case 3:
			userchoice = "item"
			choose(userchoice)
		case 4:
			mongoHandler()
		case 5:
			xivapi.GetItemPrices(getprices())
		default:
			fmt.Println("Invalid Case Selected.")
			continue
		}
	}
}
