package gocards

import (
	"fmt"
	"testing"
)

func TestCards(t *testing.T) {
	testCases := []struct {
		def          DeckDef
		expectedSize int
	}{
		{DeckDef{}, 0},
		{DeckDef{1}, 1},
		{DeckDef{1, 2, 3}, 6},
		{DeckDef{13, 13, 13, 13}, 52},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.def), func(t *testing.T) {

			d := NewOrderedCards(tc.def)

			var prev Card
			var count int
			for i, card := range d {
				if i > 0 {
					if !prev.LessThan(card) {
						t.Fatalf("out of order: %s <= %s", card.Value(), prev.Value())
					}
				}
				count++
			}
			if count != tc.expectedSize {
				t.Fatalf("size mismatch: %d != %d", count, tc.expectedSize)
			}

			if DeckDefSize(tc.def) != tc.expectedSize {
				t.Fatalf("DeckDefSize mismatch: %d != %d",
					DeckDefSize(tc.def), tc.expectedSize)
			}

			sd1, seed := Shuffle(d)
			if len(sd1) != tc.expectedSize {
				t.Fatalf("sorted deck is wrong size %d != %d",
					len(sd1), tc.expectedSize)
			}

			if tc.expectedSize > 1 {

				// after shuffle, the cards should be in different places
				matchCount, err := MatchCount(d, sd1)
				if err != nil {
					t.Fatalf("MatchCount failed: %s", err)
				}
				if matchCount == tc.expectedSize {
					t.Fatalf("shuffled decks are identical")
				}

				// a shuffle with the same seed should produce the same order
				sd2 := SeededShuffle(d, seed)
				matchCount, err = MatchCount(sd1, sd2)
				if err != nil {
					t.Fatalf("MatchCount failed: %s", err)
				}
				if matchCount != tc.expectedSize {
					t.Fatalf("seeded shuffle is not identical")
				}

			}

		})
	}
}

func TestMatchCountError(t *testing.T) {
	deckDef1 := DeckDef{1}
	deckDef1a := DeckDef{1}
	deckDef2 := DeckDef{2}

	c1 := NewOrderedCards(deckDef1)
	c1a := NewOrderedCards(deckDef1a)
	c2 := NewOrderedCards(deckDef2)

	_, err := MatchCount(c1, c1a)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	_, err = MatchCount(c1, c2)
	if err == nil {
		t.Fatalf("expected error")
	}

}
