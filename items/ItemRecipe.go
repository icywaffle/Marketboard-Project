package items

import (
	"encoding/json" // Passes the byteValue to our struct.
	"fmt"           // Println etc.
	"io/ioutil"

	// Converts jsonFile into a byteValue, which is our byte array.
	"reflect"

	// Opens files and store it into jsonFile, in our memory
	"log"
	"net/http"
	"strconv" // Converts ints to strings etc.
)

type Recipe struct {
	Name               string `json:"Name"`
	ItemResultTargetID int    `json:"ItemResultTargetID"`
	ID                 int    `json:"ID"`
	Url                string `json:"Url"`
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

// Outer Container
type IngredientRecipe struct {
	//An object, with arrays of objects
	ItemIngredientRecipe0 []struct {
		CraftType struct {
			ID string `json:"ID"`
		}
		ItemResult struct {
			ID int `json:"ID"`
		}
	} `json:"ItemIngredientRecipe0"`
	ItemIngredientRecipe1 []struct {
		CraftType struct {
			ID string `json:"ID"`
		}
		ItemResult struct {
			ID int `json:"ID"`
		}
	} `json:"ItemIngredientRecipe1"`
	ItemIngredientRecipe2 []struct {
		CraftType struct {
			ID string `json:"ID"`
		}
		ItemResult struct {
			ID int `json:"ID"`
		}
	} `json:"ItemIngredientRecipe2"`
	ItemIngredientRecipe3 []struct {
		CraftType struct {
			ID string `json:"ID"`
		}
		ItemResult struct {
			ID int `json:"ID"`
		}
	} `json:"ItemIngredientRecipe3"`
	ItemIngredientRecipe4 []struct {
		CraftType struct {
			ID string `json:"ID"`
		}
		ItemResult struct {
			ID int `json:"ID"`
		}
	} `json:"ItemIngredientRecipe4"`
	ItemIngredientRecipe5 []struct {
		CraftType struct {
			ID string `json:"ID"`
		}
		ItemResult struct {
			ID int `json:"ID"`
		}
	} `json:"ItemIngredientRecipe5"`
	ItemIngredientRecipe6 []struct {
		CraftType struct {
			ID string `json:"ID"`
		}
		ItemResult struct {
			ID int `json:"ID"`
		}
	} `json:"ItemIngredientRecipe6"`
	ItemIngredientRecipe7 []struct {
		CraftType struct {
			ID string `json:"ID"`
		}
		ItemResult struct {
			ID int `json:"ID"`
		}
	} `json:"ItemIngredientRecipe7"`
	ItemIngredientRecipe8 []struct {
		CraftType struct {
			ID string `json:"ID"`
		}
		ItemResult struct {
			ID int `json:"ID"`
		}
	} `json:"ItemIngredientRecipe8"`
	ItemIngredientRecipe9 []struct {
		CraftType struct {
			ID string `json:"ID"`
		}
		ItemResult struct {
			ID int `json:"ID"`
		}
	} `json:"ItemIngredientRecipe9"`
}

//Pass a struct item to ItemRecipe.
func GetRecipe(itemweb string) {
	// TODO: We can split the URL using categories, to get smaller payloads of JSON.
	// ABOUT TODO: You want to find an optimal amount of splitting, or just having one big payload (or one reduced payload would be ideal).
	//What this does, is open the file, and read it
	jsonFile, err := http.Get(itemweb)
	if err != nil {
		log.Fatalln(err)
	}
	// Takes the jsonFile.Body, and put it into memory as byteValue array.
	byteValue, err := ioutil.ReadAll(jsonFile.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Body.Close()

	// Recipe Struct information
	var item Recipe
	json.Unmarshal(byteValue, &item)
	//Print out our data to check.
	fmt.Println("Item:" + item.Name)
	fmt.Println("Recipe ID: " + strconv.Itoa(item.ID))
	fmt.Println("URL: " + item.Url)

	// Amount of Ingredients Information
	var amount AmountIngredient
	json.Unmarshal(byteValue, &amount)
	// We need to change the ugly AmountIngredient struct into an array.
	r_amount := reflect.ValueOf(amount)
	n_amount := r_amount.NumField()
	AmountIngredients := make([]string, n_amount)
	for i := 0; i < n_amount; i++ {
		AmountIngredients[i] = fmt.Sprintf(`%v`, r_amount.Field(i))
	}
	fmt.Println(AmountIngredients)

	//Ingredient Recipe Information
	// We need to convert into an array. stuff[i], with [j=2], where it's ctID,iID
	var ingredients IngredientRecipe
	json.Unmarshal(byteValue, &ingredients)
	r_ingred := reflect.ValueOf(ingredients)
	n_ingred := r_ingred.NumField()
	// TODO: Fix the i,j'th elements to be able to put both, CraftTypeID, and Item Result ID.
	CraftingIngredients := make([][]string, n_ingred)
	for i := 0; i < n_ingred; i++ {
		for j := 0; j <= 2; j++ {
			CraftingIngredients[i][j] = fmt.Sprintf(`%v`, r_ingred.Field(i))
		}
	}
	fmt.Println(CraftingIngredients)

	//TODO: Put this information into the database!
}
