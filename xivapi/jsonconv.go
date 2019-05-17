package xivapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	database "./database"
)

const SIZEOF_INT32 = 4 // bytes

// Converts Recipe Pages of json, to arrays.

/////////////////Recipe Struct Here//////////////////////////

type AmountIngredient struct {
	//The outer values
	AmountIngredient0 int `json:"AmountIngredient0"`
	AmountIngredient1 int `json:"AmountIngredient1"`
	AmountIngredient2 int `json:"AmountIngredient2"`
	AmountIngredient3 int `json:"AmountIngredient3"`
	AmountIngredient4 int `json:"AmountIngredient4"`
	AmountIngredient5 int `json:"AmountIngredient5"`
	AmountIngredient6 int `json:"AmountIngredient6"`
	AmountIngredient7 int `json:"AmountIngredient7"`
	AmountIngredient8 int `json:"AmountIngredient8"`
	AmountIngredient9 int `json:"AmountIngredient9"`
}

type ItemIngredient struct {
	ItemIngredient0TargetID int `json:"ItemIngredient0TargetID"`
	ItemIngredient1TargetID int `json:"ItemIngredient1TargetID"`
	ItemIngredient2TargetID int `json:"ItemIngredient2TargetID"`
	ItemIngredient3TargetID int `json:"ItemIngredient3TargetID"`
	ItemIngredient4TargetID int `json:"ItemIngredient4TargetID"`
	ItemIngredient5TargetID int `json:"ItemIngredient5TargetID"`
	ItemIngredient6TargetID int `json:"ItemIngredient6TargetID"`
	ItemIngredient7TargetID int `json:"ItemIngredient7TargetID"`
	ItemIngredient8TargetID int `json:"ItemIngredient8TargetID"`
	ItemIngredient9TargetID int `json:"ItemIngredient9TargetID"`
}

// Outer Container
type IngredientRecipe struct {
	ItemIngredientRecipe0 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe0"`
	ItemIngredientRecipe1 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe1"`
	ItemIngredientRecipe2 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe2"`
	ItemIngredientRecipe3 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe3"`
	ItemIngredientRecipe4 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe4"`
	ItemIngredientRecipe5 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe5"`
	ItemIngredientRecipe6 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe6"`
	ItemIngredientRecipe7 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe7"`
	ItemIngredientRecipe8 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe8"`
	ItemIngredientRecipe9 []struct {
		ID int `json:"ID"`
	} `json:"ItemIngredientRecipe9"`
}

//////////////////Item Struct Here////////////////////////////
type Item struct {
	Name string `json:"Name"`
	ID   int    `json:"ID"`
	Icon string `json:"Icon"`
}

