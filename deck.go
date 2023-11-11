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
	cardSuits := []string{"Ouros", "Espadas", "Copas", "Paus"}
	cardValues := []string{"Ás", "Dois", "Três", "Quatro", "Cinco"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" de "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newIndex := r.Intn(len(d) - 1)
		d[i], d[newIndex] = d[newIndex], d[i]
	}
	return d
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ", ")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) deck {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(500)
	}
	return strings.Split(string(file), ", ")
}

func getFullDeckLen() int {
	d := newDeck()
	return len(d)
}
