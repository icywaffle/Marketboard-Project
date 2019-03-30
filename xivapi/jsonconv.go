package xivapi

import (
	// Passes the byteValue to our struct.
	"encoding/json"
	"fmt" // Println etc.
	"io/ioutil"
	"strings"

	// Converts jsonFile into a byteValue, which is our byte array.
	"reflect"

	// Opens files and store it into jsonFile, in our memory
	"log"
	"net/http"
	// Converts ints to strings etc.
)

const SIZEOF_INT32 = 4 // bytes

// Converts Recipe Pages of json, to arrays.

type Recipe struct {
	Name               string `json:"Name"`
	ItemResultTargetID int    `json:"ItemResultTargetID"`
	ID                 int    `json:"ID"`
	CraftType          struct {
		ID int `json:"ID"`
	} `json:"CraftType"`
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
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe0"`
	ItemIngredientRecipe1 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe1"`
	ItemIngredientRecipe2 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe2"`
	ItemIngredientRecipe3 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe3"`
	ItemIngredientRecipe4 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe4"`
	ItemIngredientRecipe5 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe5"`
	ItemIngredientRecipe6 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe6"`
	ItemIngredientRecipe7 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe7"`
	ItemIngredientRecipe8 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe8"`
	ItemIngredientRecipe9 []struct {
		CraftTypeTargetID  int `json:"CraftTypeTargetID"`
		ItemResultTargetID int `json:"ItemResultTargetID"`
	} `json:"ItemIngredientRecipe9"`
}

type Item struct {
	Name             string `json:"Name"`
	ID               int    `json:"ID"`
	GameContentLinks struct {
		Recipe struct {
			ItemResult []int `json:"ItemResult"`
		} `json:"Recipe"`
	} `json:"GameContentLinks"`
}

// This function allows us to pass these awful structs into this function and obtain a clean slice.
func Jsontoslice(anystruct interface{}, slicename []string) {
	r_any := reflect.ValueOf(anystruct)
	n_any := r_any.NumField()
	slicename = slicename[:n_any] //Resize the slice to fit the number of fields.
	for i := 0; i < n_any; i++ {
		slicename[i] = fmt.Sprintf(`%v`, r_any.Field(i))
	}
	// Unfortunately, array elements are strings instead of ints.
	// Don't know if it can put ints into the slice element instead.
}

func Get(itemjson string, userchoiceinput string) {
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

	var recipeinfo Recipe
	json.Unmarshal(byteValue, &recipeinfo)
	// Can directly access children of structs.

	if userchoiceinput == "recipe" {
		var amount AmountIngredient
		json.Unmarshal(byteValue, &amount)
		amountslice := make([]string, 10) // Initializes a Slice
		Jsontoslice(amount, amountslice)  // <- Accesses Slice Elements.
		fmt.Println(amountslice)          // Prints out the slice.

		var matitemID ItemIngredient
		json.Unmarshal(byteValue, &matitemID)
		matitemIDslice := make([]string, 10)
		Jsontoslice(matitemID, matitemIDslice)
		fmt.Println(matitemIDslice)

		var matrecipeID IngredientRecipe
		json.Unmarshal(byteValue, &matrecipeID)
		matrecipeIDslice := make([]string, 10)
		Jsontoslice(matrecipeID, matrecipeIDslice)
		fmt.Println(matrecipeIDslice)

		// Check if it's ingredient is a base item.
		// If the length of the element is > 2, it must have recipes inside of it.
		// Else, it's a base ingredient and we don't need any more information.
		n := len(matrecipeIDslice)
		for i := 0; i < n; i++ {
			if len(matrecipeIDslice[i]) > 2 {
				// An ingredient has a recipe, we pass the ID, back into the function and redo.
				itemurl := UrlRecipe("item", strings.Trim(fmt.Sprintf(matitemIDslice[i]), "[]"))
				Get(itemurl, "item")
				fmt.Println(itemurl)
			}
		}
	} else if userchoiceinput == "item" {
		var items Item
		json.Unmarshal(byteValue, &items)
		//fmt.Println(items.Name, items.ID)
		fmt.Println(items.GameContentLinks.Recipe.ItemResult)
	}
}
