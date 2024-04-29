package main

import (
	"math/rand"
)

type Deck struct {
	cards []Card
}

func NewDeck() Deck {
	deck := Deck{}
	deck.init()
	deck.shuffle()
	return deck
}

func (d *Deck) init() {
	for _, face := range faces {
		for _, value := range values {
			d.cards = append(d.cards, Card{face, value})
		}
	}
}

func (d *Deck) shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}
