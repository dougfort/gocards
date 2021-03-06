package standard

import (
	"fmt"
	"testing"

	"github.com/dougfort/gocards"
)

func TestSingleDeck(t *testing.T) {
	testCases := []int{1, 2, 3, 4, 5, 100}

	for _, count := range testCases {
		t.Run(fmt.Sprintf("deck count = %d", count), func(t *testing.T) {

			var d gocards.OrderedDeck
			if count == 1 {
				d = NewDeck()
			} else {
				d = NewDecks(count)
			}

			expectedSize := count * gocards.DeckDefSize(deckDef)
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

			if sd.RemainingCards() != d.Size() {
				t.Fatalf("size mismatch: %d != %d", sd.RemainingCards(), d.Size())
			}
			if ss.RemainingCards() != d.Size() {
				t.Fatalf("size mismatch: %d != %d", ss.RemainingCards(), d.Size())
			}

			cardCount := make(map[gocards.Card]int)
			for i := 0; i < d.Size(); i++ {
				cd, ok := sd.Next()
				if !ok {
					t.Fatalf("unexpected EOF")
				}

				cs, ok := ss.Next()
				if !ok {
					t.Fatalf("unexpected EOF")
				}

				if !cs.Equal(cd) {
					t.Fatalf("card mismatch: %v != %v", cs, cd)
				}

				cardCount[cd]++
			}

			if len(cardCount) != d.Size()/count {
				t.Fatalf("len(cardCount) (%d) != d.Size() (%d)", len(cardCount), d.Size())
			}

			for card, n := range cardCount {
				if n != count {
					t.Fatalf("invalid card count %s: %d, expected: %d",
						card.Value(), n, count)
				}
			}

			_, ok := sd.Next()
			if ok {
				t.Fatalf("expected EOF")
			}

			_, ok = ss.Next()
			if ok {
				t.Fatalf("expected EOF")
			}

		})
	}
}
