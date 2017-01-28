package standard

import (
	"testing"

	"github.com/dougfort/gocards"
)

func TestSingleDeck(t *testing.T) {
	d := NewDeck()
	expectedSize := gocards.DeckDefSize(deckDef)
	if d.Size() != expectedSize {
		t.Fatalf("size mismatch: expected %d found %d", expectedSize, d.Size())
	}

}

func TestDoubleleDeck(t *testing.T) {
	const size = 2
	d := NewDecks(size)
	expectedSize := size * gocards.DeckDefSize(deckDef)
	if d.Size() != expectedSize {
		t.Fatalf("size mismatch: expected %d found %d", expectedSize, d.Size())
	}

}
