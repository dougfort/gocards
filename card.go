package gocards

import "fmt"

// Equal returns true if the cards have the same value
func (c Card) Equal(x Card) bool {
	return c.Suit == x.Suit && c.Rank == x.Rank
}

// LessThan returns true if the card is lower in value
func (c Card) LessThan(x Card) bool {
	return (c.Suit < x.Suit) || (c.Suit == x.Suit && c.Rank < x.Rank)
}

// Value displays the inner value of a card
func (c Card) Value() string {
	return fmt.Sprintf("(%d, %d)", c.Rank, c.Suit)
}
