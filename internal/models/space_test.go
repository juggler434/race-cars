package models

import (
	"reflect"
	"testing"
)

func TestNewSpace(t *testing.T) {
	tests := []struct {
		name        string
		next        Space
		previous    Space
		corner      int
		finishLine  bool
		description string
	}{
		{
			name:        "Regular space",
			next:        nil,
			previous:    nil,
			corner:      1,
			finishLine:  false,
			description: "Create a regular space with corner 1",
		},
		{
			name:        "Finish line space",
			next:        nil,
			previous:    nil,
			corner:      0,
			finishLine:  true,
			description: "Create a finish line space",
		},
		{
			name:        "Space with negative corner",
			next:        nil,
			previous:    nil,
			corner:      -1,
			finishLine:  false,
			description: "Create a space with negative corner",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			space := NewSpace(tt.next, tt.previous, tt.corner, tt.finishLine)

			// Test that the space implements the Space interface
			var _ Space = space

			// Test initial values
			if len(space.GetCars()) != 0 {
				t.Errorf("NewSpace() cars = %v, want empty slice", space.GetCars())
			}
			if space.GetNext() != tt.next {
				t.Errorf("NewSpace() next = %v, want %v", space.GetNext(), tt.next)
			}
			if space.GetPrevious() != tt.previous {
				t.Errorf("NewSpace() previous = %v, want %v", space.GetPrevious(), tt.previous)
			}
			if space.GetCorner() != tt.corner {
				t.Errorf("NewSpace() corner = %d, want %d", space.GetCorner(), tt.corner)
			}
			if space.IsFinishLine() != tt.finishLine {
				t.Errorf("NewSpace() finishLine = %t, want %t", space.IsFinishLine(), tt.finishLine)
			}
			if space.IsFull() {
				t.Error("NewSpace() should not be full initially")
			}
			if space.IsOccupied() {
				t.Error("NewSpace() should not be occupied initially")
			}
		})
	}
}

func TestSpace_GetCars(t *testing.T) {
	space := NewSpace(nil, nil, 1, false)

	// Test initial empty cars
	cars := space.GetCars()
	if len(cars) != 0 {
		t.Errorf("Initial cars = %v, want empty slice", cars)
	}

	// Test after adding cars
	car1 := NewCar("red", 3)
	car2 := NewCar("blue", 2)
	space.AddCar(car1)
	space.AddCar(car2)

	cars = space.GetCars()
	expected := []Car{car1, car2}
	if !reflect.DeepEqual(cars, expected) {
		t.Errorf("Cars after adding = %v, want %v", cars, expected)
	}

	// Test that returned slice is a copy
	cars[0] = nil
	originalCars := space.GetCars()
	if originalCars[0] == nil {
		t.Error("GetCars() returned a reference, not a copy")
	}
}

func TestSpace_GetNext(t *testing.T) {
	nextSpace := NewSpace(nil, nil, 2, false)
	space := NewSpace(nextSpace, nil, 1, false)

	if space.GetNext() != nextSpace {
		t.Errorf("GetNext() = %v, want %v", space.GetNext(), nextSpace)
	}

	// Test with nil next
	space2 := NewSpace(nil, nil, 1, false)
	if space2.GetNext() != nil {
		t.Errorf("GetNext() = %v, want nil", space2.GetNext())
	}
}

func TestSpace_GetPrevious(t *testing.T) {
	prevSpace := NewSpace(nil, nil, 0, true)
	space := NewSpace(nil, prevSpace, 1, false)

	if space.GetPrevious() != prevSpace {
		t.Errorf("GetPrevious() = %v, want %v", space.GetPrevious(), prevSpace)
	}

	// Test with nil previous
	space2 := NewSpace(nil, nil, 1, false)
	if space2.GetPrevious() != nil {
		t.Errorf("GetPrevious() = %v, want nil", space2.GetPrevious())
	}
}

func TestSpace_IsFull(t *testing.T) {
	space := NewSpace(nil, nil, 1, false)

	// Test initial state
	if space.IsFull() {
		t.Error("New space should not be full")
	}

	// Test with one car
	car1 := NewCar("red", 3)
	space.AddCar(car1)
	if space.IsFull() {
		t.Error("Space with one car should not be full")
	}

	// Test with two cars
	car2 := NewCar("blue", 2)
	space.AddCar(car2)
	if !space.IsFull() {
		t.Error("Space with two cars should be full")
	}
}

