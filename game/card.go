package game

import "fmt"

// Card object containing card rank, suit, and point value(s)
type Card struct {
	rank   string
	suit   string
	values []int // Two values for Aces [1,11]
}

func (card Card) String() string {
	return fmt.Sprintf("%s of %s", card.rank, card.suit)
}
