package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a sice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
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
	return d[:handSize], d[handSize:] // 0 -> handSize - 1 , handSize -> end
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666) // 0666 - permission code
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		// Option 1 - log the error and return a call to newDeck()
		// Option 2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1) // Option 2
	}

	ss := strings.Split(string(bs), ",")
	return deck(ss)
}

func (d deck) shuffle() {
	// // Steps to get int64
	// // take the current time
	// t := time.Now()

	// // to generate int64 value
	// i64 := t.UnixNano()

	// // To create a new Source for randomization
	// source := rand.NewSource(i64)

	source := rand.NewSource(time.Now().UnixNano())

	// To create a new Random generator
	r := rand.New(source)

	for i := range d {
		// To generate a randome number from the r generator
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
