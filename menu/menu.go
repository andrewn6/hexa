package menu

import (
	"fmt"
	"hexa/game"
  "hexa/tutorial"
	"os"
	"github.com/i582/cfmt/cmd/cfmt"
)

func Menu() {
	cfmt.Print(`
          ----
         /    \
    ____/      \____
   /    \      /    \
  /      \____/      \
  \      /    \      /
   \____/      \____/
   /    \      /    \
  /      \____/      \
  \      /    \      /
   \____/      \____/
        \      /
         \____/ 
`)
	cfmt.Println("\n\n Hexa, {{the adventure game}}::cyan")
	fmt.Println("────────────────────────────────────")
	cfmt.Println("{{Select an option to get started}}::cyan")
	cfmt.Println("1) {{Start}}::red")
	cfmt.Println("2) {{How to}}::red")
	cfmt.Println("3) {{Quit}}::red")

	var input string

	for input != "1" && input != "2" && input != "3" {
		fmt.Print("→")
		fmt.Scan(&input)

		switch input {
		case "1":
			game.Game()
		case "2":
      tutorial.Tutorial()
		case "3":
			cfmt.Println("{{Quiting...}}::red")
			os.Exit(0)
		}
	}

}
