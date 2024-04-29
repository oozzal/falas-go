package main

import "fmt"

func main() {
	game := NewGame(3)
	game.deal()
	for _, hand := range game.hands {
		fmt.Println(hand.Display())
	}
	fmt.Println("Winner:")
	fmt.Println(game.Show(0).Display())

	// hand := Hand{
	// 	cards: []Card{
	// 		{face: faces[0], value: Value{"4", 4}},
	// 		{face: faces[2], value: Value{"10", 10}},
	// 		{face: faces[3], value: Value{"J", 11}},
	// 	},
	// }
	// hand.Sort()
	// hand.Identify()
	// other := Hand{
	// 	cards: []Card{
	// 		{face: faces[0], value: Value{"5", 5}},
	// 		{face: faces[1], value: Value{"10", 10}},
	// 		{face: faces[2], value: Value{"10", 10}},
	// 	},
	// }
	// other.Sort()
	// other.Identify()
	// fmt.Println(hand.Display())
	// fmt.Println(other.Display())
	// fmt.Println(hand.IsGreaterThan(other))
}
