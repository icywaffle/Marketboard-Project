## Marketboard Project
Designed to calculate whether a crafted item will give you profit.

## Motivation
New Motivation: Since it will be impossible to grab information about every single item in the game without putting a stressed load on any API,
it's better to just ask whether or not an item that you see will net you some profit, and what amount of profit.

Then we can store this information into the database to help see which items will actually net you a better amount of profit.

This database begins to become better over time, the more searches that are used, since they will give the database more information with each request.



## Tech/framework used
<b>Built with</b>
- [Golang](https://golang.org/)
- [MongoDB](https://www.mongodb.com/)
- [MongoDB-Go-Driver](https://github.com/mongodb/mongo-go-driver)

## Features
Profits / Costs of Items you want to craft.
Item Recipes shown along with the cost of those materials.

## Code Example

`weburl := xivapi.UrlItemRecipe(userchoice, strconv.Itoa(userID)`

What this code does is to give you the URL corresponding to your choice of item, or recipe id. This is the xivapi url.

`xivapi.Getitem(weburl), userchoice)`

This is the example of the simplified backend code.
What it does, is take the url from the xivapi, and then parse it into the database.
However, there's no need to parse the information, if we already have it in the database. The backend code will handle this. (Later it will skip the weburl all together.)

For recipe ID, it will give you all the information about what composes an item with that recipe.
For item IDs, eventually it will give you the information about the item itself, and then gives you a bunch of recipeIDs.

Eventually it will also spit out the prices, or manipulate the prices through the backend, and give you a result profit at the end.

## Installation
For current build,

Install MongoDB, and create a server that uses the default port 27017.

Create an XIVAPI account and obtain your own private key.

Create a folder inside xivapi called,

`database`

Then create a go file that contains

`package keys`

`var XivAuthKey string = "private_key=#######"`

Finally, if you have installed Go, in the console, CD into the main folder directory,

`C:\Marketboard-Project> go run main.go`

## API Reference
- [XIVAPI] (https://xivapi.com/)

## How to use?
Development usage only.

## License
MIT © [2019] (Jacob Nguyen)