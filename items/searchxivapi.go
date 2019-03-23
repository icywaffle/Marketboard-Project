package items

import (
	"fmt"

	getKeys "./keys"
)

// Appends the strings to the xivapi.com

// TODO: ?columns=Attributes,Object.Attribute will significantly lower payload

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
