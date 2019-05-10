package xivapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	database "../database"
)

func GetItemPrices(websiteurl string) {

	// MAX Rate limit is 20 Req/s -> 0.05s/Req, but safer to use 15req/s -> 0.06s/req
	time.Sleep(100 * time.Millisecond)
	// This ensures that when this function is called, it does not exceed the rate limit.
	// TODO: Use a channel to rate limit instead to allow multiple users to use this.

	//Get request to create the bytevalue to unload into the struct
	jsonFile, err := http.Get(websiteurl)
	if err != nil {
		log.Fatalln(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Body.Close()

	var prices database.Prices
	json.Unmarshal(byteValue, &prices)

	database.MongoInsertPrices(prices)
}
