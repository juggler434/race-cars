package models

// Icon represents the type of icon on a card
// It is represented as an integer, but can be converted to a string

type Icon int

const (
	IconBoost Icon = iota
	IconCooling
)

var iconName = map[Icon]string{
	IconBoost: "Boost",
	IconCooling: "Cooling",
}

func (i Icon) String() string {
	return iconName[i]
}