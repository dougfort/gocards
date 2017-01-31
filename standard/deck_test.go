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

	pd := d.PassThroughShuffle()
	if pd.Seed() != 0 {
		t.Fatalf("expected unshuffled deck; found: %d", pd.Seed())
	}

	sd := d.Shuffle()

	if sd.Seed() == 0 {
		t.Fatalf("expected shuffled deck; found: %d", sd.Seed())
	}

	ss := d.SeededShuffle(sd.Seed())

	if ss.Seed() != sd.Seed() {
		t.Fatalf("seed mismatch: %d != %d", ss.Seed(), sd.Seed())
	}

}

func TestDoubleleDeck(t *testing.T) {
	const size = 2

	d := NewDecks(size)

	expectedSize := size * gocards.DeckDefSize(deckDef)
	if d.Size() != expectedSize {
		t.Fatalf("size mismatch: expected %d found %d", expectedSize, d.Size())
	}

	pd := d.PassThroughShuffle()
	if pd.Seed() != 0 {
		t.Fatalf("expected unshuffled deck; found: %d", pd.Seed())
	}

	sd := d.Shuffle()

	if sd.Seed() == 0 {
		t.Fatalf("expected shuffled deck; found: %d", sd.Seed())
	}

	ss := d.SeededShuffle(sd.Seed())

	if ss.Seed() != sd.Seed() {
		t.Fatalf("seed mismatch: %d != %d", ss.Seed(), sd.Seed())
	}

}
