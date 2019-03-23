package items

import (
	"fmt"

	getKeys "./keys"
)

// TODO: ?columns=Attributes,Object.Attribute will significantly lower payload.

func UrlAPI(userfield string, useruniqueID string) string {
	basewebsite := []byte("https://xivapi.com/")
	field := []byte(userfield)
	uniqueID := []byte(useruniqueID)
	authkey := []byte(getKeys.XivAuthKey)

	//The field isn't complete, we need to append the forward slash at the end.
	completefield := append(field[:], '/')

	//We need to combine the user input stuff
	userinputurl := append(append(basewebsite[:], completefield[:]...), uniqueID[:]...)

	//Now we need the complete URL
	websiteurl := append(userinputurl[:], authkey[:]...)

	s := string(websiteurl)
	fmt.Println(s)

	return s
}

// Now that we can append, we need to choose options.
// We need the user to be able to search for an item.
//It's tedious to write
// AmountIngredient0-9,
// Name,ItemResultTargetID,ID,Url,
// and this 9 times -> ItemIngredientRecipe0.*.CraftType,ItemIngredientRecipe0.*.ItemResult.ID,ItemIngredientRecipe0.*.ItemResult.Name
