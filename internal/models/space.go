package models

import "fmt"

type Space interface {
	GetCars() []Car
	GetNext() Space
	GetPrevious() Space
	IsFull() bool
	AddCar(Car) error
	RemoveCar(Car) error
	IsOccupied() bool
	GetCorner() int
	IsFinishLine() bool
}

type space struct {
	cars       []Car
	next       Space
	previous   Space
	corner     int
	finishLine bool
}

func NewSpace(next Space, previous Space, corner int, finishLine bool) Space {
	return &space{
		cars:       make([]Car, 0),
		next:       next,
		previous:   previous,
		corner:     corner,
		finishLine: finishLine,
	}
}

func (s *space) GetCars() []Car {
	// Return a copy to prevent external modification
	result := make([]Car, len(s.cars))
	copy(result, s.cars)
	return result
}

func (s *space) GetNext() Space {
	return s.next
}

func (s *space) GetPrevious() Space {
	return s.previous
}

func (s *space) IsFull() bool {
	return len(s.cars) == 2
}

func (s *space) AddCar(car Car) error {
	if car == nil {
		return fmt.Errorf("cannot add nil car")
	}
	if s.IsFull() {
		return fmt.Errorf("space is full")
	}
	s.cars = append(s.cars, car)
	return nil
}

func (s *space) RemoveCar(car Car) error {
	for i, c := range s.cars {
		if c == car {
			s.cars = append(s.cars[:i], s.cars[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("car not found")
}

func (s *space) IsOccupied() bool {
	return len(s.cars) > 0
}

func (s *space) GetCorner() int {
	return s.corner
}

func (s *space) IsFinishLine() bool {
	return s.finishLine
}
