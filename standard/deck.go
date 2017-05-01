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

const (
	// RuneBack represents the back of a(n unseen) Card displayed as a rune
	RuneBack = '🂠'

	// StringBack represents the back of a(n unseen) Card displayed as a string
	StringBack = "(.....)"
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
		gocards.Card{Suit: Clubs, Rank: Ace}:      "(C,  A)",
		gocards.Card{Suit: Clubs, Rank: Two}:      "(C,  2)",
		gocards.Card{Suit: Clubs, Rank: Three}:    "(C,  3)",
		gocards.Card{Suit: Clubs, Rank: Four}:     "(C,  4)",
		gocards.Card{Suit: Clubs, Rank: Five}:     "(C,  5)",
		gocards.Card{Suit: Clubs, Rank: Six}:      "(C,  6)",
		gocards.Card{Suit: Clubs, Rank: Seven}:    "(C,  7)",
		gocards.Card{Suit: Clubs, Rank: Eight}:    "(C,  8)",
		gocards.Card{Suit: Clubs, Rank: Nine}:     "(C,  9)",
		gocards.Card{Suit: Clubs, Rank: Ten}:      "(C, 10)",
		gocards.Card{Suit: Clubs, Rank: Jack}:     "(C,  J)",
		gocards.Card{Suit: Clubs, Rank: Queen}:    "(C,  Q)",
		gocards.Card{Suit: Clubs, Rank: King}:     "(C,  K)",
		gocards.Card{Suit: Diamonds, Rank: Ace}:   "(D,  A)",
		gocards.Card{Suit: Diamonds, Rank: Two}:   "(D,  2)",
		gocards.Card{Suit: Diamonds, Rank: Three}: "(D,  3)",
		gocards.Card{Suit: Diamonds, Rank: Four}:  "(D,  4)",
		gocards.Card{Suit: Diamonds, Rank: Five}:  "(D,  5)",
		gocards.Card{Suit: Diamonds, Rank: Six}:   "(D,  6)",
		gocards.Card{Suit: Diamonds, Rank: Seven}: "(D,  7)",
		gocards.Card{Suit: Diamonds, Rank: Eight}: "(D,  8)",
		gocards.Card{Suit: Diamonds, Rank: Nine}:  "(D,  9)",
		gocards.Card{Suit: Diamonds, Rank: Ten}:   "(D, 10)",
		gocards.Card{Suit: Diamonds, Rank: Jack}:  "(D,  J)",
		gocards.Card{Suit: Diamonds, Rank: Queen}: "(D,  Q)",
		gocards.Card{Suit: Diamonds, Rank: King}:  "(D,  K)",
		gocards.Card{Suit: Hearts, Rank: Ace}:     "(H,  A)",
		gocards.Card{Suit: Hearts, Rank: Two}:     "(H,  2)",
		gocards.Card{Suit: Hearts, Rank: Three}:   "(H,  3)",
		gocards.Card{Suit: Hearts, Rank: Four}:    "(H,  4)",
		gocards.Card{Suit: Hearts, Rank: Five}:    "(H,  5)",
		gocards.Card{Suit: Hearts, Rank: Six}:     "(H,  6)",
		gocards.Card{Suit: Hearts, Rank: Seven}:   "(H,  7)",
		gocards.Card{Suit: Hearts, Rank: Eight}:   "(H,  8)",
		gocards.Card{Suit: Hearts, Rank: Nine}:    "(H,  9)",
		gocards.Card{Suit: Hearts, Rank: Ten}:     "(H, 10)",
		gocards.Card{Suit: Hearts, Rank: Jack}:    "(H,  J)",
		gocards.Card{Suit: Hearts, Rank: Queen}:   "(H,  Q)",
		gocards.Card{Suit: Hearts, Rank: King}:    "(H,  K)",
		gocards.Card{Suit: Spades, Rank: Ace}:     "(S,  A)",
		gocards.Card{Suit: Spades, Rank: Two}:     "(S,  2)",
		gocards.Card{Suit: Spades, Rank: Three}:   "(S,  3)",
		gocards.Card{Suit: Spades, Rank: Four}:    "(S,  4)",
		gocards.Card{Suit: Spades, Rank: Five}:    "(S,  5)",
		gocards.Card{Suit: Spades, Rank: Six}:     "(S,  6)",
		gocards.Card{Suit: Spades, Rank: Seven}:   "(S,  7)",
		gocards.Card{Suit: Spades, Rank: Eight}:   "(S,  8)",
		gocards.Card{Suit: Spades, Rank: Nine}:    "(S,  9)",
		gocards.Card{Suit: Spades, Rank: Ten}:     "(S, 10)",
		gocards.Card{Suit: Spades, Rank: Jack}:    "(S,  J)",
		gocards.Card{Suit: Spades, Rank: Queen}:   "(S,  Q)",
		gocards.Card{Suit: Spades, Rank: King}:    "(S,  K)",
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
