// Appends the strings to the xivapi.com/ <- Append things here

package xivapi

import (
	"strings"

	getKeys "./keys"
)

// TODO: ?columns=Attributes,Object.Attribute will significantly lower payload

//Once you have an item ID, it plugs it into the websiteurl.
func UrlRecipe(userfield string, useruniqueID string) string {
	//Example: https://xivapi.com/Recipe/33180?key=
	basewebsite := []byte("https://xivapi.com/")
	field := []byte(userfield)
	uniqueID := []byte(useruniqueID)
	authkey := []byte(getKeys.XivAuthKey)

	//The field isn't complete, we need to append the forward slash at the end.
	completefield := append(field[:], '/')

	//We need to combine the user input stuff
	userinputurl := append(append(basewebsite[:], completefield[:]...), uniqueID[:]...)

	//Now we need the complete URL
	websiteurl := append(append(userinputurl[:], '?'), authkey[:]...)

	s := string(websiteurl)
	return s
}

//UserInputs some item to search. This appends it to the websiteurl.
func UrlSearch(usersearch string) string {
	//Example: https://xivapi.com/search?string=High+Mythrite+Ingot&key=
	var replacer = strings.NewReplacer(" ", "+")
	fixedusersearch := replacer.Replace(usersearch)

	basewebsite := []byte("https://xivapi.com/search?string=")
	authkey := []byte(getKeys.XivAuthKey)
	searchfield := []byte(fixedusersearch)
	userinputurl := append(append(basewebsite[:], searchfield[:]...), '&')
	websiteurl := append(userinputurl[:], authkey[:]...)

	s := string(websiteurl)
	return s
}
