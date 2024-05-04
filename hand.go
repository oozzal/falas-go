package main

import (
	"cmp"
	"slices"
)

type Hand struct {
	cards  []Card
	result Result
}

type Result struct {
	name string
	rank int
}

var (
	Result_Badi      = Result{"Badi", 1}
	Result_Joot      = Result{"Joot", 2}
	Result_Falas     = Result{"Falas", 3}
	Result_Run       = Result{"Run", 4}
	Result_DoubleRun = Result{"DoubleRun", 5}
	Result_Trial     = Result{"Trial", 6}
)

func (h Hand) Rank() int {
	return h.result.rank
}

func (h Hand) Sort() {
	slices.SortFunc(h.cards, func(a, b Card) int {
		return cmp.Compare(a.Rank(), b.Rank())
	})
}

func (h *Hand) Identify() {
	output := Result_Badi

	ranks := []int{h.cards[0].Rank(), h.cards[1].Rank(), h.cards[2].Rank()}
	slices.Sort(ranks)
	ranks = slices.Compact(ranks)

	if len(ranks) == 1 {
		h.result = Result_Trial
		return
	}

	if len(ranks) == 2 {
		h.result = Result_Joot
		return
	}

	faces := []string{h.cards[0].face, h.cards[1].face, h.cards[2].face}
	slices.Sort(faces)
	faces = slices.Compact(faces)

	if len(faces) == 1 {
		output = Result_Falas
	}

	if ranks[2]-ranks[0] == 2 || ranks[2]-ranks[1] == 11 {
		if output == Result_Falas {
			h.result = Result_DoubleRun
			return
		} else {
			h.result = Result_Run
			return
		}
	}

	h.result = output
}

func (h Hand) Display() string {
	output := ""
	for _, card := range h.cards {
		output += card.Display() + " "
	}
	return output + "=> " + h.result.name
}

func (h1 Hand) IsGreaterThan(h2 Hand) bool {
	if h1.Rank() > h2.Rank() {
		return true
	} else if h2.Rank() > h1.Rank() {
		return false
	}

	h1Ranks := []int{h1.cards[0].Rank(), h1.cards[1].Rank(), h1.cards[2].Rank()}
	h2Ranks := []int{h2.cards[0].Rank(), h2.cards[1].Rank(), h2.cards[2].Rank()}

	if h1Ranks[2] > h2Ranks[2] {
		return true
	} else if h1Ranks[2] < h2Ranks[2] {
		return false
	}
	if h1Ranks[1] > h2Ranks[1] {
		return true
	} else if h1Ranks[1] < h2Ranks[1] {
		return false
	}
	if h1Ranks[0] > h2Ranks[0] {
		return true
	} else if h1Ranks[0] < h2Ranks[0] {
		return false
	}
	return false
}
