package models

import (
	"fmt"
	"math"
)

type Color string

const (
	Red    Color = "Red"
	Blue   Color = "Blue"
	Green  Color = "Green"
	Yellow Color = "Yellow"
	Orange Color = "Orange"
	Black  Color = "Black"
	Gray   Color = "Gray"
)

type Car interface {
	GetColor() string
	GetSpeed() int
	SetSpeed(int)
	GetPassedCorners() []int
	AddPassedCorner(int)
	ResetPassedCorners()
	GetLap() int
	IncreaseLap()
	GetGear() int
	SetGear(int, DiscardPile) (map[Icon]int, error)
	GetEngine() int
}

type car struct {
	color         string
	speed         int
	passedCorners []int
	lap           int
	gear          int
	engine        int
}

// NewCar creates a new car instance
func NewCar(color string, engine int) Car {
	return &car{
		color:         color,
		speed:         0,
		passedCorners: make([]int, 0),
		lap:           0,
		gear:          1,
		engine:        engine,
	}
}

// GetColor returns the car's color
func (c *car) GetColor() string {
	return c.color
}

// GetSpeed returns the car's current speed
func (c *car) GetSpeed() int {
	return c.speed
}

// SetSpeed sets the car's speed
func (c *car) SetSpeed(speed int) {
	c.speed = speed
}

// GetPassedCorners returns the list of corners the car has passed
func (c *car) GetPassedCorners() []int {
	// Return a copy to prevent external modification
	result := make([]int, len(c.passedCorners))
	copy(result, c.passedCorners)
	return result
}

// AddPassedCorner adds a corner to the list of passed corners
func (c *car) AddPassedCorner(corner int) {
	c.passedCorners = append(c.passedCorners, corner)
}

// ResetPassedCorners resets the list of passed corners
func (c *car) ResetPassedCorners() {
	c.passedCorners = make([]int, 0)
}

// GetLap returns the current lap number
func (c *car) GetLap() int {
	return c.lap
}

// IncreaseLap increments the lap counter
func (c *car) IncreaseLap() {
	c.lap++
}

// GetGear returns the current gear
func (c *car) GetGear() int {
	return c.gear
}

// SetGear sets the car's gear and discards the old gear to the discard pile
func (c *car) SetGear(gear int, discardPile DiscardPile) (map[Icon]int, error) {

	if gear < 1 || gear > 5 {
		return nil, fmt.Errorf("gear must be between 1 and 5")
	}

	err := c.calculateGearShift(gear, discardPile)
	if err != nil {
		return nil, err
	}

	icons := c.getCoolingIconsForGear(gear)

	return icons, nil
}

// GetEngine returns the engine value
func (c *car) GetEngine() int {
	return c.engine
}

func (c *car) calculateGearShift(gear int, discardPile DiscardPile) error {
	noOfShifts := math.Abs(float64(gear - c.gear))

	switch noOfShifts {
	case 0:
		return nil
	case 1:
		c.gear = gear
	case 2:
		if c.engine == 0 {
			return fmt.Errorf("cannot shift up to gear %d with engine 0", gear)
		}
		c.gear = gear
		c.engine--
		if discardPile != nil {
			discardPile.AddCard(NewHeatCard())
		}
	default:
		return fmt.Errorf("cannot shift more than 2 gears at once")
	}
	return nil
}

func (c *car) getCoolingIconsForGear(gear int) map[Icon]int {
	icons := make(map[Icon]int)

	switch gear {
	case 1:
		icons[IconCooling] = 3
	case 2:
		icons[IconCooling] = 1
	}

	return icons
}