func Getitem(websiteurl string, userchoiceinput string) {

	//If statement here. Use MongoFind.
	// If we have it, we just put use profit.go
	// If we dont have it, we just parse the information into Profit.go in a readable form.

	// MAX Rate limit is 20 Req/s -> 0.05s/Req, but safer to use 15req/s -> 0.06s/req
	time.Sleep(100 * time.Millisecond)
	// This ensures that when this function is called, it does not exceed the rate limit.
	// TODO: Use a channel to rate limit instead to allow multiple users to use this.

	//What this does, is open the file, and read it
	//TODO : At this point, we need an if statement to check if we have the data or not.
	// If we do, then there's no need to http.Get
	jsonFile, err := http.Get(websiteurl)
	if err != nil {
		log.Fatalln(err)
	}
	// Takes the jsonFile.Body, and put it into memory as byteValue array.
	byteValue, err := ioutil.ReadAll(jsonFile.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Body.Close()

	var recipes database.Recipes
	json.Unmarshal(byteValue, &recipes)

	if userchoiceinput == "recipe" {
		var amount AmountIngredient
		json.Unmarshal(byteValue, &amount)
		// Passing the struct into the array instead to obtain a cleaner slice.
		amountslice := []int{amount.AmountIngredient0,
			amount.AmountIngredient1,
			amount.AmountIngredient2,
			amount.AmountIngredient3,
			amount.AmountIngredient4,
			amount.AmountIngredient5,
			amount.AmountIngredient6,
			amount.AmountIngredient7,
			amount.AmountIngredient8,
			amount.AmountIngredient9}

		var matitemID ItemIngredient
		json.Unmarshal(byteValue, &matitemID)
		matitemIDslice := []int{matitemID.ItemIngredient0TargetID,
			matitemID.ItemIngredient1TargetID,
			matitemID.ItemIngredient2TargetID,
			matitemID.ItemIngredient3TargetID,
			matitemID.ItemIngredient4TargetID,
			matitemID.ItemIngredient5TargetID,
			matitemID.ItemIngredient6TargetID,
			matitemID.ItemIngredient7TargetID,
			matitemID.ItemIngredient8TargetID,
			matitemID.ItemIngredient9TargetID}

		//Pass all this information into the database
		database.MongoInsertRecipe(recipes, matitemIDslice, amountslice)

		//Finally, we need to go through each recipe that is possible.
		matrecipeIDslice := matrecipeslice(byteValue)
		for i := 0; i < len(matrecipeIDslice); i++ {
			for j := 0; j < len(matrecipeIDslice[i]); j++ {
				Getitem(UrlItemRecipe("recipe", strconv.Itoa(matrecipeIDslice[i][j])), "recipe")
			}
		}

		//This is for requesting information about an item.
	} else if userchoiceinput == "item" {
		var items Item
		json.Unmarshal(byteValue, &items)
		// We need to iterate over the elements of the array
		fmt.Println(items.ID, items.Icon, items.Name)
	} // TODO: Store these array information into a caching layer, which we can call instead of calling the server for the same pages over and over etc.
}

func GetItemPrices(websiteurl string, userID int) {

	//If statement here. Use MongoFind

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

	//database.MongoInsertPrices(prices, userID)

	//Finally, we need to go through each recipe that is possible.
	matrecipeIDslice := matrecipeslice(byteValue)
	for i := 0; i < len(matrecipeIDslice); i++ {
		for j := 0; j < len(matrecipeIDslice[i]); j++ {
			GetItemPrices(UrlItemRecipe("recipe", strconv.Itoa(matrecipeIDslice[i][j])), matrecipeIDslice[i][j])
		}
	}
}

func matrecipeslice(byteValue []byte) [][]int {
	fmt.Println(byteValue)
	// We need this information in order to go through every single possible recipe that can make this item.
	var matrecipeID IngredientRecipe
	json.Unmarshal(byteValue, &matrecipeID)
	matrecipeIDslice := make([][]int, 10)

	//No choice but to unravel for each element, the possible Material Ingredient Recipe IDs 10 times.
	// There is variable length for different elements.
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe0); i++ {
		matrecipeIDslice[0] = append(matrecipeIDslice[0], matrecipeID.ItemIngredientRecipe0[i].ID)
	}
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe1); i++ {
		matrecipeIDslice[1] = append(matrecipeIDslice[1], matrecipeID.ItemIngredientRecipe1[i].ID)
	}
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe2); i++ {
		matrecipeIDslice[2] = append(matrecipeIDslice[2], matrecipeID.ItemIngredientRecipe2[i].ID)
	}
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe3); i++ {
		matrecipeIDslice[3] = append(matrecipeIDslice[3], matrecipeID.ItemIngredientRecipe3[i].ID)
	}
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe4); i++ {
		matrecipeIDslice[4] = append(matrecipeIDslice[4], matrecipeID.ItemIngredientRecipe4[i].ID)
	}
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe5); i++ {
		matrecipeIDslice[5] = append(matrecipeIDslice[5], matrecipeID.ItemIngredientRecipe5[i].ID)
	}
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe6); i++ {
		matrecipeIDslice[6] = append(matrecipeIDslice[6], matrecipeID.ItemIngredientRecipe6[i].ID)
	}
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe7); i++ {
		matrecipeIDslice[7] = append(matrecipeIDslice[7], matrecipeID.ItemIngredientRecipe7[i].ID)
	}
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe8); i++ {
		matrecipeIDslice[8] = append(matrecipeIDslice[8], matrecipeID.ItemIngredientRecipe8[i].ID)
	}
	for i := 0; i < len(matrecipeID.ItemIngredientRecipe9); i++ {
		matrecipeIDslice[9] = append(matrecipeIDslice[9], matrecipeID.ItemIngredientRecipe9[i].ID)
	}
	fmt.Println(matrecipeID.ItemIngredientRecipe0)
	return matrecipeIDslice
}
