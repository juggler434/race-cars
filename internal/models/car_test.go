package models

import (
	"reflect"
	"testing"
)

func TestNewCar(t *testing.T) {
	tests := []struct {
		name     string
		color    string
		engine   int
		expected Car
	}{
		{
			name:   "Create red car with engine 3",
			color:  "red",
			engine: 3,
		},
		{
			name:   "Create blue car with engine 0",
			color:  "blue",
			engine: 0,
		},
		{
			name:   "Create car with empty color",
			color:  "",
			engine: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			car := NewCar(tt.color, tt.engine)

			// Test that the car implements the Car interface
			var _ Car = car

			// Test initial values
			if car.GetColor() != tt.color {
				t.Errorf("NewCar() color = %s, want %s", car.GetColor(), tt.color)
			}
			if car.GetEngine() != tt.engine {
				t.Errorf("NewCar() engine = %d, want %d", car.GetEngine(), tt.engine)
			}
			if car.GetSpeed() != 0 {
				t.Errorf("NewCar() speed = %d, want 0", car.GetSpeed())
			}
			if car.GetLap() != 0 {
				t.Errorf("NewCar() lap = %d, want 0", car.GetLap())
			}
			if car.GetGear() != 1 {
				t.Errorf("NewCar() gear = %d, want 1", car.GetGear())
			}
			if len(car.GetPassedCorners()) != 0 {
				t.Errorf("NewCar() passed corners = %v, want empty slice", car.GetPassedCorners())
			}
		})
	}
}

func TestCar_GetColor(t *testing.T) {
	tests := []struct {
		name     string
		color    string
		expected string
	}{
		{
			name:     "Red car",
			color:    "red",
			expected: "red",
		},
		{
			name:     "Blue car",
			color:    "blue",
			expected: "blue",
		},
		{
			name:     "Empty color",
			color:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			car := NewCar(tt.color, 3)
			if car.GetColor() != tt.expected {
				t.Errorf("GetColor() = %s, want %s", car.GetColor(), tt.expected)
			}
		})
	}
}

func TestCar_GetSpeed(t *testing.T) {
	car := NewCar("red", 3)

	// Test initial speed
	if car.GetSpeed() != 0 {
		t.Errorf("Initial speed = %d, want 0", car.GetSpeed())
	}

	// Test after setting speed
	car.SetSpeed(50)
	if car.GetSpeed() != 50 {
		t.Errorf("Speed after SetSpeed(50) = %d, want 50", car.GetSpeed())
	}

	// Test negative speed
	car.SetSpeed(-10)
	if car.GetSpeed() != -10 {
		t.Errorf("Speed after SetSpeed(-10) = %d, want -10", car.GetSpeed())
	}
}

func TestCar_SetSpeed(t *testing.T) {
	tests := []struct {
		name     string
		speed    int
		expected int
	}{
		{
			name:     "Set positive speed",
			speed:    100,
			expected: 100,
		},
		{
			name:     "Set zero speed",
			speed:    0,
			expected: 0,
		},
		{
			name:     "Set negative speed",
			speed:    -20,
			expected: -20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			car := NewCar("red", 3)
			car.SetSpeed(tt.speed)
			if car.GetSpeed() != tt.expected {
				t.Errorf("SetSpeed(%d) resulted in speed %d, want %d", tt.speed, car.GetSpeed(), tt.expected)
			}
		})
	}
}

func TestCar_GetPassedCorners(t *testing.T) {
	car := NewCar("red", 3)

	// Test initial empty corners
	corners := car.GetPassedCorners()
	if len(corners) != 0 {
		t.Errorf("Initial passed corners = %v, want empty slice", corners)
	}

	// Test after adding corners
	car.AddPassedCorner(1)
	car.AddPassedCorner(2)
	corners = car.GetPassedCorners()
	expected := []int{1, 2}
	if !reflect.DeepEqual(corners, expected) {
		t.Errorf("Passed corners after adding = %v, want %v", corners, expected)
	}

	// Test that returned slice is a copy
	corners[0] = 999
	originalCorners := car.GetPassedCorners()
	if originalCorners[0] == 999 {
		t.Error("GetPassedCorners() returned a reference, not a copy")
	}
}

func TestCar_AddPassedCorner(t *testing.T) {
	car := NewCar("red", 3)

	// Test adding single corner
	car.AddPassedCorner(1)
	corners := car.GetPassedCorners()
	if len(corners) != 1 || corners[0] != 1 {
		t.Errorf("After AddPassedCorner(1), corners = %v, want [1]", corners)
	}

	// Test adding multiple corners
	car.AddPassedCorner(2)
	car.AddPassedCorner(3)
	corners = car.GetPassedCorners()
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(corners, expected) {
		t.Errorf("After adding multiple corners = %v, want %v", corners, expected)
	}

	// Test adding negative corner
	car.AddPassedCorner(-1)
	corners = car.GetPassedCorners()
	if corners[len(corners)-1] != -1 {
		t.Errorf("After AddPassedCorner(-1), last corner = %d, want -1", corners[len(corners)-1])
	}
}

