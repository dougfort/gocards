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

// Cards is a group of cards
type Cards []Card

// SuitSize defines the number of ranks in a suit
type SuitSize int

// DeckDef defines the content of a Deck of cards
type DeckDef []SuitSize
