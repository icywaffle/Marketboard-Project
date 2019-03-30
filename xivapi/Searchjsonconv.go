package xivapi

import (
	"encoding/json" // Passes the byteValue to our struct.
	"fmt"           // Println etc.
	"io/ioutil"

	// Converts jsonFile into a byteValue, which is our byte array.

	// Opens files and store it into jsonFile, in our memory
	"log"
	"net/http"
	"strconv" // Converts ints to strings etc.
)

// Converts Search Pages of json, to arrays.

type SearchPages struct {
	Pagination struct {
		Page           int `json:"Page"`
		PageTotal      int `json:"PageTotal"`
		ResultsTotal   int `json:"ResultsTotal"`
		ResultsPerPage int `json:"ResultsPerPage"`
	} `json:"Pagination"`
	Results []struct {
		ID      int    `json:"ID"`
		Name    string `json:"Name"`
		Icon    string `json:"Icon"`
		UrlType string `json:"UrlType"`
	} `json:"Results"`
}

//Pass a struct item to ItemRecipe.
func SearchItem(itemsearchjson string) {
	// TODO: We can split the URL using categories, to get smaller payloads of JSON.
	// ABOUT TODO: You want to find an optimal amount of splitting, or just having one big payload (or one reduced payload would be ideal).
	//What this does, is open the file, and read it
	jsonFile, err := http.Get(itemsearchjson)
	if err != nil {
		log.Fatalln(err)
	}
	// Takes the jsonFile.Body, and put it into memory as byteValue array.
	byteValue, err := ioutil.ReadAll(jsonFile.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Body.Close()

	// Pagination information
	var pages SearchPages
	json.Unmarshal(byteValue, &pages)
	//Print out our data to check.
	fmt.Println("Page:" + strconv.Itoa(pages.Pagination.Page))
	fmt.Println("Total Pages: " + strconv.Itoa(pages.Pagination.PageTotal))
	fmt.Println("Total Results: " + strconv.Itoa(pages.Pagination.ResultsTotal))
	fmt.Println("Results per Page: " + strconv.Itoa(pages.Pagination.ResultsPerPage))

	// Search Results Information
	// Arrayed in order to flexibily add items to the search results
	fmt.Println(pages.Results) // Output Array of information.

}
