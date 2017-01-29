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
				if matchCount(t, d, sd1) == tc.expectedSize {
					t.Fatalf("shuffled decks are identical")
				}

				// a shuffle with the same seed should produce the same order
				sd2 := SeededShuffle(d, seed)
				if matchCount(t, sd1, sd2) != tc.expectedSize {
					t.Fatalf("seeded shuffle is not identical")
				}

			}

		})
	}
}

// matchCount returns the number of matches between to slices of cards
func matchCount(t *testing.T, c1 Cards, c2 Cards) int {
	if len(c1) != len(c2) {
		t.Fatalf("matchCount: size mismatch: %d != %d", len(c1), len(c2))
	}

	var count int

	for i := 0; i < len(c1); i++ {
		if c1[i].Equal(c2[i]) {
			count++
		}
	}

	return count
}
