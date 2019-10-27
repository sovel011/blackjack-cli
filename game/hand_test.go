package game

import "testing"

func TestNewHand(t *testing.T) {
	name := "Bob"
	hand := NewHand(name)
	if hand.name != name {
		t.Error("Expected name to be Bob, got ", hand.name)
	}
}

func TestPush(t *testing.T) {
	hand := Hand{}
	card := Card{}
	hand.Push(card)
	if len(hand.cards) != 1 {
		t.Error("Expected hand size of 1, got ", len(hand.cards))
	}
}

func TestCalculateTotalSimple(t *testing.T) {
	hand := Hand{}
	hand.Push(Card{rank: "King", values: []int{10}})
	hand.Push(Card{rank: "9", values: []int{9}})
	total := hand.CalculateTotal()

	if total != 19 {
		t.Error("Expected total of 19, got ", total)
	}
}

func TestCalculateTotalComplex(t *testing.T) {
	hand := Hand{}
	hand.Push(Card{rank: "Ace", values: []int{1, 11}})
	hand.Push(Card{rank: "9", values: []int{9}})
	hand.Push(Card{rank: "Ace", values: []int{1, 11}})
	total := hand.CalculateTotal()

	if total != 21 {
		t.Error("Expected total of 21, got ", total)
	}
}

func TestCalculateTotalBust(t *testing.T) {
	hand := Hand{}
	hand.Push(Card{rank: "King", values: []int{10}})
	hand.Push(Card{rank: "King", values: []int{10}})
	hand.Push(Card{rank: "King", values: []int{10}})
	total := hand.CalculateTotal()

	if total != -1 {
		t.Error("Expected total of -1, got ", total)
	}
}

func TestIsBust(t *testing.T) {
	hand := Hand{}
	hand.Push(Card{rank: "King", values: []int{10}})
	hand.Push(Card{rank: "King", values: []int{10}})
	hand.Push(Card{rank: "King", values: []int{10}})

	if !hand.IsBust() {
		t.Error("Hand with 3 kings not considered a bust")
	}
}

func TestHandString(t *testing.T) {
	hand := Hand{}
	if hand.String() == "" {
		t.Error("Hand string value is empty")
	}
}
