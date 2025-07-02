package models

import "math/rand"

// Deck is a collection of cards that can be drawn from and shuffled
type Deck interface {
	DrawCard() Card
	Shuffle()
	AddCardsToTop(cards []Card)
}

// deck is an implementation of the Deck interface
type deck struct {
	cards []Card
}

// NewDeck creates a new deck of cards
// Input: cards - a slice of Cards
// Returns: a new Deck
func NewDeck(cards []Card) Deck {
	return &deck{
		cards: cards,
	}
}

// DrawCard draws a card from the top of the deck
// Returns nil if the deck is empty
// Input: none
// Returns: a Card
func (d *deck) DrawCard() Card {
	if len(d.cards) == 0 {
		return nil
	}
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}

// Shuffle shuffles the deck
// Input: none
// Returns: none
func (d *deck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

// AddCardsToTop adds cards to the top of the deck
// Input: cards - a slice of Cards
// Returns: none
func (d *deck) AddCardsToTop(cards []Card) {
	d.cards = append(cards, d.cards...)
}
