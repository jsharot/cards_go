package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

//Card type  associations

type card struct {
	Suit  string
	Value int
}

var inputMappings = map[string]string{
	"h": "Hearts",
	"d": "Diamonds",
	"s": "Spades",
	"c": "Clubs",
}

// Deck Type
type deck []card

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	for _, suit := range cardSuits {
		for i := 1; i <= 13; i++ {
			cards = append(cards, card{suit, i})
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

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	//fmt.Print(bs)
	if err != nil {
		//Option one: log error and give a new deck
		//Option two: raise an error
		//fmt.Println("Error: ", err)
		return nil
	}
	fDeck := deck{}
	err2 := json.Unmarshal(bs, &fDeck)
	if err2 != nil {
		//fmt.Println("Error:", err)
		return nil

	}
	return fDeck
}

func (d deck) shuffleDeck() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		swap := r.Intn(len(d) - 1)
		d[i], d[swap] = d[swap], d[i]
	}

}
