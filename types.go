package gocards

// Version is the release vrsion
const Version = "0.0.1"

// Suit represents the suit of a card
type Suit uint8

// Rank represents the order within a Suit
type Rank uint8

// Card represents a single card
type Card struct {
	Suit Suit
	Rank Rank
}

// Cards is a group of cards
type Cards []Card

// SuitSize defines the number of ranks in a suit
type SuitSize int

// DeckDef defines the content of a Deck of cards
type DeckDef []SuitSize

// OrderedDeck defines the interface to a newly created deck
type OrderedDeck interface {

	// Size returns the number of cards available from the Deck
	Size() int

	// Shuffle randomly reorders the Cards
	// This is Seeded shuffle with the seed set to the current time
	Shuffle() PlayableDeck

	// SeededShuffle reorders the Cards to the sequence determined by seed.
	// We expect the result to be the same every time for a given seed.
	SeededShuffle(seed int64) PlayableDeck

	// PassThroughShuffle makes the deck playable without shuffle
	PassThroughShuffle() PlayableDeck
}

// PlayableDeck defines the interface to a deck that is 'in play'
// You can't shuffle a deck when it's in play
type PlayableDeck interface {

	// RemainingCards returns the number of cards available from the Deck
	RemainingCards() int

	// Seed returns the random seed used to uniquely identify the game
	// Seed == 0 means the cars haven't been shuffled
	Seed() int64

	// Next returns the next available card and true if there is such a card.
	// returns false if all cards have been consumed.
	Next() (Card, bool)
}
