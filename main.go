package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	xivapi "./xivapi"
)

func main() {

	start := time.Now()

	fmt.Printf("XIVAPI Search:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	usersearchinput := scanner.Text()
	//Searches the item, and put the json into an array.
	xivapi.SearchItem(xivapi.UrlSearch(usersearchinput))

	//After finding the results, and user chooses put the item ID into here.
	fmt.Printf("ItemID:")
	scanner.Scan()
	useriteminput := scanner.Text()
	fmt.Printf("Item/Recipe:")
	scanner.Scan()
	userchoiceinput := scanner.Text()
	//We need to get the URL
	testitem := xivapi.UrlRecipe(userchoiceinput, useriteminput)
	//We need to use the GET request, and put JSON data into the database.
	xivapi.GetRecipe(testitem)

	elapsed := time.Since(start)
	log.Printf("Runtime: %s", elapsed)
	// Current Status: Able to search, run, and obtain the Amount of ingredients, and ingredients of items
}
