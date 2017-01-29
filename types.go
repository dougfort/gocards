package gocards

// Version is the reease vrsion
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

// Deck defines the interface that card definitons must support
type Deck interface {

	// Size returns the number of cards available from the Deck
	Size() int

	// Shuffle randomly reorders the Cards
	// This is Seeded shuffle wiht the seed set to the current time
	Shuffle()

	// SeededShuffle reorders the Cards to the sequence determined by seed
	// We expect the result to be the same every tie for a given seed
	SeededShuffle(seed int64)
}
