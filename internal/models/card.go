package models

const (
	Stress string = "Stress"
	Heat   string = "Heat"
)

type Card interface {
	GetName() string
	GetSpeed() int
	GetIcons() map[Icon]int
	IsDiscardable() bool
	IsPlayable() bool
	IsBasic() bool
}

type card struct {
	name        string
	speed       int
	icons       map[Icon]int
	discardable bool
	playable    bool
	basic       bool
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

// NewHeatCard creates a new heat card instance
func NewHeatCard() Card {
	return NewCard(Heat, 0, nil, false, false, false)
}

// NewStressCard creates a new stress card instance
func NewStressCard() Card {
	return NewCard(Stress, 0, nil, false, true, false)
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
	icons := make(map[Icon]int)
	for icon, count := range c.icons {
		icons[icon] = count
	}
	return icons
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