func TestCar_ResetPassedCorners(t *testing.T) {
	car := NewCar("red", 3)

	// Add some corners
	car.AddPassedCorner(1)
	car.AddPassedCorner(2)
	car.AddPassedCorner(3)

	// Verify corners were added
	if len(car.GetPassedCorners()) != 3 {
		t.Error("Corners were not added properly")
	}

	// Reset corners
	car.ResetPassedCorners()

	// Verify corners were reset
	if len(car.GetPassedCorners()) != 0 {
		t.Errorf("After ResetPassedCorners(), corners = %v, want empty slice", car.GetPassedCorners())
	}
}

func TestCar_GetLap(t *testing.T) {
	car := NewCar("red", 3)

	// Test initial lap
	if car.GetLap() != 0 {
		t.Errorf("Initial lap = %d, want 0", car.GetLap())
	}

	// Test after increasing lap
	car.IncreaseLap()
	if car.GetLap() != 1 {
		t.Errorf("After IncreaseLap(), lap = %d, want 1", car.GetLap())
	}

	// Test multiple increases
	car.IncreaseLap()
	car.IncreaseLap()
	if car.GetLap() != 3 {
		t.Errorf("After multiple IncreaseLap(), lap = %d, want 3", car.GetLap())
	}
}

func TestCar_IncreaseLap(t *testing.T) {
	car := NewCar("red", 3)

	// Test initial state
	if car.GetLap() != 0 {
		t.Errorf("Initial lap = %d, want 0", car.GetLap())
	}

	// Test single increase
	car.IncreaseLap()
	if car.GetLap() != 1 {
		t.Errorf("After IncreaseLap(), lap = %d, want 1", car.GetLap())
	}

	// Test multiple increases
	for i := 0; i < 5; i++ {
		car.IncreaseLap()
	}
	if car.GetLap() != 6 {
		t.Errorf("After 6 increases, lap = %d, want 6", car.GetLap())
	}
}

func TestCar_GetGear(t *testing.T) {
	car := NewCar("red", 3)

	// Test initial gear
	if car.GetGear() != 1 {
		t.Errorf("Initial gear = %d, want 1", car.GetGear())
	}

	// Test after changing gear
	discardPile := NewDiscardPile()
	_, err := car.SetGear(3, discardPile)
	if err != nil {
		t.Errorf("SetGear(3) failed: %v", err)
	}
	if car.GetGear() != 3 {
		t.Errorf("After SetGear(3), gear = %d, want 3", car.GetGear())
	}
}

func TestCar_SetGear(t *testing.T) {
	tests := []struct {
		name        string
		initialGear int
		newGear     int
		engine      int
		expectError bool
		errorMsg    string
		expectIcons map[Icon]int
	}{
		{
			name:        "Shift to same gear",
			initialGear: 1,
			newGear:     1,
			engine:      3,
			expectError: false,
			expectIcons: map[Icon]int{IconCooling: 3},
		},
		{
			name:        "Shift up one gear",
			initialGear: 1,
			newGear:     2,
			engine:      3,
			expectError: false,
			expectIcons: map[Icon]int{IconCooling: 1},
		},
		{
			name:        "Shift up two gears",
			initialGear: 1,
			newGear:     3,
			engine:      3,
			expectError: false,
			expectIcons: map[Icon]int{},
		},
		{
			name:        "Shift up two gears with engine 0",
			initialGear: 1,
			newGear:     3,
			engine:      0,
			expectError: true,
			errorMsg:    "cannot shift up to gear 3 with engine 0",
		},
		{
			name:        "Shift up three gears",
			initialGear: 1,
			newGear:     4,
			engine:      3,
			expectError: true,
			errorMsg:    "cannot shift more than 2 gears at once",
		},
		{
			name:        "Shift to gear 0",
			initialGear: 1,
			newGear:     0,
			engine:      3,
			expectError: true,
			errorMsg:    "gear must be between 1 and 5",
		},
		{
			name:        "Shift to gear 6",
			initialGear: 1,
			newGear:     6,
			engine:      3,
			expectError: true,
			errorMsg:    "gear must be between 1 and 5",
		},
		{
			name:        "Shift to gear 1 with cooling",
			initialGear: 3,
			newGear:     1,
			engine:      3,
			expectError: false,
			expectIcons: map[Icon]int{IconCooling: 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			car := NewCar("red", tt.engine)

			// Set initial gear
			discardPile := NewDiscardPile()
			_, err := car.SetGear(tt.initialGear, discardPile)
			if err != nil {
				t.Fatalf("Failed to set initial gear: %v", err)
			}

			// Test gear shift
			icons, err := car.SetGear(tt.newGear, discardPile)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else if err.Error() != tt.errorMsg {
					t.Errorf("Expected error '%s', got '%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if !reflect.DeepEqual(icons, tt.expectIcons) {
					t.Errorf("Expected icons %v, got %v", tt.expectIcons, icons)
				}
			}
		})
	}
}

