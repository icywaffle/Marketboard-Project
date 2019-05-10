package xivapi

import (
	// Passes the byteValue to our struct.
	"encoding/json"
	"fmt" // Println etc.
	"io/ioutil"
	"strconv"
	"time"

	// Converts jsonFile into a byteValue, which is our byte array.

	// Opens files and store it into jsonFile, in our memory
	"log"
	"net/http"
	// Converts ints to strings etc.
)

const SIZEOF_INT32 = 4 // bytes

// Converts Recipe Pages of json, to arrays.

/////////////////Recipe Struct Here//////////////////////////
type Recipe struct {
	Name               string `json:"Name"`
	ItemResultTargetID int    `json:"ItemResultTargetID"` // This is the Item ID
	ID                 int    `json:"ID"`                 // This is the recipeID
	CraftTypeTargetID  int    `json:"CraftTypeTargetID"`
}

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

func Getitem(itemjson string, userchoiceinput string) {
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

	var recipeinfo Recipe
	json.Unmarshal(byteValue, &recipeinfo)

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

		//  This slice is meant to search for every item recipe that we want.
		var matrecipeID IngredientRecipe
		json.Unmarshal(byteValue, &matrecipeID)
		matrecipeIDslice := make([][]int, 10)

		//No choice but to unravel for each element, the possible Material Ingredient Recipe IDs 10 times.
		// There is variable length for different elements.
		for i := 0; i < len(matrecipeID.ItemIngredientRecipe0); i++ {
			// Add to each element, the matrecipeIDs that are possible for one item
			matrecipeIDslice[0] = append(matrecipeIDslice[0], matrecipeID.ItemIngredientRecipe0[i].ID)
		}
		for i := 0; i < len(matrecipeID.ItemIngredientRecipe1); i++ {
			// Add to each element, the matrecipeIDs that are possible for one item
			matrecipeIDslice[1] = append(matrecipeIDslice[1], matrecipeID.ItemIngredientRecipe1[i].ID)
		}
		for i := 0; i < len(matrecipeID.ItemIngredientRecipe2); i++ {
			// Add to each element, the matrecipeIDs that are possible for one item
			matrecipeIDslice[2] = append(matrecipeIDslice[2], matrecipeID.ItemIngredientRecipe2[i].ID)
		}
		for i := 0; i < len(matrecipeID.ItemIngredientRecipe3); i++ {
			// Add to each element, the matrecipeIDs that are possible for one item
			matrecipeIDslice[3] = append(matrecipeIDslice[3], matrecipeID.ItemIngredientRecipe3[i].ID)
		}
		for i := 0; i < len(matrecipeID.ItemIngredientRecipe4); i++ {
			// Add to each element, the matrecipeIDs that are possible for one item
			matrecipeIDslice[4] = append(matrecipeIDslice[4], matrecipeID.ItemIngredientRecipe4[i].ID)
		}
		for i := 0; i < len(matrecipeID.ItemIngredientRecipe5); i++ {
			// Add to each element, the matrecipeIDs that are possible for one item
			matrecipeIDslice[5] = append(matrecipeIDslice[5], matrecipeID.ItemIngredientRecipe5[i].ID)
		}
		for i := 0; i < len(matrecipeID.ItemIngredientRecipe6); i++ {
			// Add to each element, the matrecipeIDs that are possible for one item
			matrecipeIDslice[6] = append(matrecipeIDslice[6], matrecipeID.ItemIngredientRecipe6[i].ID)
		}
		for i := 0; i < len(matrecipeID.ItemIngredientRecipe7); i++ {
			// Add to each element, the matrecipeIDs that are possible for one item
			matrecipeIDslice[7] = append(matrecipeIDslice[7], matrecipeID.ItemIngredientRecipe7[i].ID)
		}
		for i := 0; i < len(matrecipeID.ItemIngredientRecipe8); i++ {
			// Add to each element, the matrecipeIDs that are possible for one item
			matrecipeIDslice[8] = append(matrecipeIDslice[8], matrecipeID.ItemIngredientRecipe8[i].ID)
		}
		for i := 0; i < len(matrecipeID.ItemIngredientRecipe9); i++ {
			// Add to each element, the matrecipeIDs that are possible for one item
			matrecipeIDslice[9] = append(matrecipeIDslice[9], matrecipeID.ItemIngredientRecipe9[i].ID)
		}

		//Pass all this information into the database
		//database.MongoInsertRecipe(recipeinfo.Name, recipeinfo.ID, recipeinfo.ItemResultTargetID, recipeinfo.CraftTypeTargetID, matitemIDslice, amountslice)
		fmt.Println(amountslice, matitemIDslice)
		//Finally, we need to go through each recipe that is possible.
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
