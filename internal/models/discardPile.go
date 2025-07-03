package models

type DiscardPile interface {
	AddCard(card Card)
	ResetDeck(deck Deck)
}

type discardPile struct {
	cards []Card
}

// NewDiscardPile creates a new discard pile
// Input: none
// Returns: a new DiscardPile
func NewDiscardPile() DiscardPile {
	return &discardPile{
		cards: make([]Card, 0),
	}
}

// AddCard adds a card to the discard pile
// Input: card - a Card
// Returns: none
func (d *discardPile) AddCard(card Card) {
	if card == nil {
		return
	}
	d.cards = append(d.cards, card)
}

// ResetDeck returns the cards to the deck and shuffles it
// Set the discard pile to empty
// Input: deck - a Deck
// Returns: none
func (d *discardPile) ResetDeck(deck Deck) {
	if deck == nil {
		return
	}
	deck.AddCardsToTop(d.cards)
	d.cards = make([]Card, 0)
	deck.Shuffle()
}
