package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	 cards := deck{}

	 cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	 cardValues := []string{"Ace", "Two", "Three", "Four"}
	 
	 for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, suit + " of " + value)
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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	file, error := os.ReadFile(fileName)

	if(error != nil) {
		fmt.Println("Filename is not correct")
	} 

		parsedFile := string(file)
		stringSlice := strings.Split(parsedFile, ",")
		deck := deck(stringSlice)

	return deck
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)



	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}