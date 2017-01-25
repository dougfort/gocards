package gocards

// NewOrderedDeck creates a deck ordered by Suit and by Rank within suit
func NewOrderedDeck(def DecDef) Deck {
	var size int

	for _, suitSize := range def {
		size += int(suitSize)
	}

	deck := make(Deck, size)

	var i int
	for suit, suitSize := range def {
		for rank := 0; rank < int(suitSize); rank++ {
			deck[i] = Card{Suit(suit), Rank(rank)}
			i++
		}
	}

	return deck
}
