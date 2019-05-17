package main

import (
	"strconv"

	xivapi "./xivapi"
)

func main() {
	xivapi.GetItemPrices(xivapi.UrlPrices(strconv.Itoa(33180)), 33180)
	/*
		// Allows user to select what they want to do.
		for {
			var input int
			var userchoice string
			var userID int
			fmt.Printf("Input Integer: ")
			fmt.Scanln(&input)

			switch input {
			case 1:
				fmt.Printf("Recipe ID:")
				fmt.Scan(&userID)
				userchoice = "recipe"
				//Gets the item recipe, and puts it into the database.
				xivapi.Getitem(xivapi.UrlItemRecipe(userchoice, strconv.Itoa(userID)), userchoice)
			case 2:
				fmt.Printf("Item ID:")
				fmt.Scan(&userID)
				userchoice = "item"
				xivapi.Getitem(xivapi.UrlItemRecipe(userchoice, strconv.Itoa(userID)), userchoice)
			case 3:
				database.MongoFind("recipeid", 33180)
			case 4:
				xivapi.GetItemPrices(xivapi.UrlPrices(strconv.Itoa(33180)), 33180)
			default:
				fmt.Println("Invalid Case Selected.")
				continue
			}
		}*/

}
