package items

import (
	"fmt"

	getKeys "./keys"
)

func UrlAPI(userfield string, useruniqueID string) string {
	basewebsite := []byte("https://xivapi.com/")
	field := []byte(userfield)
	uniqueID := []byte(useruniqueID)
	pretty := []byte("?pretty=1")
	authkey := []byte(getKeys.XivAuthKey)

	//The field isn't complete, we need to append the forward slash at the end.
	completefield := append(field[:], '/')

	//We need to combine the user input stuff
	userinputurl := append(append(basewebsite[:], completefield[:]...), uniqueID[:]...)

	//Now we need the complete URL
	websiteurl := append(append(userinputurl[:], pretty[:]...), authkey[:]...)

	s := string(websiteurl)
	fmt.Println(s)
	return s
}
