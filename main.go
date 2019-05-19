package main

import (
	"fmt"

	xivapi "./xivapi"
	database "./xivapi/database"
)

func main() {
	// Allows user to select what they want to do.
	for {
		var input int
		var userID int
		fmt.Printf("Input Integer: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			fmt.Printf("Recipe ID:")
			fmt.Scan(&userID)
			// If it is not in the database, access the API and put it into the database
			if !database.MongoFind("Recipes", "recipeid", userID) {
				xivapi.Getitem(userID)
			}

		case 2:
			fmt.Printf("Item ID:")
			fmt.Scan(&userID)
			database.TotalProfits(userID)
		default:
			fmt.Println("Invalid Case Selected.")
			continue
		}
	}
}
