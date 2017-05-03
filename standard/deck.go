package standard

import (
	"github.com/dougfort/gocards"
)

// Suit values
const (
	NullSuit gocards.Suit = iota
	Clubs
	Diamonds
	Hearts
	Spades
)

// Rank Values
const (
	NullRank gocards.Rank = iota
	Ace
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

type orderedImpl struct {
	cards gocards.Cards
}

type playableImpl struct {
	cards gocards.Cards

	// seed uniquely identifies the game
	seed int64

	// next is the index to th next available card
	next int
}

// NewDeck returns an entity that implements the Deck interface
func NewDeck() gocards.OrderedDeck {
	return orderedImpl{cards: gocards.NewOrderedCards(deckDef)}
}

// NewDecks returns an entity that implements the Deck interface
// it will consist of 'count' single decks appended
func NewDecks(count int) gocards.OrderedDeck {
	var impl orderedImpl

	for i := 0; i < count; i++ {
		impl.cards = append(impl.cards, gocards.NewOrderedCards(deckDef)...)
	}

	return impl
}

// Size returns the number of cards available from the Deck
func (i orderedImpl) Size() int {
	return len(i.cards)
}

// Shuffle randomly reorders the Cards
// This is Seeded shuffle with the seed set to the current time
func (i orderedImpl) Shuffle() gocards.PlayableDeck {
	var p playableImpl
	p.cards, p.seed = gocards.Shuffle(i.cards)
	return &p
}

// SeededShuffle reorders the Cards to the sequence determined by seed
// We expect the result to be the same every tie for a given seed
func (i orderedImpl) SeededShuffle(seed int64) gocards.PlayableDeck {
	var p playableImpl
	p.seed = seed
	p.cards = gocards.SeededShuffle(i.cards, p.seed)
	return &p
}

// PassThroughShuffle makes the deck layable without shuffle
func (i orderedImpl) PassThroughShuffle() gocards.PlayableDeck {
	return &playableImpl{cards: i.cards}
}

// RemainingCards returns the number of cards available from the Deck
func (p *playableImpl) RemainingCards() int {
	return len(p.cards) - p.next
}

// Seed returns the random seed used to uniquely identify the game
// Seed == 0 means the cars haven't been shuffled
func (p *playableImpl) Seed() int64 {
	return p.seed
}

// Next returns the next available card and true if there is such a card.
// returns false if all cards ave been consumed.
func (p *playableImpl) Next() (gocards.Card, bool) {
	if p.next >= len(p.cards) {
		return gocards.Card{}, false
	}

	i := p.next
	p.next++

	return p.cards[i], true
}
