package gocards

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

// DeckDefSize computes the defined size of a deck
func DeckDefSize(deckDef DeckDef) int {
	var totalSize int

	for _, size := range deckDef {
		totalSize += int(size)
	}

	return totalSize
}