package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//Create a new type of 'deck' which is a slice of strings, defining a new type (deck) that extends an existing type (string)
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for i := 0; i < len(cardSuits); i++ {
		for j := 0; j < len(cardValues); j++ {
			cards = append(cards, cardValues[j]+" of "+cardSuits[i])
		}
	}

	return cards
}

//reciever function - any variable of type of deck now gets access to the print method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//function will return two values, both of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//reciever function that takes a deck and converts it into one complete string seperated by commas
func (d deck) toString() string {
	return strings.Join([]string(d), ",") //uses a function from strings package

}

//writes to file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
