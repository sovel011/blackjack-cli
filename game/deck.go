package game

import (
	"math/rand"
	"strconv"
	"time"
)

const numberOfCards = 52

var suits = [4]string{
	"Diamonds",
	"Hearts",
	"Spades",
	"Clubs",
}

var ranks = [13]string{
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"10",
	"Jack",
	"Queen",
	"King",
	"Ace",
}

// Deck object containing Cards and an offset representing the index of the next to be dealt
type Deck struct {
	cards  [numberOfCards]Card
	offset int
}

// InitDeck initializes a new shuffled deck
func InitDeck() Deck {
	rand.Seed(time.Now().UTC().UnixNano())
	deck := Deck{offset: 0}
	i := 0
	for _, suit := range suits {
		for _, rank := range ranks {
			values := computeCardValues(rank)
			card := Card{rank, suit, values}
			deck.cards[i] = card
			i++
		}
	}
	deck.shuffle()
	return deck
}

// DealCard returns the next card in the deck
func (deck *Deck) DealCard() Card {
	if deck.offset > numberOfCards-1 {
		deck.shuffle()
	}
	card := deck.cards[deck.offset]
	deck.offset = deck.offset + 1
	return card
}

func (deck *Deck) shuffle() {
	for i := range deck.cards {
		j := rand.Intn(numberOfCards)
		card1 := deck.cards[i]
		card2 := deck.cards[j]
		deck.cards[i] = card2
		deck.cards[j] = card1
	}
	deck.offset = 0
}

func computeCardValues(rank string) []int {
	var result []int
	switch rank {
	case "Jack", "Queen", "King":
		result = []int{10}
	case "Ace":
		result = []int{1, 11}
	default:
		intVal, _ := strconv.Atoi(rank)
		result = []int{intVal}
	}
	return result
}
