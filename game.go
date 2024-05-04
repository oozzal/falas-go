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
	for range g.noOfPlayers {
		for j := range 3 {
			g.hands[j].cards = append(g.hands[j].cards, deck.cards[0])
			deck.cards = append(deck.cards[:0], deck.cards[1:]...)
		}
	}
	for i := range g.hands {
		g.hands[i].Sort()
		g.hands[i].Identify()
	}
}

func (g Game) Show(player int) Hand {
	winner := g.hands[player]
	for _, challenger := range g.hands {
		if !winner.IsGreaterThan(challenger) {
			winner = challenger
		}
	}
	return winner
}
