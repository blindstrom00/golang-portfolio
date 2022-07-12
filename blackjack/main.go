package main

import (
	"fmt"

	"github.com/blindstrom00/deck"
)

func main() {
	cards := deck.New(deck.Deck(4), deck.Shuffle)
	var card deck.Card
	for i := 0; i < 10; i++ {
		card, cards = cards[0], cards[1:]
		fmt.Println(card)
	}
}
