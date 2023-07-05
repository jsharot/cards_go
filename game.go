package main

import (
	"fmt"
	"os"
	"strings"
)

type game interface {
	//some method that all will define
	play(deck) deck
	name() string
	minCards() int
}

// Play the game!
func playGame(g game, d deck) deck {
	return g.play(d)
}

func validateGame(input string) game {
	if strings.ToLower(input) == "a" {
		return guessTheCard{}
	} else {
		//TODO: Add more games
		fmt.Println("Hmm.. I don't know any other games.")
		os.Exit(1)
		panic("Bye!")
	}

}
