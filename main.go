package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sovel011/blackjack-cli/game"
)

const dealerHitLimit = 17

func main() {
	play()
}

func play() {
	deck := game.InitDeck()

	// Initialize Dealer hand
	dealer := game.NewHand("Dealer")
	dealer.Push(deck.DealCard())
	fmt.Print(dealer.String())
	dealer.Push(deck.DealCard())

	// Initialize Player hand
	player := game.NewHand("You")
	player.Push(deck.DealCard())
	player.Push(deck.DealCard())
	fmt.Print(player.String())

	blackjack := checkBlackjack(player)
	if !blackjack {
		executePlayerTurn(&player, &deck)
		executeDealerTurn(&dealer, player.CalculateTotal(), &deck)
	}
	printWinner(dealer, player)
}

func checkBlackjack(hand game.Hand) bool {
	result := false
	if hand.CalculateTotal() == game.BlackjackGoal {
		fmt.Println("BLACKJACK!")
		result = true
	}
	return result
}

func executePlayerTurn(hand *game.Hand, deck *game.Deck) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Hit? (y/n)")
		response, err := reader.ReadString('\n')
		if isNegativeResponse(response) || err != nil {
			break
		}

		hand.Push(deck.DealCard())
		fmt.Print(hand.String())
		if hand.IsBust() {
			break
		}
	}
}

func isNegativeResponse(response string) bool {
	cleansedResponse := strings.ReplaceAll(response, "\n", "")
	return cleansedResponse == "n" || cleansedResponse == "N"
}

func executeDealerTurn(dealer *game.Hand, playerTotal int, deck *game.Deck) {
	fmt.Print(dealer.String())

	if dealer.CalculateTotal() < playerTotal {
		executeDealerLoop(dealer, deck)
	}
}

func executeDealerLoop(dealer *game.Hand, deck *game.Deck) {
	for {
		if dealer.IsBust() || isHandExceedingHitLimit(*dealer) {
			break
		}

		dealer.Push(deck.DealCard())
		fmt.Print(dealer.String())
	}
}

func isHandExceedingHitLimit(hand game.Hand) bool {
	return hand.CalculateTotal() >= dealerHitLimit
}

func printWinner(dealer, player game.Hand) {
	if dealer.CalculateTotal() >= player.CalculateTotal() {
		fmt.Print("Dealer wins\n\n")
	} else {
		fmt.Print("You won!\n\n")
	}
}
