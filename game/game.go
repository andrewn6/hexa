package game

import (
  "os"
  "bufio"
  "fmt"
)

type Player struct {
  Health int
  Username string
  CoinGathered int
  Score int
}

type Item struct {
  Name string
  Action string
  
}

var Items = map[int]*Item{
  1: {Name:"Sword"},
  2: {Name: "Health Drink"}
  3: {Name: "Coin"}
}

func Game() {
  PlayerInput()
}
