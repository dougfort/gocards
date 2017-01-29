package gocards

import (
	"math/rand"
	"time"
)

// NewOrderedCards creates a Cards ordered by Suit and by Rank within suit
func NewOrderedCards(def DeckDef) Cards {
	var size int

	for _, suitSize := range def {
		size += int(suitSize)
	}

	cards := make(Cards, size)

	var i int
	for suit, suitSize := range def {
		for rank := 0; rank < int(suitSize); rank++ {
			cards[i] = Card{Suit(suit), Rank(rank)}
			i++
		}
	}

	return cards
}

// Shuffle randomly reorders the Cards
// This is Seeded shuffle with the seed set to the current time (returned)
func Shuffle(orig Cards) (Cards, int64) {
	seed := time.Now().UTC().UnixNano()
	return SeededShuffle(orig, seed), seed
}

// SeededShuffle reorders the Cards to the sequence determined by seed
// We expect the result to be the same every tie for a given seed
func SeededShuffle(orig Cards, seed int64) Cards {
	src := rand.NewSource(seed)
	r := rand.New(src)
	size := len(orig)
	perm := r.Perm(size)

	result := make(Cards, size)
	for i, j := range perm {
		result[i] = orig[j]
	}

	return result
}

// DeckDefSize computes the defined size of a deck
func DeckDefSize(deckDef DeckDef) int {
	var totalSize int

	for _, size := range deckDef {
		totalSize += int(size)
	}

	return totalSize
}
