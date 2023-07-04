package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type Card struct {
	Suit  string
	Value int
}

//Create a new type of deck
//Which is a slice of a strings
//Eventually we are going to upgrade this so that it is a
//Slice of a map {suit: is a constant string, number: 1-10 }
//Handle the face values as  1 (ACE) ints 11 (Jack) 12 (Queen) 13 (King)

type deck []Card

func emptyDeck() deck {
	return deck{}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	for _, suit := range cardSuits {
		for i := 1; i <= 13; i++ {
			cards = append(cards, Card{suit, i})
		}

	}
	return cards

}
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)

	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

func (d deck) toFile(filename string) error {
	out, err := json.Marshal(d)
	if err != nil {
		fmt.Println("Failed")
	}

	return ioutil.WriteFile(filename, out, 0666)

}

// func newDeckFromFile(filename string) deck {
// 	bs, err := ioutil.ReadFile(filename)
// 	fmt.Print(bs)
// 	if err != nil {
// 		//Option one: log error and give a new deck
// 		//Option two: raise an error
// 		fmt.Println("Error: ", err)
// 		os.Exit(1)
// 	}

// }

func (d deck) shuffleDeck() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		swap := r.Intn(len(d) - 1)
		d[i], d[swap] = d[swap], d[i]
	}

}
