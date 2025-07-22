package models

// Player represents a player in the racing game
// A player has a name, car, deck, hand, discard pile, and manages played cards and icons
type Player interface {
	// GetName returns the player's name
	// Returns: the player's name as a string
	GetName() string

	// GetCar returns the player's car
	// Returns: the player's car instance
	GetCar() Car

	// GetDiscardPile returns the player's discard pile
	// Returns: the player's discard pile instance
	GetDiscardPile() DiscardPile

	// GetDeck returns the player's deck
	// Returns: the player's deck instance
	GetDeck() Deck

	// GetHand returns the player's hand
	// Returns: the player's hand instance
	GetHand() Hand

	// DrawCard draws a card from the specified deck into the player's hand
	// Input: deck - the deck to draw from
	// Returns: none
	DrawCard(deck Deck)

	// DiscardCard discards a card from the player's hand to their discard pile
	// Input: index - the index of the card in the hand to discard
	// Returns: an error if the card cannot be discarded
	DiscardCard(index int) error

	// GetIcons returns the player's current accumulated icons
	// Returns: a map of icon types to their counts
	GetIcons() map[Icon]int

	// AddIcons adds icons to the player's accumulated icon count
	// Input: icons - a map of icon types to counts to add
	// Returns: none
	AddIcons(icons map[Icon]int)

	// PlayCard plays a card from the player's hand and adds it to played cards
	// Input: index - the index of the card in the hand to play
	// Returns: an error if the card cannot be played
	PlayCard(index int) error

	// ResolvePlayedCards processes all played cards, calculating speed and adding icons
	// Also handles special cards like Stress cards
	// Returns: none
	ResolvePlayedCards()
}

type player struct {
	name        string
	car         Car
	discardPile DiscardPile
	deck        Deck
	hand        Hand
	playedCards []Card
	icons       map[Icon]int
}

// NewPlayer creates a new player instance
// Input: name - the player's name
//
//	car - the player's car
//	discardPile - the player's discard pile
//	deck - the player's deck
//	hand - the player's hand
//
// Returns: a new Player instance
func NewPlayer(name string, car Car, discardPile DiscardPile, deck Deck, hand Hand) Player {
	return &player{
		name:        name,
		car:         car,
		discardPile: discardPile,
		deck:        deck,
		hand:        hand,
		playedCards: make([]Card, 0),
		icons:       make(map[Icon]int),
	}
}

// GetName returns the player's name
// Input: none
// Returns: the player's name as a string
func (p *player) GetName() string {
	return p.name
}

// GetCar returns the player's car
// Input: none
// Returns: the player's car instance
func (p *player) GetCar() Car {
	return p.car
}

// GetDiscardPile returns the player's discard pile
// Input: none
// Returns: the player's discard pile instance
func (p *player) GetDiscardPile() DiscardPile {
	return p.discardPile
}

// GetDeck returns the player's deck
// Input: none
// Returns: the player's deck instance
func (p *player) GetDeck() Deck {
	return p.deck
}

// GetHand returns the player's hand
// Input: none
// Returns: the player's hand instance
func (p *player) GetHand() Hand {
	return p.hand
}

// DrawCard draws a card from the specified deck into the player's hand
// Input: deck - the deck to draw from
// Returns: none
func (p *player) DrawCard(deck Deck) {
	p.hand.DrawCard(deck)
}

// DiscardCard discards a card from the player's hand to their discard pile
// Input: index - the index of the card in the hand to discard
// Returns: an error if the card cannot be discarded
func (p *player) DiscardCard(index int) error {
	err := p.hand.DiscardCard(index, p.discardPile)
	if err != nil {
		return err
	}
	return nil
}

// GetIcons returns the player's current accumulated icons
// Input: none
// Returns: a map of icon types to their counts
func (p *player) GetIcons() map[Icon]int {
	return p.icons
}

// PlayCard plays a card from the player's hand and adds it to played cards
// Input: index - the index of the card in the hand to play
// Returns: an error if the card cannot be played
func (p *player) PlayCard(index int) error {
	card, err := p.hand.PlayCard(index)
	if err != nil {
		return err
	}

	p.playedCards = append(p.playedCards, card)

	return nil
}

// AddIcons adds icons to the player's accumulated icon count
// Input: icons - a map of icon types to counts to add
// Returns: none
func (p *player) AddIcons(icons map[Icon]int) {
	for icon, count := range icons {
		p.icons[icon] += count
	}
}

// ResolvePlayedCards processes all played cards, calculating speed and adding icons
// Also handles special cards like Stress cards
// Input: none
// Returns: none
func (p *player) ResolvePlayedCards() {
	speed := 0

	for _, card := range p.playedCards {
		if card.GetName() == Stress {
			speed += p.resolveStressCard()
		}

		speed += card.GetSpeed()
		p.AddIcons(card.GetIcons())
	}
	p.car.SetSpeed(speed)
}

// resolveStressCard handles the special Stress card effect
// Draws cards until a basic card is found, discarding non-basic cards
// Input: none
// Returns: the speed value of the basic card found
func (p *player) resolveStressCard() int {
	isBasic := false
	speed := 0

	for !isBasic {
		if p.deck.IsEmpty() {
			p.discardPile.ResetDeck(p.deck)
		}

		card := p.deck.DrawCard()
		if card.IsBasic() {
			isBasic = true
			speed = card.GetSpeed()
		}

		p.discardPile.AddCard(card)
	}

	return speed
}
