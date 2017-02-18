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

const (
	// RuneBack represents the back of a(n unseen) Card displayed as a rune
	RuneBack = '🂠'

	// StringBack represents the back of a(n unseen) Card displayed as a string
	StringBack = "()"
)

// RuneMap maps cards to unicode images
type RuneMap map[gocards.Card]rune

// StringMap maps cards to string names
type StringMap map[gocards.Card]string

var (

	// Runes are Cards represented as unicode runes
	Runes = RuneMap{
		gocards.Card{Suit: Clubs, Rank: Ace}:      '🃑',
		gocards.Card{Suit: Clubs, Rank: Two}:      '🃒',
		gocards.Card{Suit: Clubs, Rank: Three}:    '🃓',
		gocards.Card{Suit: Clubs, Rank: Four}:     '🃔',
		gocards.Card{Suit: Clubs, Rank: Five}:     '🃕',
		gocards.Card{Suit: Clubs, Rank: Six}:      '🃖',
		gocards.Card{Suit: Clubs, Rank: Seven}:    '🃗',
		gocards.Card{Suit: Clubs, Rank: Eight}:    '🃘',
		gocards.Card{Suit: Clubs, Rank: Nine}:     '🃙',
		gocards.Card{Suit: Clubs, Rank: Ten}:      '🃚',
		gocards.Card{Suit: Clubs, Rank: Jack}:     '🃛',
		gocards.Card{Suit: Clubs, Rank: Queen}:    '🃝',
		gocards.Card{Suit: Clubs, Rank: King}:     '🃞',
		gocards.Card{Suit: Diamonds, Rank: Ace}:   '🃁',
		gocards.Card{Suit: Diamonds, Rank: Two}:   '🃂',
		gocards.Card{Suit: Diamonds, Rank: Three}: '🃃',
		gocards.Card{Suit: Diamonds, Rank: Four}:  '🃄',
		gocards.Card{Suit: Diamonds, Rank: Five}:  '🃅',
		gocards.Card{Suit: Diamonds, Rank: Six}:   '🃆',
		gocards.Card{Suit: Diamonds, Rank: Seven}: '🃇',
		gocards.Card{Suit: Diamonds, Rank: Eight}: '🃈',
		gocards.Card{Suit: Diamonds, Rank: Nine}:  '🃉',
		gocards.Card{Suit: Diamonds, Rank: Ten}:   '🃊',
		gocards.Card{Suit: Diamonds, Rank: Jack}:  '🃋',
		gocards.Card{Suit: Diamonds, Rank: Queen}: '🃍',
		gocards.Card{Suit: Diamonds, Rank: King}:  '🃎',
		gocards.Card{Suit: Hearts, Rank: Ace}:     '🂱',
		gocards.Card{Suit: Hearts, Rank: Two}:     '🂲',
		gocards.Card{Suit: Hearts, Rank: Three}:   '🂳',
		gocards.Card{Suit: Hearts, Rank: Four}:    '🂴',
		gocards.Card{Suit: Hearts, Rank: Five}:    '🂵',
		gocards.Card{Suit: Hearts, Rank: Six}:     '🂶',
		gocards.Card{Suit: Hearts, Rank: Seven}:   '🂷',
		gocards.Card{Suit: Hearts, Rank: Eight}:   '🂸',
		gocards.Card{Suit: Hearts, Rank: Nine}:    '🂹',
		gocards.Card{Suit: Hearts, Rank: Ten}:     '🂺',
		gocards.Card{Suit: Hearts, Rank: Jack}:    '🂻',
		gocards.Card{Suit: Hearts, Rank: Queen}:   '🂽',
		gocards.Card{Suit: Hearts, Rank: King}:    '🂾',
		gocards.Card{Suit: Spades, Rank: Ace}:     '🂡',
		gocards.Card{Suit: Spades, Rank: Two}:     '🂢',
		gocards.Card{Suit: Spades, Rank: Three}:   '🂣',
		gocards.Card{Suit: Spades, Rank: Four}:    '🂤',
		gocards.Card{Suit: Spades, Rank: Five}:    '🂥',
		gocards.Card{Suit: Spades, Rank: Six}:     '🂦',
		gocards.Card{Suit: Spades, Rank: Seven}:   '🂧',
		gocards.Card{Suit: Spades, Rank: Eight}:   '🂨',
		gocards.Card{Suit: Spades, Rank: Nine}:    '🂩',
		gocards.Card{Suit: Spades, Rank: Ten}:     '🂪',
		gocards.Card{Suit: Spades, Rank: Jack}:    '🂫',
		gocards.Card{Suit: Spades, Rank: Queen}:   '🂭',
		gocards.Card{Suit: Spades, Rank: King}:    '🂮',
	}

	// Strings are cards represented as stringd
	Strings = StringMap{
		gocards.Card{Suit: Clubs, Rank: Ace}:      "(Clubs, Ace)",
		gocards.Card{Suit: Clubs, Rank: Two}:      "(Clubs, 2)",
		gocards.Card{Suit: Clubs, Rank: Three}:    "(Clubs, 3)",
		gocards.Card{Suit: Clubs, Rank: Four}:     "(Clubs, 4)",
		gocards.Card{Suit: Clubs, Rank: Five}:     "(Clubs, 5)",
		gocards.Card{Suit: Clubs, Rank: Six}:      "(Clubs, 6)",
		gocards.Card{Suit: Clubs, Rank: Seven}:    "(Clubs, 7)",
		gocards.Card{Suit: Clubs, Rank: Eight}:    "(Clubs, 8)",
		gocards.Card{Suit: Clubs, Rank: Nine}:     "(Clubs, 9)",
		gocards.Card{Suit: Clubs, Rank: Ten}:      "(Clubs, 10)",
		gocards.Card{Suit: Clubs, Rank: Jack}:     "(Clubs, Jack)",
		gocards.Card{Suit: Clubs, Rank: Queen}:    "(Clubs, Queen)",
		gocards.Card{Suit: Clubs, Rank: King}:     "(Clubs, King)",
		gocards.Card{Suit: Diamonds, Rank: Ace}:   "(Diamonds, Ace)",
		gocards.Card{Suit: Diamonds, Rank: Two}:   "(Diamonds, 2)",
		gocards.Card{Suit: Diamonds, Rank: Three}: "(Diamonds, 3)",
		gocards.Card{Suit: Diamonds, Rank: Four}:  "(Diamonds, 4)",
		gocards.Card{Suit: Diamonds, Rank: Five}:  "(Diamonds, 5)",
		gocards.Card{Suit: Diamonds, Rank: Six}:   "(Diamonds, 6)",
		gocards.Card{Suit: Diamonds, Rank: Seven}: "(Diamonds, 7)",
		gocards.Card{Suit: Diamonds, Rank: Eight}: "(Diamonds, 8)",
		gocards.Card{Suit: Diamonds, Rank: Nine}:  "(Diamonds, 9)",
		gocards.Card{Suit: Diamonds, Rank: Ten}:   "(Diamonds, 10)",
		gocards.Card{Suit: Diamonds, Rank: Jack}:  "(Diamonds, Jack)",
		gocards.Card{Suit: Diamonds, Rank: Queen}: "(Diamonds, Queen)",
		gocards.Card{Suit: Diamonds, Rank: King}:  "(Diamonds, King)",
		gocards.Card{Suit: Hearts, Rank: Ace}:     "(Hearts, Ace)",
		gocards.Card{Suit: Hearts, Rank: Two}:     "(Hearts, 2)",
		gocards.Card{Suit: Hearts, Rank: Three}:   "(Hearts, 3)",
		gocards.Card{Suit: Hearts, Rank: Four}:    "(Hearts, 4)",
		gocards.Card{Suit: Hearts, Rank: Five}:    "(Hearts, 5)",
		gocards.Card{Suit: Hearts, Rank: Six}:     "(Hearts, 6)",
		gocards.Card{Suit: Hearts, Rank: Seven}:   "(Hearts, 7)",
		gocards.Card{Suit: Hearts, Rank: Eight}:   "(Hearts, 8)",
		gocards.Card{Suit: Hearts, Rank: Nine}:    "(Hearts, 9)",
		gocards.Card{Suit: Hearts, Rank: Ten}:     "(Hearts, 10)",
		gocards.Card{Suit: Hearts, Rank: Jack}:    "(Hearts, Jack)",
		gocards.Card{Suit: Hearts, Rank: Queen}:   "(Hearts, Queen)",
		gocards.Card{Suit: Hearts, Rank: King}:    "(Hearts, King)",
		gocards.Card{Suit: Spades, Rank: Ace}:     "(Spades, Ace)",
		gocards.Card{Suit: Spades, Rank: Two}:     "(Spades, 2)",
		gocards.Card{Suit: Spades, Rank: Three}:   "(Spades, 3)",
		gocards.Card{Suit: Spades, Rank: Four}:    "(Spades, 4)",
		gocards.Card{Suit: Spades, Rank: Five}:    "(Spades, 5)",
		gocards.Card{Suit: Spades, Rank: Six}:     "(Spades, 6)",
		gocards.Card{Suit: Spades, Rank: Seven}:   "(Spades, 7)",
		gocards.Card{Suit: Spades, Rank: Eight}:   "(Spades, 8)",
		gocards.Card{Suit: Spades, Rank: Nine}:    "(Spades, 9)",
		gocards.Card{Suit: Spades, Rank: Ten}:     "(Spades, 10)",
		gocards.Card{Suit: Spades, Rank: Jack}:    "(Spades, Jack)",
		gocards.Card{Suit: Spades, Rank: Queen}:   "(Spades, Queen)",
		gocards.Card{Suit: Spades, Rank: King}:    "(Spades, King)",
	}
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
