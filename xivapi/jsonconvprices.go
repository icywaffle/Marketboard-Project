package xivapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/////Price Struct//////
type Prices struct {
	Sargatanas struct {
		History []struct {
			Added        int  `json:"Added"` // Time is in Unix epoch time
			IsHQ         bool `json:"IsHQ"`
			PricePerUnit int  `json:"PricePerUnit"`
			PriceTotal   int  `json:"PriceTotal"`
			PurchaseDate int  `json:"PurchaseDate"`
			Quantity     int  `json:"Quantity"`
		} `json:"History"`
	} `json:"Adamantoise"`
}

func GetItemPrices(itemjson string, userchoiceinput string) {

	// MAX Rate limit is 20 Req/s -> 0.05s/Req, but safer to use 15req/s -> 0.06s/req
	time.Sleep(100 * time.Millisecond)
	// This ensures that when this function is called, it does not exceed the rate limit.
	// TODO: Use a channel to rate limit instead to allow multiple users to use this.

	//What this does, is open the file, and read it
	//TODO : At this point, we need an if statement to check if we have the data or not.
	// If we do, then there's no need to http.Get
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

	var prices Prices
	json.Unmarshal(byteValue, &prices)

	fmt.Println(prices.Sargatanas)
}
