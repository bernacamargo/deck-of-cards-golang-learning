package main

func main() {
	cards := newDeck()
	cards = cards.shuffle()
	cards.print()
	//err := cards.saveToFile("fulldeck.txt")
	//if err != nil {
	//	return
	//}
	//loadFromFile := readFromFile("fulldeck.txt")
	//loadFromFile.print()

	//firstHand, remainingCards := deal(cards, 3)
	//fmt.Println("\nFirst hand:\n")
	//firstHand.print()
	//fmt.Println("\nRemaining cards:\n")
	//remainingCards.print()
}
