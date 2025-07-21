package models

import "errors"

// Hand is a collection of cards that a player can draw from and discard to the discard pile
type Hand interface {
	AddCards(cards []Card)
	DrawCard(deck Deck)
	DiscardCard(index int, discardPile DiscardPile) error
	PlayCard(index int) (Card, error)
}

type hand struct {
	cards []Card
}

// NewHand creates a new hand
// Input: none
// Returns: a new Hand
func NewHand() Hand {
	return &hand{
		cards: make([]Card, 0),
	}
}

// AddCards adds cards to the hand
// Input: cards - a slice of Cards
// Returns: none
func (h *hand) AddCards(cards []Card) {
	if cards == nil {
		return
	}
	h.cards = append(h.cards, cards...)
}

// DrawCard draws a card from the deck
// Input: deck - a Deck
// Returns: none
func (h *hand) DrawCard(deck Deck) {
	if deck == nil {
		return
	}

	card := deck.DrawCard()
	if card == nil {
		return
	}

	h.cards = append(h.cards, card)
}

// DiscardCard discards a card from the hand
// Input: index - an int, the index of the card to discard
// Returns: an error if the card is not discardable or the index is out of bounds
func (h *hand) DiscardCard(index int, discardPile DiscardPile) error {
	if index < 0 || index >= len(h.cards) {
		return errors.New("invalid card index")
	}

	if h.cards[index] == nil {
		return errors.New("card is nil")
	}

	if !h.cards[index].IsDiscardable() {
		return errors.New("card is not discardable")
	}

	if discardPile == nil {
		return errors.New("discard pile is nil")
	}

	discardPile.AddCard(h.cards[index])
	h.cards = append(h.cards[:index], h.cards[index+1:]...)
	return nil
}

// PlayCard plays a card from the hand
// Input: index - an int, the index of the card to play
// Returns: Card at the position of index, anerror if the card is not playable or the index is out of bounds
func (h *hand) PlayCard(index int) (Card, error) {
	if index < 0 || index >= len(h.cards) {
		return nil, errors.New("invalid card index")
	}

	if h.cards[index] == nil {
		return nil, errors.New("card is nil")
	}

	if !h.cards[index].IsPlayable() {
		return nil, errors.New("card is not playable")
	}

	card := h.cards[index]

	h.cards = append(h.cards[:index], h.cards[index+1:]...)
	return card, nil
}
