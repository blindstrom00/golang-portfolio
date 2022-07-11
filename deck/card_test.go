package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: 4, Suit: Spade})
	fmt.Println(Card{Rank: 9, Suit: Diamond})
	fmt.Println(Card{Rank: 2, Suit: Heart})
	fmt.Println(Card{Rank: 10, Suit: Club})
	fmt.Println(Card{Rank: King, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Heart})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Four of Spades
	// Nine of Diamonds
	// Two of Hearts
	// Ten of Clubs
	// King of Diamonds
	// Jack of Hearts
	// Joker
}

func TestDeck(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Error("Incorrect amount of cards in the deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	huh := Card{Rank: Ace, Suit: Spade}
	if cards[0] != huh {
		t.Error("Expected to get Ace of Spades, received", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	huh := Card{Rank: Ace, Suit: Spade}
	if cards[0] != huh {
		t.Error("Expected to get Ace of Spades, received", cards[0])
	}
}
