package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 28 {
		t.Errorf("Expected deck of length 20,got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Firs Card As ace of spades,bt got %s", d[0])
	}
}
