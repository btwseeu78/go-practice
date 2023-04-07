package main

import (
	"fmt"
	"subroutine"
)

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
	cards.saveToFile("testsave.txt")
	newDeck := newDeckFromFile("testsave.txt")
	newDeck.print()
	newDeck.shufle()
	newDeck.print()
	subroutine.TestLinks()
}
