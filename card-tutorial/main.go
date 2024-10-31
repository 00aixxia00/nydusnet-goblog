package main

import "fmt"

var decksize int

func main() {
	cards := []string{newCard(), newCard()}
	cards = append(cards, newCard())

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Ace of Spades"
}
