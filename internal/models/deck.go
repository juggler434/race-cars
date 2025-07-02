package models

import "math/rand"

type Deck interface {
	DrawCard() Card
	Shuffle()
	AddCardsToTop(cards []Card)
}
type deck struct {
	cards []Card
}

func NewDeck(cards []Card) Deck {
	return &deck{
		cards: cards,
	}
}

func (d *deck) DrawCard() Card {
	if len(d.cards) == 0 {
		return nil
	}
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}

func (d *deck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *deck) AddCardsToTop(cards []Card) {
	d.cards = append(cards, d.cards...)
}
