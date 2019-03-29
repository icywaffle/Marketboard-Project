// Converts Recipe Pages of json, to arrays.

package xivapi

import (
	// Passes the byteValue to our struct.
	"encoding/json"
	"fmt" // Println etc.
	"io/ioutil"

	// Converts jsonFile into a byteValue, which is our byte array.
	"reflect"

	// Opens files and store it into jsonFile, in our memory
	"log"
	"net/http"
	// Converts ints to strings etc.
)

type Recipe struct {
	Name               string `json:"Name"`
	ItemResultTargetID int    `json:"ItemResultTargetID"`
	ID                 int    `json:"ID"`
	Url                string `json:"Url"`
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

// This function allows us to pass these awful structs into this function and obtain a clean array.
func jsontoarray(anystruct interface{}) {
	r_any := reflect.ValueOf(anystruct)
	n_any := r_any.NumField()
	newarray := make([]string, n_any)
	for i := 0; i < n_any; i++ {
		newarray[i] = fmt.Sprintf(`%v`, r_any.Field(i))
	}
}

func GetRecipe(itemjson string) {
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

	var amount AmountIngredient
	json.Unmarshal(byteValue, &amount)
	jsontoarray(amount)
	fmt.Println(amount)

}
