package main

import (
	"fmt"
	"strings"

	"github.com/blindstrom00/deck"
)

func main() {
	cards := deck.New(deck.Deck(4), deck.Shuffle)
	var card deck.Card
	var player, dealer Hand
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}
	var input string
	for input != "s" {
		fmt.Println("Player:", player)
		fmt.Println("Dealer:", dealer.DealerString())
		fmt.Println("What would you like to (h)it or (s)tand?")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, cards = draw(cards)
			player = append(player, card)
		}
	}
	fmt.Println("-- Final Hands --")
	fmt.Println("Player:", player)
	fmt.Println("Dealer:", dealer)
}

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", HIDDEN"
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
