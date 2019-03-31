package xivapi

import (
	// Passes the byteValue to our struct.
	"encoding/json"
	"fmt" // Println etc.
	"io/ioutil"
	"strings"
	"time"

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

type Item struct {
	Name string `json:"Name"`
	ID   int    `json:"ID"`
	Icon string `json:"Icon"`
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
	// MAX Rate limit is 20 Req/s -> 0.066s/Req
	time.Sleep(60 * time.Millisecond)
	// TODO: Use a channel to rate limit instead

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

	// Then unmarshals the byteValues into the struct.
	var recipeinfo Recipe
	json.Unmarshal(byteValue, &recipeinfo)

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

		// Check if it's ingredient is a base item.
		// If the length of the element is > 2, it must have recipes inside of it.
		// Else, it's a base ingredient and we don't need any more information.
		// Ex: matrecipeIDslice = [[{31482} {31843}] [{31486}] [{31484}] [] [] [] [] [] [] []]
		// Empty arrays have length of 2.
		n := len(matrecipeIDslice)
		for i := 0; i < n; i++ {
			if len(matrecipeIDslice[i]) > 2 {
				// An ingredient has a recipe, we pass the ID, back into the function and redo.
				fmt.Println(matitemIDslice[i], matrecipeIDslice[i], i)
				match(matrecipeIDslice[i])
			}
		}
		// If all we have is itemIDs, we need to search for the possible RecipeID.
	} else if userchoiceinput == "item" {
		var items Item
		json.Unmarshal(byteValue, &items)
		// We need to iterate over the elements of the array
		fmt.Println(items.ID, items.Icon, items.Name)
	} // TODO: Store these array information into a caching layer, which we can call instead of calling the server for the same pages over and over etc.
}

func match(input string) {
	for {
		starting := strings.Index(input, "{") // Will return the indext of the first instance.
		ending := strings.Index(input[starting:], "}")
		fmt.Println(starting, ending)

		if starting >= 0 {
			if ending >= 0 {
				result := input[starting+1 : ending+1]
				fmt.Println(result)
				if len(input) != 9 { // Length of input = 9 , means that there's only one ID!
					input = input[ending+2:]
					fmt.Println(input)
				} else {
					break
				}
			}
		}
	}

	/*
		for i := 0; i < 15; i++ {
			if starting >= 0 { // If we find start index, then iterate
				if ending >= 0 { //  If we find end index, then finish
					result[i] = input[starting+1 : ending+1]
					// cut the string off, and do it again.

				}
			} else {
				result[i] = ""
			}
		}
		// To end, we are unable to find any more -> {
		// result[i] has a set length.
		return result
	*/
}
