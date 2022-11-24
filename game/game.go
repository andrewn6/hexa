package game

import (
	"fmt"
	"os"

	"github.com/i582/cfmt/cmd/cfmt"
)

type Player struct {
	Health       int
	Username     string
	CoinGathered int
	Score        int
	Item_Name    string
	Action       string
}

var Items = map[int]*Player{
	1: {Item_Name: "Sword"},
	2: {Item_Name: "Health Drink"},
	3: {Item_Name: "Coin"},
}

func Game() {
	for {
		cfmt.Println("{{You woke up in a old abandoned house, after a huge combat fight with the biggest dragon in the village. Do you stay to explore? or do you exit the house}}::cyan.\n")
		cfmt.Println("If you choose to stay, type 1, if you choose to leave type 2.\n")

		var Health int
		var input string

		for input != "1" && input != "2" {
			fmt.Print("→")
			fmt.Scan(&input)

			switch input {

			case "1":
				cfmt.Println("{{Uh oh.. looks like a booby trap got set off. You lost!}}::red.\n")
				os.Exit(0)
			}

			switch input {
				case "2":
				cfmt.Println("{{Great choice! You escaped the house and you see an unfamiliar town in the distance. You run over to town for help and meet a friendly local who offers to let you sleep in their home}}::green.\n")
			}

		}

		if Health <= 0 {
			cfmt.Println("{{You did not pass, game over :(.}}::red")
		}
	}
}
