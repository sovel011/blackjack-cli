package game

import "testing"

func TestCardString(t *testing.T) {
	card := Card{rank: "King", suit: "Hearts"}
	str := card.String()
	if str != "King of Hearts" {
		t.Error("Expected 'King of Hearts',got ", str)
	}
}
