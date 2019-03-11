package items

import (
	"encoding/json" // Passes the byteValue to our struct.
	"fmt"           // Println etc.
	"io/ioutil"     // Converts jsonFile into a byteValue, which is our byte array.

	// Opens files and store it into jsonFile, in our memory
	"log"
	"net/http"
	"strconv" // Converts ints to strings etc.
)

// An ItemRecipe will look like this
// A container, which we can call the ItemRecipe Structure
type ItemRecipe struct {
	//The outer values
	AmountIngredient0  int    `json:"AmountIngredient0"`
	AmountIngredient1  int    `json:"AmountIngredient1"`
	AmountIngredient2  int    `json:"AmountIngredient2"`
	AmountIngredient3  int    `json:"AmountIngredient3"`
	AmountIngredient4  int    `json:"AmountIngredient4"`
	AmountIngredient5  int    `json:"AmountIngredient5"`
	AmountIngredient6  int    `json:"AmountIngredient6"`
	AmountIngredient7  int    `json:"AmountIngredient7"`
	AmountIngredient8  int    `json:"AmountIngredient8"`
	AmountIngredient9  int    `json:"AmountIngredient9"`
	Name               string `json:"Name"`
	ItemResultTargetID int    `json:"ItemResultTargetID"`
	ID                 int    `json:"ID"`
	Url                string `json:"Url"`
	//The outer objects
	ItemIngredient0 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient0"`
	ItemIngredient1 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient1"`
	ItemIngredient2 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient2"`
	ItemIngredient3 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient3"`
	ItemIngredient4 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient4"`
	ItemIngredient5 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient5"`
	ItemIngredient6 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient6"`
	ItemIngredient7 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient7"`
	ItemIngredient8 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient8"`
	ItemIngredient9 struct {
		ID   int    `json:"ID"`
		Name string `json:"Name"`
	} `json:"ItemIngredient9"`

	//The outer arrays of objects
	ItemIngredientRecipe0 []struct {
		//The array elements are objects, containing objects
		CraftType struct {
			Name string `json:"Name"`
		}
		ItemResult struct {
			Name string `json:"Name"`
		}
	} `json:"ItemIngredientRecipe0"`
	ItemIngredientRecipe1 []struct {
		CraftType struct {
			Name string `json:"Name"`
		}
		ItemResult struct {
			Name string `json:"Name"`
		}
	} `json:"ItemIngredientRecipe1"`
	ItemIngredientRecipe2 []struct {
		CraftType struct {
			Name string `json:"Name"`
		}
		ItemResult struct {
			Name string `json:"Name"`
		}
	} `json:"ItemIngredientRecipe2"`
	ItemIngredientRecipe3 []struct {
		CraftType struct {
			Name string `json:"Name"`
		}
		ItemResult struct {
			Name string `json:"Name"`
		}
	} `json:"ItemIngredientRecipe3"`
	ItemIngredientRecipe4 []struct {
		CraftType struct {
			Name string `json:"Name"`
		}
		ItemResult struct {
			Name string `json:"Name"`
		}
	} `json:"ItemIngredientRecipe4"`
	ItemIngredientRecipe5 []struct {
		CraftType struct {
			Name string `json:"Name"`
		}
		ItemResult struct {
			Name string `json:"Name"`
		}
	} `json:"ItemIngredientRecipe5"`
	ItemIngredientRecipe6 []struct {
		CraftType struct {
			Name string `json:"Name"`
		}
		ItemResult struct {
			Name string `json:"Name"`
		}
	} `json:"ItemIngredientRecipe6"`
	ItemIngredientRecipe7 []struct {
		CraftType struct {
			Name string `json:"Name"`
		}
		ItemResult struct {
			Name string `json:"Name"`
		}
	} `json:"ItemIngredientRecipe7"`
	ItemIngredientRecipe8 []struct {
		CraftType struct {
			Name string `json:"Name"`
		}
		ItemResult struct {
			Name string `json:"Name"`
		}
	} `json:"ItemIngredientRecipe8"`
	ItemIngredientRecipe9 []struct {
		CraftType struct {
			Name string `json:"Name"`
		}
		ItemResult struct {
			Name string `json:"Name"`
		}
	} `json:"ItemIngredientRecipe9"`
}

//Pass a struct item to ItemRecipe.
func GetRecipe(itemweb string) {

	// TODO: ?columns=Attributes,Object.Attribute will significantly lower load.

	//What this does, is open the file, and store it as memory jsonFile, and we need to read the body component as a byte slice.
	jsonFile, err := http.Get(itemweb)
	if err != nil {
		log.Fatalln(err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile.Body)
	if err != nil {
		log.Fatalln(err)
	}

	/* This is the regular file equivalent for Debug.
	jsonFile, err := os.Open("SeeingHordeAxe.json")
		if err != nil {
			fmt.Println(err)
			os.Exit(1001) // Error 1001 - No Json file found.
		}


		fmt.Println("Success.")

		// We need to keep the memory alive.
		defer jsonFile.Close()


	// We need to read the memory body as a byte slice.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	*/

	var item ItemRecipe

	//We now just unmarshal the byteArray into the struct
	json.Unmarshal(byteValue, &item)

	//Print out our data to check.
	fmt.Println("Item:" + item.Name)
	fmt.Println("Recipe ID: " + strconv.Itoa(item.ID))
	fmt.Println("URL: " + item.Url)
	fmt.Println("First Ingredient: " + item.ItemIngredient0.Name)
	fmt.Println("First Ingredient Amount: " + strconv.Itoa(item.AmountIngredient0))
	for i := 0; i < len(item.ItemIngredientRecipe0); i++ {
		fmt.Println("Ingredient0: " + item.ItemIngredientRecipe0[i].ItemResult.Name)
		fmt.Println("Ingredient0: " + item.ItemIngredientRecipe0[i].CraftType.Name)
	}

	/*
		// We can equate these string values, and check if it's right. Which it will be.
		if item.ItemIngredientRecipe0[0].ItemResult.Name == item.ItemIngredient0.Name {
			fmt.Println("Success")
		} else {
			fmt.Println("Failure")
		}
	*/

	//TODO: Put this information into the database!
}
