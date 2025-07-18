package models

import (
	"fmt"
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
	space4 := NewSpace(nil, nil, 4, false)

	car1 := NewCar("red", 3)
	car2 := NewCar("blue", 2)
	car3 := NewCar("green", 1)
	car4 := NewCar("yellow", 1) // Same lap as green car
	car5 := NewCar("purple", 1) // Same space as yellow car

	// Set up cars with specific lap counts to trigger the insertion logic
	car1.IncreaseLap()
	car1.IncreaseLap() // red car is on lap 2
	car2.IncreaseLap() // blue car is on lap 1
	// green car is on lap 0
	// yellow car is on lap 0 (same as green)
	// purple car is on lap 0 (same as green and yellow)

	// Add cars to spaces in order that will trigger insertRacerInTurnOrder
	space1.AddCar(car1) // lap 2 - will be first in turn order
	space2.AddCar(car3) // lap 0 - will trigger insertRacerInTurnOrder
	space3.AddCar(car2) // lap 1 - will trigger insertRacerInTurnOrder
	space4.AddCar(car4) // lap 0 - will trigger insertRacerInTurnOrder
	space4.AddCar(car5) // lap 0 - same space as yellow car

	spaces := []Space{space1, space2, space3, space4}
	finishLine := NewSpace(nil, nil, 0, true)
	board := NewBoard(spaces, finishLine, 3)

	board.SetRacerTurnOrder()

	// Verify the turn order is correct: blue (lap 1), red (lap 2), green (lap 0), yellow (lap 0), purple (lap 0)
	// The actual order depends on how insertRacerInTurnOrder works
	turnOrder := board.GetRacerTurnOrder()
	expectedColors := []string{"blue", "red", "green", "yellow", "purple"}

	if len(turnOrder) != len(expectedColors) {
		t.Fatalf("Expected %d cars in turn order, got %d", len(expectedColors), len(turnOrder))
	}

	for i, expectedColor := range expectedColors {
		if turnOrder[i].GetColor() != expectedColor {
			t.Errorf("Turn order[%d] = %s, want %s", i, turnOrder[i].GetColor(), expectedColor)
		}
	}
}

func BenchmarkNewBoard(b *testing.B) {
	finishLine := NewSpace(nil, nil, 0, true)
	spaces := make([]Space, 10) // Reasonable number of spaces for a race track
	for i := 0; i < 10; i++ {
		spaces[i] = NewSpace(nil, nil, i+1, false)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewBoard(spaces, finishLine, 3)
	}
}

func BenchmarkSetRacerTurnOrder_SmallBoard(b *testing.B) {
	// Small board with few cars (2-3 cars)
	space1 := NewSpace(nil, nil, 1, false)
	space2 := NewSpace(nil, nil, 2, false)
	car1 := NewCar("red", 3)
	car2 := NewCar("blue", 2)
	space1.AddCar(car1)
	space2.AddCar(car2)
	spaces := []Space{space1, space2}
	finishLine := NewSpace(nil, nil, 0, true)
	board := NewBoard(spaces, finishLine, 3)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.SetRacerTurnOrder()
	}
}

func BenchmarkSetRacerTurnOrder_MediumBoard(b *testing.B) {
	// Medium board with 6-8 cars (typical game scenario)
	spaces := make([]Space, 6)
	for i := 0; i < 6; i++ {
		spaces[i] = NewSpace(nil, nil, i+1, false)
		// Add 1-2 cars per space (max 2 per space constraint)
		for j := 0; j < 1+(i%2); j++ {
			car := NewCar(fmt.Sprintf("car%d_%d", i, j), 3)
			car.IncreaseLap() // Add some lap variation
			if i%3 == 0 {
				car.IncreaseLap()
			}
			spaces[i].AddCar(car)
		}
	}
	finishLine := NewSpace(nil, nil, 0, true)
	board := NewBoard(spaces, finishLine, 3)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.SetRacerTurnOrder()
	}
}

func BenchmarkSetRacerTurnOrder_MaxCars(b *testing.B) {
	// Board with maximum 8 cars (4 spaces with 2 cars each)
	spaces := make([]Space, 4)
	for i := 0; i < 4; i++ {
		spaces[i] = NewSpace(nil, nil, i+1, false)
		// Add exactly 2 cars per space (max constraint)
		for j := 0; j < 2; j++ {
			car := NewCar(fmt.Sprintf("car%d_%d", i, j), 3)
			// Add lap variation
			for k := 0; k < i%4; k++ {
				car.IncreaseLap()
			}
			spaces[i].AddCar(car)
		}
	}
	finishLine := NewSpace(nil, nil, 0, true)
	board := NewBoard(spaces, finishLine, 3)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.SetRacerTurnOrder()
	}
}

func BenchmarkGetNextRacer(b *testing.B) {
	// Setup board with cars
	space1 := NewSpace(nil, nil, 1, false)
	space2 := NewSpace(nil, nil, 2, false)
	car1 := NewCar("red", 3)
	car2 := NewCar("blue", 2)
	car3 := NewCar("green", 1)
	space1.AddCar(car1)
	space2.AddCar(car2)
	space1.AddCar(car3)
	spaces := []Space{space1, space2}
	finishLine := NewSpace(nil, nil, 0, true)
	board := NewBoard(spaces, finishLine, 3)
	board.SetRacerTurnOrder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Reset the board for each iteration to avoid running out of racers
		if len(board.GetRacerTurnOrder()) == 0 {
			board.SetRacerTurnOrder()
		}
		board.GetNextRacer()
	}
}

func BenchmarkBoardOperations_Complete(b *testing.B) {
	// Benchmark complete board operations cycle with realistic game constraints
	spaces := make([]Space, 4)
	for i := 0; i < 4; i++ {
		spaces[i] = NewSpace(nil, nil, i+1, false)
		// Add 1-2 cars per space (max 2 per space)
		for j := 0; j < 1+(i%2); j++ {
			car := NewCar(fmt.Sprintf("car%d_%d", i, j), 3)
			// Add some lap variation
			for k := 0; k < i%3; k++ {
				car.IncreaseLap()
			}
			spaces[i].AddCar(car)
		}
	}
	finishLine := NewSpace(nil, nil, 0, true)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board := NewBoard(spaces, finishLine, 3)
		board.SetRacerTurnOrder()

		// Process all racers (max 8 cars)
		for len(board.GetRacerTurnOrder()) > 0 {
			board.GetNextRacer()
		}
	}
}
