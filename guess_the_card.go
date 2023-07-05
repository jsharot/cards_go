package main

import (
	"fmt"
	"strings"
)

type guessTheCard struct{}

func (r guessTheCard) name() string {
	return "guess_the_card"
}

func (r guessTheCard) minCards() int {
	return 4
}

func (r guessTheCard) play(d deck) deck {
	//House rules : 1 round is 4 questions at a time.
	//1: Red or Black
	//2: Higher or lower
	//3: Inside or Outside (Inside contains the card in questions)
	//4: Suit

	questions := []string{
		"Red (R) or Black (B)?",
		"Higher (H) or Lower (L)",
		"Inside (I) or Outside(O)",
		"Suit? Diamonds (D), Hearts(H), Spade (S),Clubs (C) "}

	var answer string
	var hand deck
	guessedCorrectly := 0
	totalRounds := 0
	for {
		hand, d = deal(d, 4)
		for i, q := range questions {
			fmt.Println("The Question is:", q)
			fmt.Println("Your guess?")
			fmt.Scanln(&answer)

			result := validateAnswer(i, answer, hand)
			fmt.Println("Cards are...", i)
			for j := 0; j <= i; j++ {
				fmt.Println("[", hand[j].Suit, ",", hand[j].Value, "]")
			}

			if result {
				fmt.Println("CORRECT! 10 POINTS TO SLYTHERIN")
				guessedCorrectly += 1

			} else {
				fmt.Println("You have chosen..Poorly.")
			}
		}
		totalRounds += 4
		if len(d) < r.minCards() {
			printResults(guessedCorrectly, totalRounds)
			return d

		}
		fmt.Println("Play again? Y or N")
		fmt.Scanln(&answer)
		if strings.ToLower(answer) == "n" {
			printResults(guessedCorrectly, totalRounds)
			return d
		}

	}

}

func printResults(correct int, total int) {
	fmt.Println("Game Over! You answered", correct, "correctly out of", total, "total cards.")
}

func validateAnswer(q int, input string, d deck) bool {
	if q == 0 {
		//Pass one: Get card[0] and see if the suit matches what was guessed
		var suit string
		card := d[0]
		fmt.Println(card)
		if card.Suit == "Diamonds" || card.Suit == "Hearts" {
			suit = "r"
		} else {
			suit = "b"
		}
		fmt.Println(suit, strings.ToLower(input), suit == strings.ToLower(input))
		return suit == strings.ToLower(input)

	}
	if q == 1 {
		//Pass two: Get card[0] and card[1]
		//See if card 1 is higher or lower based on what was guessed
		previous := d[0]
		current := d[1]
		pVal := previous.Value
		cVal := current.Value

		if strings.ToLower(input) == "h" {
			res := cVal > pVal
			fmt.Println(res)
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

		var minCard card
		var maxCard card
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
	fmt.Println(val, current)

	return val == current.Suit
}