func TestSpace_AddCar(t *testing.T) {
	tests := []struct {
		name        string
		carsToAdd   []Car
		expectError bool
		errorMsg    string
	}{
		{
			name:        "Add first car",
			carsToAdd:   []Car{NewCar("red", 3)},
			expectError: false,
		},
		{
			name:        "Add second car",
			carsToAdd:   []Car{NewCar("red", 3), NewCar("blue", 2)},
			expectError: false,
		},
		{
			name:        "Add third car should fail",
			carsToAdd:   []Car{NewCar("red", 3), NewCar("blue", 2), NewCar("green", 1)},
			expectError: true,
			errorMsg:    "space is full",
		},
		{
			name:        "Add nil car should fail",
			carsToAdd:   []Car{nil},
			expectError: true,
			errorMsg:    "cannot add nil car",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			space := NewSpace(nil, nil, 1, false)

			for i, car := range tt.carsToAdd {
				err := space.AddCar(car)

				if tt.expectError && i == len(tt.carsToAdd)-1 {
					// Only expect error on the last car for "space is full" case
					if err == nil {
						t.Errorf("Expected error but got none for car %d", i)
					} else if err.Error() != tt.errorMsg {
						t.Errorf("Expected error '%s', got '%s'", tt.errorMsg, err.Error())
					}
				} else if tt.expectError && tt.errorMsg == "cannot add nil car" {
					// For nil car case, expect error immediately
					if err == nil {
						t.Errorf("Expected error but got none for car %d", i)
					} else if err.Error() != tt.errorMsg {
						t.Errorf("Expected error '%s', got '%s'", tt.errorMsg, err.Error())
					}
					break
				} else {
					if err != nil {
						t.Errorf("AddCar() failed for car %d: %v", i, err)
					}
				}
			}

			// Verify final state
			if tt.expectError {
				// Should have no cars if we tried to add a nil car, or 2 cars if we tried to add a third
				expectedCount := 0
				if tt.errorMsg == "space is full" {
					expectedCount = 2
				}
				if len(space.GetCars()) != expectedCount {
					t.Errorf("Expected %d cars after failed add, got %d", expectedCount, len(space.GetCars()))
				}
			} else {
				if len(space.GetCars()) != len(tt.carsToAdd) {
					t.Errorf("Expected %d cars, got %d", len(tt.carsToAdd), len(space.GetCars()))
				}
			}
		})
	}
}

func TestSpace_RemoveCar(t *testing.T) {
	space := NewSpace(nil, nil, 1, false)
	car1 := NewCar("red", 3)
	car2 := NewCar("blue", 2)

	// Add cars
	space.AddCar(car1)
	space.AddCar(car2)

	tests := []struct {
		name        string
		carToRemove Car
		expectError bool
		errorMsg    string
	}{
		{
			name:        "Remove first car",
			carToRemove: car1,
			expectError: false,
		},
		{
			name:        "Remove second car",
			carToRemove: car2,
			expectError: false,
		},
		{
			name:        "Remove non-existent car",
			carToRemove: NewCar("green", 1),
			expectError: true,
			errorMsg:    "car not found",
		},
		{
			name:        "Remove nil car",
			carToRemove: nil,
			expectError: true,
			errorMsg:    "car not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a fresh space for each test
			testSpace := NewSpace(nil, nil, 1, false)
			testSpace.AddCar(car1)
			testSpace.AddCar(car2)

			err := testSpace.RemoveCar(tt.carToRemove)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				} else if err.Error() != tt.errorMsg {
					t.Errorf("Expected error '%s', got '%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				// Verify car was removed
				cars := testSpace.GetCars()
				for _, car := range cars {
					if car == tt.carToRemove {
						t.Error("Car was not removed")
					}
				}
			}
		})
	}
}

func TestSpace_IsOccupied(t *testing.T) {
	space := NewSpace(nil, nil, 1, false)

	// Test initial state
	if space.IsOccupied() {
		t.Error("New space should not be occupied")
	}

	// Test with one car
	car1 := NewCar("red", 3)
	space.AddCar(car1)
	if !space.IsOccupied() {
		t.Error("Space with one car should be occupied")
	}

	// Test with two cars
	car2 := NewCar("blue", 2)
	space.AddCar(car2)
	if !space.IsOccupied() {
		t.Error("Space with two cars should be occupied")
	}

	// Test after removing all cars
	space.RemoveCar(car1)
	space.RemoveCar(car2)
	if space.IsOccupied() {
		t.Error("Space with no cars should not be occupied")
	}
}

func TestSpace_GetCorner(t *testing.T) {
	tests := []struct {
		name     string
		corner   int
		expected int
	}{
		{
			name:     "Corner 1",
			corner:   1,
			expected: 1,
		},
		{
			name:     "Corner 0",
			corner:   0,
			expected: 0,
		},
		{
			name:     "Negative corner",
			corner:   -1,
			expected: -1,
		},
		{
			name:     "Large corner number",
			corner:   100,
			expected: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			space := NewSpace(nil, nil, tt.corner, false)
			if space.GetCorner() != tt.expected {
				t.Errorf("GetCorner() = %d, want %d", space.GetCorner(), tt.expected)
			}
		})
	}
}

func TestSpace_IsFinishLine(t *testing.T) {
	tests := []struct {
		name       string
		finishLine bool
		expected   bool
	}{
		{
			name:       "Regular space",
			finishLine: false,
			expected:   false,
		},
		{
			name:       "Finish line space",
			finishLine: true,
			expected:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			space := NewSpace(nil, nil, 1, tt.finishLine)
			if space.IsFinishLine() != tt.expected {
				t.Errorf("IsFinishLine() = %t, want %t", space.IsFinishLine(), tt.expected)
			}
		})
	}
}

