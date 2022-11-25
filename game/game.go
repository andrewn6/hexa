package game

// Imports
import (
	"fmt"

	"github.com/i582/cfmt/cmd/cfmt"
)

// Types, Player & Items.
type Player struct {
	Health       int
	Username     string
	CoinGathered int
	Score        int
	Coins        int
	Item_Name    string
	Action       string
}

var Items = map[int]*Player{
	1: {Item_Name: "Sword"},
	2: {Item_Name: "Health Drink"},
	3: {Item_Name: "Coin"},
	4: {Item_Name: "Wood"},
	5: {Item_Name: "Leafs"},
}

// Main game logic
func Game() {
	var Health int = 100
	var Coins int = 200
	var input string
	var Wood int = 0
	var Leafs int = 0

	cfmt.Println("{{You woke up in a old abandoned house, after a huge combat fight with the biggest dragon in the village. Do you stay to explore? or do you exit the house}}::cyan.\n")
	cfmt.Println("If you choose to stay, type 1, if you choose to leave type 2.\n")

	for input != "1" && input != "2" {
		fmt.Print("→")
		fmt.Scan(&input)

		switch input {
		case "1":
			hp := Health - 50
			cfmt.Println("{{Uh oh.. looks like a booby trap got set off. Your health is now at}}::red: ", hp, "%\n")
		case "2":
			cfmt.Println("{{You left the house and went to the a foreign village in the disance where a friendly local offers to feed and give you a place to rest.}}::green\n")
		}
	}

	cfmt.Println("{{You are now in a village, you can go to the shop to buy items, or you can go to the forest to gather resources.}}::cyan.\n")
	cfmt.Println("If you choose to go to the shop type 1, and if you choose to go to the forest to gather resources type 2.\n")

	var input_2 string
	for input_2 != "1" && input_2 != "2" {
		fmt.Print("→")
		fmt.Scan(&input_2)

		switch input_2 {
		case "1":
			coins := Coins - 100
			cfmt.Println("{{You go to the shop and buy a great steel rustic sword for}}::cyan: ", coins)
		case "2":
			coins := Coins + 20
			wood := Wood + 2
			leafs := Leafs + 5
			cfmt.Println("{{You enter the forest and gather some wood, leafs. You even found}}::green.\n", coins, wood, leafs)
		}

	}
	if Health <= 0 {
		cfmt.Println("{{You did not pass, game over :(.\n}}::red")
	}
}
