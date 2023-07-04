package main

import (
	"fmt"
	"os"
	"strings"
)

var inputMappings = map[string]string{
	"h": "Hearts",
	"d": "Diamonds",
	"s": "Spades",
	"c": "Clubs",
}

func validateGame(input string) string {
	if input == "A" {
		return "Guess Next Card!"
	}
	fmt.Println("Unknown option", input)
	os.Exit(1)
	panic("Bye!")
}

func validateAnswer(q int, input string, d deck) bool {
	if q == 0 {
		//Pass one: Get card[0] and see if the suit matches what was guessed
		var suit string
		card := d[0]
		if card.Suit == "Diamonds" || card.Suit == "Hearts" {
			suit = "R"
		} else {
			suit = "B"
		}
		return suit == input

	}
	if q == 1 {
		//Pass two: Get card[0] and card[1]
		//See if card 1 is higher or lower based on what was guessed
		previous := d[0]
		current := d[1]
		pVal := previous.Value
		cVal := current.Value

		if input == "H" {
			return cVal > pVal
		} else {
			return cVal < pVal
		}

	}
	if q == 2 {
		//Pass three: Get card[0] and card[1] and see if card[3] is outside
		//or inbetween

		//Step 1: find the higher and lower constraints
		start := d[0]
		previous := d[1]

		sVal := start.Value
		pVal := previous.Value

		var minCard Card
		var maxCard Card
		if sVal < pVal {
			minCard = start
			maxCard = previous
		} else {
			minCard = previous
			maxCard = start
		}

		current := d[2]

		if input == "I" {
			return minCard.Value <= current.Value && current.Value <= maxCard.Value
		} else {
			return current.Value < minCard.Value || current.Value > maxCard.Value
		}

	}
	current := d[3]
	//TODO: Make a map of input to value to reuse
	val := inputMappings[strings.ToLower(input)]

	return val == current.Suit
}

func main() {
	//Let's play a game!
	var game string
	//Print out selections & get user input
	fmt.Println("Select game you wish to play!")
	fmt.Println("[A] Guess Next Card!")
	fmt.Scanln(&game)

	//Validate that the selection is valid
	validateGame(game)
	fmt.Println("Game is: ", game)

	//Create a new deck
	deck := newDeck()
	//Shuffle deck
	deck.shuffleDeck()

	//We play for 5 rounds!
	guessedCorrectly := 0
	questions := []string{
		"Red (R) or Black (B)?",
		"Higher (H) or Lower (L)",
		"Inside (I) or Outside(O)",
		"Suit? Diamonds (D), Hearts(H), Spade (S),Clubs (C) "}

	for i := 0; i < 1; i++ {
		//Pick up a 4 card & remove from deck
		current := emptyDeck()
		current, deck = deal(deck, 4)

		for j, value := range questions {
			fmt.Println(value)
			var answer string
			fmt.Scanln(&answer)

			result := validateAnswer(j, answer, current)
			fmt.Println("Card value : ", current[j])

			if result {
				fmt.Println("CORRECT! 10 POINTS TO SLYTHERIN")
				guessedCorrectly += 1

			} else {
				fmt.Println("You have chosen..Poorly.")
			}

		}
	}
	fmt.Println("End of Game! You answered", guessedCorrectly, "correctly. ")

}
