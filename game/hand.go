package game

import "fmt"

// BlackjackGoal - the goal number of points to have in a given hand
const BlackjackGoal = 21
const bustValue = -1

// Hand object containing list of cards and name of the player
type Hand struct {
	name  string
	cards []Card
}

// NewHand initializes a new hand struct with a given name
func NewHand(name string) Hand {
	return Hand{name: name}
}

// Push adds a new card to the hand
func (hand *Hand) Push(card Card) {
	hand.cards = append(hand.cards, card)
}

// CalculateTotal returns the number of points in the hand. If the hand is a bust -1 is returned
func (hand Hand) CalculateTotal() int {
	return calTotal(0, hand.cards)
}

// IsBust returns true if the cards in the given hand exceed 21
func (hand Hand) IsBust() bool {
	return hand.CalculateTotal() < 0
}

func calTotal(sum int, cards []Card) int {
	var total int
	if sum > BlackjackGoal {
		return bustValue
	} else if len(cards) == 0 {
		return sum
	}

	card := cards[0]
	if len(card.values) == 1 {
		newSum := sum + card.values[0]
		total = calTotal(newSum, cards[1:])
	} else {
		sumA := sum + card.values[0]
		sumB := sum + card.values[1]
		total = max(calTotal(sumA, cards[1:]), calTotal(sumB, cards[1:]))
	}
	return total
}

func max(a, b int) int {
	var result int
	if a > b {
		result = a
	} else {
		result = b
	}
	return result
}

func (hand Hand) String() string {
	result := fmt.Sprintf("\n%s:\n", hand.name)
	for _, card := range hand.cards {
		result += fmt.Sprintf("%v\t", card.String())
	}

	total := hand.CalculateTotal()
	if total > 0 {
		result += fmt.Sprintf("\nShowing: %v \n\n", total)
	} else {
		result += "\nBUST\n\n"
	}
	return result
}