func TestCar_SetGear_EngineReduction(t *testing.T) {
	car := NewCar("red", 3)
	discardPile := NewDiscardPile()

	// Shift up two gears (should reduce engine by 1)
	_, err := car.SetGear(3, discardPile)
	if err != nil {
		t.Fatalf("SetGear(3) failed: %v", err)
	}

	if car.GetEngine() != 2 {
		t.Errorf("After shifting up 2 gears, engine = %d, want 2", car.GetEngine())
	}

	// Check that heat card was added
	deck := NewDeck([]Card{})
	discardPile.ResetDeck(deck)
	card := deck.DrawCard()
	if card == nil {
		t.Error("Expected heat card to be added to discard pile")
	}
}

func TestCar_GetEngine(t *testing.T) {
	tests := []struct {
		name     string
		engine   int
		expected int
	}{
		{
			name:     "Engine 3",
			engine:   3,
			expected: 3,
		},
		{
			name:     "Engine 0",
			engine:   0,
			expected: 0,
		},
		{
			name:     "Engine 5",
			engine:   5,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			car := NewCar("red", tt.engine)
			if car.GetEngine() != tt.expected {
				t.Errorf("GetEngine() = %d, want %d", car.GetEngine(), tt.expected)
			}
		})
	}
}

func TestCar_InterfaceCompliance(t *testing.T) {
	car := NewCar("red", 3)

	// Test that car implements the Car interface
	var _ Car = car

	// Test that we can call all interface methods
	discardPile := NewDiscardPile()

	// These should not panic
	_ = car.GetColor()
	_ = car.GetSpeed()
	car.SetSpeed(50)
	_ = car.GetPassedCorners()
	car.AddPassedCorner(1)
	car.ResetPassedCorners()
	_ = car.GetLap()
	car.IncreaseLap()
	_ = car.GetGear()
	_, _ = car.SetGear(2, discardPile)
	_ = car.GetEngine()
}

func TestCar_EdgeCases(t *testing.T) {
	t.Run("SetGear with nil discard pile", func(t *testing.T) {
		car := NewCar("red", 3)

		// This should not panic
		_, err := car.SetGear(2, nil)
		if err != nil {
			t.Errorf("SetGear with nil discard pile failed: %v", err)
		}
	})

	t.Run("Multiple gear shifts", func(t *testing.T) {
		car := NewCar("red", 5)
		discardPile := NewDiscardPile()

		// Shift up multiple times
		_, err := car.SetGear(2, discardPile)
		if err != nil {
			t.Errorf("First gear shift failed: %v", err)
		}

		_, err = car.SetGear(4, discardPile)
		if err != nil {
			t.Errorf("Second gear shift failed: %v", err)
		}

		if car.GetGear() != 4 {
			t.Errorf("Final gear = %d, want 4", car.GetGear())
		}

		if car.GetEngine() != 4 {
			t.Errorf("Final engine = %d, want 4", car.GetEngine())
		}
	})

	t.Run("Passed corners immutability", func(t *testing.T) {
		car := NewCar("red", 3)
		car.AddPassedCorner(1)
		car.AddPassedCorner(2)

		// Get corners and modify the slice
		corners := car.GetPassedCorners()
		corners[0] = 999

		// Get corners again and verify they weren't modified
		originalCorners := car.GetPassedCorners()
		if originalCorners[0] == 999 {
			t.Error("GetPassedCorners() returned a reference, not a copy")
		}
	})
}

// Benchmark tests
func BenchmarkNewCar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewCar("red", 3)
	}
}

func BenchmarkCar_SetSpeed(b *testing.B) {
	car := NewCar("red", 3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		car.SetSpeed(i % 100)
	}
}

func BenchmarkCar_AddPassedCorner(b *testing.B) {
	car := NewCar("red", 3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		car.AddPassedCorner(i)
	}
}

func BenchmarkCar_SetGear(b *testing.B) {
	car := NewCar("red", 3)
	discardPile := NewDiscardPile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gear := (i % 5) + 1
		_, _ = car.SetGear(gear, discardPile)
	}
}
