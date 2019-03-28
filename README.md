## Marketboard Project
Designed to calculate the most profitable crafted item from a popular MMORPG, Final Fantasy XIV.

## Motivation
This is similar to ffxivmb.com, where they grab user obtained data, and find organize the data to show items that have the greatest profit.
However, it's difficult to navigate through ffxivmb, and you have to individually go through the items to show how much profit it gives when creating the item.
This project is meant to allow you to easily see whether or not an item will make you a profit by crafting it, rather than a deficit if you could just sell the mats for more gil.
This project can add more additional features such as easier crafting and gathering, but this is way in the future.


## Tech/framework used
<b>Built with</b>
- [Golang](https://golang.org/)
- [MySQL](https://www.mysql.com/)

## Features
Profits / Costs of Items you want to craft.
Item Recipes shown along with the cost of those materials.

## Code Example
Converts JSON Payloads from GET requests on the xivapi, and convert it into a caching layer to allow easier access to identify items and materials with the costs that are implied in the database.

## Installation
For current build, create a apikeys.go file with your api key string key=######.
Then compile and use through terminal.

## API Reference
- [XIVAPI] (https://xivapi.com/)

## How to use?
Currently not fully designed for user applications.

## License
MIT © [2019] (Jacob Nguyen)