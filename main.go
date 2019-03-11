package main

import (
	"log"
	"time"

	getitems "./items"
)

func main() {

	start := time.Now()

	//We need a GET request, and pass this information to the GetRecipe function
	testitem := getitems.UrlAPI("Recipe", "33180")

	getitems.GetRecipe(testitem)

	elapsed := time.Since(start)
	log.Printf("Runtime: %s", elapsed)
}
