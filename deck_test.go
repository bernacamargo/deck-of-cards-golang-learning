package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	lenExpected := 20
	if len(d) != lenExpected {
		t.Errorf("Expected deck length of %v, but got %v", lenExpected, len(d))
	}

	expectedFirstCard := "Ás de Ouros"
	if d[0] != expectedFirstCard {
		t.Errorf("Expected first card %v but got %v", expectedFirstCard, d[0])
	}

	expectedLastCard := "Cinco de Paus"
	if d[len(d)-1] != expectedLastCard {
		t.Errorf("Expected last card %v but got %v", expectedLastCard, d[len(d)-1])
	}
}

func TestSaveDeckToFileAndReadDeckFromFile(t *testing.T) {
	fileNameDeckTesting := "_decktesting"

	d := newDeck()
	err := d.saveToFile(fileNameDeckTesting)
	if err != nil {
		fmt.Println("Error when saving file: ", err)
		os.Exit(500)
	}

	loadedDeck := readFromFile(fileNameDeckTesting)

	expectedLenCards := 20
	if len(loadedDeck) != expectedLenCards {
		t.Errorf("Expected %v cards, but got %v", expectedLenCards, len(loadedDeck))
	}

	err = os.Remove(fileNameDeckTesting)
	if err != nil {
		fmt.Println("Error when removing file: ", err)
		os.Exit(500)
	}
}
