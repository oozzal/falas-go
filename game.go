package main

type Game struct {
	hands       []Hand
	noOfPlayers int
}

func NewGame(noOfPlayers int) Game {
	return Game{make([]Hand, noOfPlayers), noOfPlayers}
}

func (g *Game) deal() {
	deck := NewDeck()
	if g.noOfPlayers > 17 {
		panic("only 17 players allowed")
	}
	for i := range g.noOfPlayers {
		g.hands[i].cards = deck.cards[:3]
		deck.cards = deck.cards[3:]
		g.hands[i].Identify()
	}
}

// player: index of the player who wants to show his hand
func (g Game) Show(player int) Hand {
	winner := g.hands[player]
	for _, challenger := range g.hands {
		if challenger.IsGreaterThan(winner) {
			winner = challenger
		}
	}
	return winner
}
