package standard

import (
	"github.com/dougfort/gocards"
)

// Suit values
const (
	Clubs gocards.Suit = iota
	Diamonds
	Hearts
	Spades
)

// Rank Values
const (
	Ace gocards.Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// deckDef defines the standard deck: 4 suits of thirteen cards each
var deckDef = gocards.DeckDef{13, 13, 13, 13}

type deckImpl struct {
	cards gocards.Cards
}

// NewDeck returns an entity that implements the Deck interface
func NewDeck() gocards.Deck {
	return &deckImpl{cards: gocards.NewOrderedCards(deckDef)}
}

// NewDecks returns an entity that implements the Deck interface
// it will consist of 'count' single decks appended
func NewDecks(count int) gocards.Deck {
	var impl deckImpl

	for i := 0; i < count; i++ {
		impl.cards = append(impl.cards, gocards.NewOrderedCards(deckDef)...)
	}

	return &impl
}

// Size returns the number of cards available from the Deck
func (i *deckImpl) Size() int {
	return len(i.cards)
}
