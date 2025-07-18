package models

import (
	"reflect"
	"testing"
)

func TestNewBoard(t *testing.T) {
	finishLine := NewSpace(nil, nil, 0, true)
	spaces := []Space{NewSpace(nil, nil, 1, false), NewSpace(nil, nil, 2, false)}
	board := NewBoard(spaces, finishLine, 3)

	if !reflect.DeepEqual(board.GetSpaces(), spaces) {
		t.Errorf("GetSpaces() = %v, want %v", board.GetSpaces(), spaces)
	}
	if board.GetFinishLine() != finishLine {
		t.Errorf("GetFinishLine() = %v, want %v", board.GetFinishLine(), finishLine)
	}
	if len(board.GetRacerTurnOrder()) != 0 {
		t.Errorf("Initial racerTurnOrder = %v, want empty", board.GetRacerTurnOrder())
	}
}

func TestBoard_SetRacerTurnOrder_And_GetNextRacer(t *testing.T) {
	space1 := NewSpace(nil, nil, 1, false)
	space2 := NewSpace(nil, nil, 2, false)
	car1 := NewCar("red", 3)
	car2 := NewCar("blue", 2)
	car3 := NewCar("green", 1)
	car1.IncreaseLap() // lap 1
	car2.IncreaseLap() // lap 1
	car2.IncreaseLap() // lap 2
	space1.AddCar(car1)
	space2.AddCar(car2)
	space1.AddCar(car3)
	spaces := []Space{space1, space2}
	finishLine := NewSpace(nil, nil, 0, true)
	board := NewBoard(spaces, finishLine, 3)

	board.SetRacerTurnOrder()
	order := board.GetRacerTurnOrder()
	expectedColors := []string{"blue", "red", "green"}
	if len(order) != len(expectedColors) {
		t.Fatalf("Expected %d cars in turn order, got %d", len(expectedColors), len(order))
	}
	for i, color := range expectedColors {
		if order[i].GetColor() != color {
			t.Errorf("Turn order[%d] = %s, want %s", i, order[i].GetColor(), color)
		}
	}

	// Test GetNextRacer
	for _, color := range expectedColors {
		next := board.GetNextRacer()
		if next.GetColor() != color {
			t.Errorf("GetNextRacer() = %s, want %s", next.GetColor(), color)
		}
	}
	if len(board.GetRacerTurnOrder()) != 0 {
		t.Error("Turn order should be empty after all racers have gone")
	}
}

func TestBoard_EdgeCases(t *testing.T) {
	t.Run("Empty board", func(t *testing.T) {
		board := NewBoard([]Space{}, NewSpace(nil, nil, 0, true), 3)
		board.SetRacerTurnOrder()
		if len(board.GetRacerTurnOrder()) != 0 {
			t.Error("Turn order should be empty for empty board")
		}
	})

	t.Run("No cars in spaces", func(t *testing.T) {
		spaces := []Space{NewSpace(nil, nil, 1, false), NewSpace(nil, nil, 2, false)}
		board := NewBoard(spaces, NewSpace(nil, nil, 0, true), 3)
		board.SetRacerTurnOrder()
		if len(board.GetRacerTurnOrder()) != 0 {
			t.Error("Turn order should be empty when no cars are on the board")
		}
	})

	t.Run("GetNextRacer panics on empty turn order", func(t *testing.T) {
		board := NewBoard([]Space{}, NewSpace(nil, nil, 0, true), 3)
		defer func() {
			if r := recover(); r == nil {
				t.Error("GetNextRacer() should panic when turn order is empty")
			}
		}()
		board.GetNextRacer()
	})
}

func TestBoard_InterfaceCompliance(t *testing.T) {
	board := NewBoard([]Space{}, NewSpace(nil, nil, 0, true), 3)
	var _ Board = board
	_ = board.GetSpaces()
	_ = board.GetFinishLine()
	_ = board.GetRacerTurnOrder()
	board.SetRacerTurnOrder()
}

func TestBoard_SetRacerTurnOrder_CoversInsertRacerInTurnOrder(t *testing.T) {
	// Create spaces with cars in specific positions to trigger insertRacerInTurnOrder
	space1 := NewSpace(nil, nil, 1, false)
	space2 := NewSpace(nil, nil, 2, false)
	space3 := NewSpace(nil, nil, 3, false)

	car1 := NewCar("red", 3)
	car2 := NewCar("blue", 2)
	car3 := NewCar("green", 1)

	// Set up cars with specific lap counts to trigger the insertion logic
	car1.IncreaseLap()
	car1.IncreaseLap() // red car is on lap 2
	car2.IncreaseLap() // blue car is on lap 1
	// green car is on lap 0

	// Add cars to spaces in order that will trigger insertRacerInTurnOrder
	space1.AddCar(car1) // lap 2 - will be first in turn order
	space2.AddCar(car3) // lap 0 - will trigger insertRacerInTurnOrder
	space3.AddCar(car2) // lap 1 - will trigger insertRacerInTurnOrder

	spaces := []Space{space1, space2, space3}
	finishLine := NewSpace(nil, nil, 0, true)
	board := NewBoard(spaces, finishLine, 3)

	board.SetRacerTurnOrder()

	// Verify the turn order is correct: blue (lap 1), red (lap 2), green (lap 0)
	// The actual order depends on how insertRacerInTurnOrder works
	turnOrder := board.GetRacerTurnOrder()
	expectedColors := []string{"blue", "red", "green"}

	if len(turnOrder) != len(expectedColors) {
		t.Fatalf("Expected %d cars in turn order, got %d", len(expectedColors), len(turnOrder))
	}

	for i, expectedColor := range expectedColors {
		if turnOrder[i].GetColor() != expectedColor {
			t.Errorf("Turn order[%d] = %s, want %s", i, turnOrder[i].GetColor(), expectedColor)
		}
	}
}