func TestSpace_InterfaceCompliance(t *testing.T) {
	space := NewSpace(nil, nil, 1, false)

	// Test that space implements the Space interface
	var _ Space = space

	// Test that we can call all interface methods
	car := NewCar("red", 3)

	// These should not panic
	_ = space.GetCars()
	_ = space.GetNext()
	_ = space.GetPrevious()
	_ = space.IsFull()
	_ = space.AddCar(car)
	_ = space.RemoveCar(car)
	_ = space.IsOccupied()
	_ = space.GetCorner()
	_ = space.IsFinishLine()
}

func TestSpace_EdgeCases(t *testing.T) {
	t.Run("AddCar with nil car", func(t *testing.T) {
		space := NewSpace(nil, nil, 1, false)

		// This should return an error for nil cars
		err := space.AddCar(nil)
		if err == nil {
			t.Error("Expected error when adding nil car")
		} else if err.Error() != "cannot add nil car" {
			t.Errorf("Expected 'cannot add nil car' error, got: %s", err.Error())
		}

		if len(space.GetCars()) != 0 {
			t.Error("AddCar with nil car should not add anything")
		}
	})

	t.Run("RemoveCar from empty space", func(t *testing.T) {
		space := NewSpace(nil, nil, 1, false)
		car := NewCar("red", 3)

		err := space.RemoveCar(car)
		if err == nil {
			t.Error("Expected error when removing car from empty space")
		} else if err.Error() != "car not found" {
			t.Errorf("Expected 'car not found' error, got: %s", err.Error())
		}
	})

	t.Run("AddCar to full space", func(t *testing.T) {
		space := NewSpace(nil, nil, 1, false)
		car1 := NewCar("red", 3)
		car2 := NewCar("blue", 2)
		car3 := NewCar("green", 1)

		// Add two cars
		space.AddCar(car1)
		space.AddCar(car2)

		// Try to add third car
		err := space.AddCar(car3)
		if err == nil {
			t.Error("Expected error when adding car to full space")
		} else if err.Error() != "space is full" {
			t.Errorf("Expected 'space is full' error, got: %s", err.Error())
		}

		// Verify only two cars remain
		if len(space.GetCars()) != 2 {
			t.Errorf("Expected 2 cars, got %d", len(space.GetCars()))
		}
	})

	t.Run("RemoveCar and add again", func(t *testing.T) {
		space := NewSpace(nil, nil, 1, false)
		car1 := NewCar("red", 3)
		car2 := NewCar("blue", 2)

		// Add two cars
		space.AddCar(car1)
		space.AddCar(car2)

		// Remove one car
		err := space.RemoveCar(car1)
		if err != nil {
			t.Errorf("Failed to remove car: %v", err)
		}

		// Should be able to add another car now
		car3 := NewCar("green", 1)
		err = space.AddCar(car3)
		if err != nil {
			t.Errorf("Failed to add car after removal: %v", err)
		}

		// Should have 2 cars now
		if len(space.GetCars()) != 2 {
			t.Errorf("Expected 2 cars after removal and addition, got %d", len(space.GetCars()))
		}
	})

	t.Run("Cars immutability", func(t *testing.T) {
		space := NewSpace(nil, nil, 1, false)
		car := NewCar("red", 3)
		space.AddCar(car)

		// Get cars and modify the slice
		cars := space.GetCars()
		cars[0] = nil

		// Get cars again and verify they weren't modified
		originalCars := space.GetCars()
		if originalCars[0] == nil {
			t.Error("GetCars() returned a reference, not a copy")
		}
	})

	t.Run("Linked spaces", func(t *testing.T) {
		// Create spaces with proper bidirectional linking
		space1 := NewSpace(nil, nil, 1, false)
		space2 := NewSpace(nil, space1, 2, false)

		// Note: We can't easily test bidirectional linking with the current NewSpace function
		// since it doesn't allow updating the next/previous references after creation.
		// This test verifies that the references are set correctly during creation.

		if space2.GetPrevious() != space1 {
			t.Error("space2.GetPrevious() should return space1")
		}
		if space1.GetNext() != nil {
			t.Error("space1.GetNext() should return nil since it wasn't set during creation")
		}
	})
}

// Benchmark tests
func BenchmarkNewSpace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewSpace(nil, nil, 1, false)
	}
}

func BenchmarkSpace_AddCar(b *testing.B) {
	car := NewCar("red", 3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Reset space for each iteration
		space := NewSpace(nil, nil, 1, false)
		space.AddCar(car)
	}
}

func BenchmarkSpace_RemoveCar(b *testing.B) {
	space := NewSpace(nil, nil, 1, false)
	car := NewCar("red", 3)
	space.AddCar(car)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Re-add car for each iteration
		space.AddCar(car)
		space.RemoveCar(car)
	}
}

func BenchmarkSpace_IsFull(b *testing.B) {
	space := NewSpace(nil, nil, 1, false)
	car1 := NewCar("red", 3)
	car2 := NewCar("blue", 2)
	space.AddCar(car1)
	space.AddCar(car2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		space.IsFull()
	}
}
