package main

import (
	"log"
	"time"

	getitems "./items"
)

func main() {

	start := time.Now()

	//We need to get the URL
	testitem := getitems.UrlAPI("Recipe", "33180")
	//We need to use the GET request, and put JSON data into the database.
	getitems.GetRecipe(testitem)

	elapsed := time.Since(start)
	log.Printf("Runtime: %s", elapsed)
}
