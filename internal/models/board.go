package models

// Board represents a racing board with spaces, finish line, and turn management
// The board manages the physical layout of the race track and the turn order of racers
type Board interface {
	// GetSpaces returns all spaces on the board
	// Returns: slice of all spaces that make up the race track
	GetSpaces() []Space

	// GetFinishLine returns the finish line space
	// Returns: the space that represents the finish line
	GetFinishLine() Space

	// GetRacerTurnOrder returns the current turn order of racers
	// Returns: slice of cars in turn order (first car is next to go)
	GetRacerTurnOrder() []Car

	// SetRacerTurnOrder calculates and sets the turn order based on current car positions
	// Cars are ordered by lap count (higher laps first), then by position on the board
	SetRacerTurnOrder()

	// GetNextRacer returns the next racer in turn order and removes them from the queue
	// Returns: the next car to take their turn
	GetNextRacer() Car
}

type board struct {
	spaces         []Space
	finishLine     Space
	racerTurnOrder []Car
	numberOfLaps   int
}

// NewBoard creates a new board instance
// Input: spaces - slice of spaces that make up the board
//
//	finishLine - the finish line space
//	numberOfLaps - the number of laps required to win the race
//
// Returns: a new Board
func NewBoard(spaces []Space, finishLine Space, numberOfLaps int) Board {
	return &board{
		spaces:         spaces,
		finishLine:     finishLine,
		racerTurnOrder: make([]Car, 0),
		numberOfLaps:   numberOfLaps,
	}
}

// GetSpaces returns all spaces on the board
// Input: none
// Returns: slice of all spaces on the board
func (b *board) GetSpaces() []Space {
	return b.spaces
}

// GetFinishLine returns the finish line space
// Input: none
// Returns: the finish line space
func (b *board) GetFinishLine() Space {
	return b.finishLine
}

// GetRacerTurnOrder returns the current turn order of racers
// Input: none
// Returns: slice of cars in turn order (first car is next to go)
func (b *board) GetRacerTurnOrder() []Car {
	return b.racerTurnOrder
}

// SetRacerTurnOrder calculates and sets the turn order based on current car positions
// Cars are ordered by lap count (higher laps first), then by position on the board
// Input: none
// Returns: none
func (b *board) SetRacerTurnOrder() {
	for _, space := range b.spaces {
		for _, car := range space.GetCars() {
			if len(b.racerTurnOrder) == 0 {
				b.racerTurnOrder = append(b.racerTurnOrder, car)
			} else if car.GetLap() > b.racerTurnOrder[0].GetLap() {
				b.racerTurnOrder = append([]Car{car}, b.racerTurnOrder...)
			} else {
				insertRacerInTurnOrder(b, car)
			}
		}
	}
}

// GetNextRacer returns the next racer in turn order and removes them from the queue
// Input: none
// Returns: the next car to take their turn
func (b *board) GetNextRacer() Car {
	nextRacer := b.racerTurnOrder[0]
	b.racerTurnOrder = b.racerTurnOrder[1:]
	return nextRacer
}

// insertRacerInTurnOrder inserts a car into the turn order based on lap count
// Cars with higher lap counts go first, cars with same lap count maintain relative order
// Input: b - the board instance
//
//	car - the car to insert
//
// Returns: none
func insertRacerInTurnOrder(b *board, car Car) {
	for i := len(b.racerTurnOrder) - 1; i >= 0; i-- {
		if car.GetLap() <= b.racerTurnOrder[i].GetLap() {
			if i == len(b.racerTurnOrder)-1 {
				b.racerTurnOrder = append(b.racerTurnOrder, car)
				break
			}
			b.racerTurnOrder = append(b.racerTurnOrder[:i+1], b.racerTurnOrder[i:]...)
			b.racerTurnOrder[i] = car
			break
		}
	}
}
