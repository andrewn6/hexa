package menu

import (
	"fmt"
	"os"
  "github.com/i582/cfmt/cmd/cfmt"
)

func main() {
}

func Menu() {
	cfmt.Println("\n\n Hexa, {{the adventure game}}::red")
	fmt.Println("────────────────────────────────────")
	fmt.Println("Select an option to get started")
	fmt.Println("1) Start")
	fmt.Println("2) How to")
	fmt.Println("3) Quit")

	var input string

	for input != "1" && input != "2" && input != "3" {
		fmt.Print("→")
		fmt.Scan(&input)

		switch input {
		case "1":
			fmt.Print("Ok! Let's get started!")
			// todo add / import game func here
		case "2":
			fmt.Print("Now entering the tutorial...")
			// todo add tutorial func
		case "3":
			fmt.Print("Ok.. see you next time!")
			os.Exit(0)
		}
	}

}
