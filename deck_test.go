package gocards

import (
	"fmt"
	"testing"
)

func TestDeck(t *testing.T) {
	testCases := []struct {
		def          DecDef
		expectedSize int
	}{
		{DecDef{}, 0},
		{DecDef{1}, 1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.def), func(t *testing.T) {
			d := NewOrderedDeck(tc.def)
			var prev Card
			var count int
			for i, card := range d {
				if i > 0 {
					if !card.LessThan(prev) {
						t.Fatalf("out of order: %s <= %s", card.Value(), prev.Value())
					}
				}
				count++
			}
			if count != tc.expectedSize {
				t.Fatalf("size mismatch: %d != %d", count, tc.expectedSize)
			}
		})
	}
}
