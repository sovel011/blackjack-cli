package game

import (
	"reflect"
	"testing"
)

func TestInitDeck(t *testing.T) {
	deck := InitDeck()

	numCards := len(deck.cards)
	if numCards != numberOfCards {
		t.Error("Unexpected number of cards in new deck, got ", numCards)
	}

	numAces := 0
	for _, card := range deck.cards {
		if card.rank == "Ace" {
			numAces++
		}
	}
	if numAces != 4 {
		t.Error("Expected 4 aces, got ", numAces)
	}

	numClubs := 0
	for _, card := range deck.cards {
		if card.suit == "Clubs" {
			numClubs++
		}
	}
	if numClubs != 13 {
		t.Error("Expected 13 clubs, got ", numClubs)
	}
}

func TestDealCard(t *testing.T) {
	deck := InitDeck()
	card := deck.DealCard()

	if card.suit == "" {
		t.Error("Card suit cannot be empty")
	}

	if card.rank == "" {
		t.Error("Card rank cannot be empty")
	}

	if deck.offset != 1 {
		t.Error("Exepected deck offset to be 1, got ", deck.offset)
	}
}

func TestDeal200Cards(t *testing.T) {
	deck := InitDeck()
	for i := 0; i < 200; i++ {
		deck.DealCard()
	}

	if deck.offset != 44 {
		t.Error("Expected deck offset of 44, got ", deck.offset)
	}
}

func TestShuffle(t *testing.T) {
	deck1 := InitDeck()
	deck2 := deck1

	deck1.shuffle()
	if reflect.DeepEqual(deck1.cards, deck2.cards) {
		t.Error("No change in shuffled deck")
	}

	deck1.DealCard()
	deck1.shuffle()
	if deck1.offset != 0 {
		t.Error("Expected deck offset to be reset, got ", deck1.offset)
	}
}

func TestComputeCardValues(t *testing.T) {
	jackValue := computeCardValues("Jack")
	if jackValue[0] != 10 {
		t.Error("Expected jack point value of 10, got ", jackValue[0])
	}

	aceValue := computeCardValues("Ace")
	if aceValue[0] != 1 && aceValue[1] != 11 {
		t.Error("Expected ace point value of [1,11], got ", aceValue)
	}

	twoValue := computeCardValues("2")
	if twoValue[0] != 2 {
		t.Error("Expected 2 point value of 2, got ", twoValue[0])
	}
}
