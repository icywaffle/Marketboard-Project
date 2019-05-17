package main

import (
	"fmt"

	xivapi "./xivapi"
)

func main() {
	// Allows user to select what they want to do.
	for {
		var input int
		var userID int
		fmt.Printf("Input Integer: Search(1), Find Recipe(2), Find Item(3), Get  Item Prices(4):")
		fmt.Scanln(&input)

		switch input {
		case 1:
			fmt.Printf("Recipe ID:")
			fmt.Scan(&userID)
			//Gets the item recipe, and puts it into the database.
			xivapi.Getitem(userID)
		case 2:
			xivapi.GetItemPrices(14146)
		default:
			fmt.Println("Invalid Case Selected.")
			continue
		}
	}
}
