package gocards

// Suit represents the suit of a card
type Suit uint8

// Rank represents the order within a Suit
type Rank uint8

// Card represents a single card
type Card struct {
	Suit Suit
	Rank Rank
}

// Deck is a group of cards
type Deck []Card

// SuitSize defines the number of ranks in a suit
type SuitSize int

// DecDef defines the content of a deck of cards
type DecDef []SuitSize
