package main

type Game struct {
	noOfPlayers int
	hands       []Hand
}

func NewGame(noOfPlayers int) Game {
	return Game{noOfPlayers, make([]Hand, noOfPlayers)}
}

func (g *Game) deal() {
	deck := NewDeck()
	if g.noOfPlayers > 17 {
		panic("only 17 players allowed")
	}
	for i := range g.noOfPlayers {
		for j := range 3 {
			g.hands[j].cards[i] = deck.cards[0]
			deck.cards = append(deck.cards[:0], deck.cards[1:]...)
		}
	}
}
