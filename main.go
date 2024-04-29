package main

import "fmt"

func main() {
	game := NewGame(3)
	game.deal()
	fmt.Println(game.hands)
}
