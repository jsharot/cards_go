package main

import (
	"fmt"
)

func main() {
	//Let's play a game!
	var input string
	//Print out selections & get user input
	fmt.Println("Select game you wish to play!")
	fmt.Println("[A] Guess the Card!")
	fmt.Scanln(&input)

	//Validate that the selection is valid
	game := validateGame(input)
	fmt.Println("Game is: ", game.name())

	//Let's get a deck!
	var deck deck
	//If there is a previous deck on hand for that game load it
	//File on local will follow convention GAMENAME_deck.json
	fmt.Println("Checking if there's a previous deck..")
	filename := game.name() + "_deck.json"

	deck = newDeckFromFile(filename)
	if deck == nil || len(deck) < game.minCards() {
		fmt.Println("The previous deck was either not found or doesn't have enough cards to play. Let's get a new deck!")
		//New Deck
		deck = newDeck()
		deck.shuffleDeck()
	}

	//Play game
	deck = playGame(game, deck)
	//Save to file for next time
	deck.toFile(filename)
	fmt.Println(len(deck))

}
