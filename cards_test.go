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
		})
	}
}