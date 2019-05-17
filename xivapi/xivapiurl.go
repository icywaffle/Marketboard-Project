// Appends the strings to the xivapi.com/ <- Append things here

package xivapi

import (
	"strings"

	getKeys "./keys"
)

// TODO: ?columns=Attributes,Object.Attribute will significantly lower payload

//Creates the URL for recipes and items
func UrlItemRecipe(userID string) string {
	//Example: https://xivapi.com/Recipe/33180
	//Example: https://xivapi.com/Item/24322
	basewebsite := []byte("https://xivapi.com/")
	field := []byte("recipe")
	uniqueID := []byte(userID)
	completefield := append(field[:], '/')
	userinputurl := append(append(basewebsite[:], completefield[:]...), uniqueID[:]...)

	//Finishing the url with the AuthKey
	authkey := []byte(getKeys.XivAuthKey)
	websiteurl := append(append(userinputurl[:], '?'), authkey[:]...)

	s := string(websiteurl)
	return s
}

//UserInputs some item to search. This appends it to the websiteurl.
func UrlSearch(usersearch string) string {
	//Example: https://xivapi.com/search?string
	basewebsite := []byte("https://xivapi.com/search?string=")

	//Example: https://xivapi.com/search?string=High+Mythrite+Ingot
	var replacer = strings.NewReplacer(" ", "+")
	fixedusersearch := replacer.Replace(usersearch)
	searchfield := []byte(fixedusersearch)
	userinputurl := append(append(basewebsite[:], searchfield[:]...), '&')

	authkey := []byte(getKeys.XivAuthKey)
	websiteurl := append(userinputurl[:], authkey[:]...)

	s := string(websiteurl)
	return s
}

func UrlPrices(useritemid string) string {
	//Example: https://xivapi.com/market/item/3?servers=Phoenix,Lich,Moogle

	//Produces : https://xivapi.com/market/item/3
	itemwebsitefield := []byte("https://xivapi.com/market/item/")
	itemid := []byte(useritemid)
	basewebsite := append(itemwebsitefield[:], itemid[:]...)

	//Produces :https://xivapi.com/market/item/3?servers=Phoenix,Lich,Moogle&
	//TODO:Let's just use Sargatanas for now for simple structs, then expand later. ?servers=Adamantoise,Cactuar,Faerie,Gilgamesh,Jenova,Midgardsormr,Sargatanas,Siren
	servers := []byte("?servers=Sargatanas")
	userinputurl := append(append(basewebsite[:], servers[:]...), '&')

	//Attaches key to the end.
	authkey := []byte(getKeys.XivAuthKey)
	websiteurl := append(userinputurl[:], authkey[:]...)

	s := string(websiteurl)
	return s
}
