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

func TestSingleDeck(t *testing.T) {
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

func TestJokers(t *testing.T) {
	cards := New(Jokers(4))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 4 {
		t.Error("Expected 4 jokers, received: ", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Seven || card.Rank == Nine
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Seven || c.Rank == Nine {
			t.Error("Expected all sevens and nines to be removed from the deck.")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(4))
	if len(cards) != 52*4 {
		t.Errorf("Expected to get %d cards, received %d cards.", 52*4, len(cards))
	}
}
