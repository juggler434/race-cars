package models

type Icon string

type Card interface {
	GetName() string
	GetSpeed() int
	GetIcons() map[Icon]int
	IsDiscardable() bool
	IsPlayable() bool
	IsBasic() bool
}

type card struct {
	name string
	speed int
	icons map[Icon]int
	discardable bool
	playable bool
	basic bool
}

// NewCard creates a new card instance
func NewCard(name string, speed int, icons map[Icon]int, discardable, playable, basic bool) Card {
	return &card{
		name:        name,
		speed:       speed,
		icons:       icons,
		discardable: discardable,
		playable:    playable,
		basic:       basic,
	}
}

// GetName returns the card's name
func (c *card) GetName() string {
	return c.name
}

// GetSpeed returns the card's speed value
func (c *card) GetSpeed() int {
	return c.speed
}

// GetIcons returns all icons on the card
func (c *card) GetIcons() map[Icon]int {
	return c.icons
}

// IsDiscardable returns whether the card can be discarded
func (c *card) IsDiscardable() bool {
	return c.discardable
}

// IsPlayable returns whether the card can be played
func (c *card) IsPlayable() bool {
	return c.playable
}

// IsBasic returns whether the card is a basic card
func (c *card) IsBasic() bool {
	return c.basic
}


