package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// custom type declaration
type deck []string

// new deck generater func
func newDeck() deck {
	cards := deck{}

	cardName := []string{"spades", "diamonds", "hearts", "clubs"}
	cardValue := []string{"ace", "two", "three", "four", "five"}

	for _, cardN := range cardName {
		for _, cardV := range cardValue {
			cards = append(cards, cardN+" of "+cardV)
		}
	}

	return cards
}

// print method for custom type object (deck)
func (d deck) deckPrinter() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// creating a function which return a specific deck within a given range
func deal(d deck, _range int) (deck, deck) {
	return d[:_range], d[_range:]
}

func (d deck) deckToString() string {
	return strings.Join(d, "\n")
}

func (d deck) saveDeckToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.deckToString()), 0666)
}

func readDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	dataFromFile := strings.Split(string(bs), "\n")

	return deck(dataFromFile)
}
