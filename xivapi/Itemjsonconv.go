package xivapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Passes the byteValue to our struct.
// Println etc.
// Converts jsonFile into a byteValue, which is our byte array.

// Opens files and store it into jsonFile, in our memory

// Converts ints to strings etc.

type Item struct {
	Name string `json:"Name"`
	ID   int    `json:"ID"`
}
type Link struct {
	GameContentLinks struct {
		Recipe struct {
			ItemResult []int `json:"ItemResult"`
		} `json:"Recipe"`
	} `json:"GameContentLinks"`
}

func GetItem(itemjson string) {
	//What this does, is open the file, and read it
	jsonFile, err := http.Get(itemjson)
	if err != nil {
		log.Fatalln(err)
	}
	// Takes the jsonFile.Body, and put it into memory as byteValue array.
	byteValue, err := ioutil.ReadAll(jsonFile.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Body.Close()

	var items Item
	json.Unmarshal(byteValue, &items)
	Jsontoarray(items)
	fmt.Println(items)
	var links Link
	json.Unmarshal(byteValue, &links)
	Jsontoarray(links)
	fmt.Println(links)

}
