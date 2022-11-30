package game

// Imports
import (
	"fmt"
  "os"
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
  6: {Item_Name: "Diamonds"},
}

// Main game logic
func Game() {
	var Health int = 100
	var Coins int = 200
	var input string
	var Wood int = 0
	var Leafs int = 0

	cfmt.Println("{{You woke up in a old abandoned house, after a huge combat fight with a huge monster. Do you stay to explore? or do you exit the house}}::cyan.\n")
	cfmt.Println("{{If you choose to stay, type 1, if you choose to leave type 2}}::green.\n")

	for input != "1" && input != "2" {
		fmt.Print("→")
		fmt.Scan(&input)

		switch input {
		case "1":
			hp := Health - 50
			cfmt.Println("{{Uh oh.. looks like a booby trap got set off. Your health is now at}}::red: ", hp, "%\n")
		case "2":
			cfmt.Println("{{You left the house and went to the a foreign village in the disance where a friendly local offers to feed and give you a place to rest.}}::cyan\n")
		}
	}

	cfmt.Println("{{After running away from the house, you end up in a village. You can go to the shop to buy items, or you can go to the forest to gather resources.}}::cyan.\n")
	cfmt.Println("{{If you choose to go to the shop type 1, and if you choose to go to the forest to gather resources type 2.}}::green\n")

	var input_2 string
	for input_2 != "1" && input_2 != "2" {
		fmt.Print("→")
		fmt.Scan(&input_2)

		switch input_2 {
		case "1":
			coins := Coins - 100
			cfmt.Println("{{You go to the shop and buy a great steel rustic sword for}}::cyan: ", coins, " coins\n")
      cfmt.Println("{{Your balance is now at}}::green ", coins, " coins\n")
		case "2":
			coins := Coins + 20
			wood := Wood + 2
			leafs := Leafs + 5
			cfmt.Println("{{You enter the forest and gather some wood, leafs.}}::green")
			cfmt.Println("{{You now have}}::green: ", coins, " coins, ", wood, " wood, ", leafs, " leafs\n ")
		}

	}

	cfmt.Println("{{You make your way back to the house of the local who offered to give you a bed for the night}}::cyan\n")
  cfmt.Println("{{The local guides you to your room for the night, you settle in and decide to rest}}::cyan\n")
  cfmt.Println("{{After falling asleep, a few hours later you hear screaming and a ton of chaos. You get out of your room and go downstairs and look out the window. You see the dragon you fought back in that abandoned mansion destroying the village}}::red\n")
  cfmt.Println("{{If you choose to attack the dragon with your sword, choose 1. If you chose to run away and follow the rest of the villagers choose 2. Lastly, if you choose to stay in the house choose 3.}}::green\n") 

  var input_3 string
  for input_3 != "1" && input_3 != "2" && input_3 != "3" {
    fmt.Print("→")
    fmt.Scan(&input_3)

    switch input_3 {
    case "1":
      coins := Coins + 100
      cfmt.Println("{{You chose to attack the dragon, after a tough fight you WIN! You collect}}::green ", coins, " coins\n")
      cfmt.Println("{{Congrats! You won hexa, the adventure game. Thanks for playing!}}::yellow|underline\n")
      os.Exit(0)
      
    case "2":
      hp := Health - 100
      cfmt.Println("{{Uh oh.. The dragon notices you and hunts you down and you lose! Health}}::red: ", hp, "%\n")
      os.Exit(0)

    case "3": 
      hp := Health - 100
      cfmt.Println("{{You decide to stay in the house but the dragon completely obilterates it and you lose! Health}}::red: ", hp, "%\n")
      os.Exit(0)
    }
  }
	if Health <= 0 {
		cfmt.Println("{{You did not pass, game over :(.\n}}::red")
    os.Exit(0)
	}
}
