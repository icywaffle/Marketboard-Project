package main

import (
	xivapi "./xivapi"
)

func main() {

	for {
		/*
			var recipeID int
			fmt.Printf("Input Recipe ID :")
			fmt.Scanln(&recipeID)
			// This function returns the current Marketboard Price, total material prices, and the total profit.
		*/
		xivapi.NetItemPrice(33180)
	}

}
