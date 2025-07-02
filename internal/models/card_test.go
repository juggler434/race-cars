package models

import (
	"reflect"
	"testing"
)

func TestNewCard(t *testing.T) {
	tests := []struct {
		name        string
		speed       int
		icons       map[Icon]int
		discardable bool
		playable    bool
		basic       bool
		wantName    string
		wantSpeed   int
		wantIcons   map[Icon]int
	}{
		{
			name:        "Speed Boost",
			speed:       5,
			icons:       map[Icon]int{IconBoost: 2},
			discardable: true,
			playable:    true,
			basic:       false,
			wantName:    "Speed Boost",
			wantSpeed:   5,
			wantIcons:   map[Icon]int{IconBoost: 2},
		},
		{
			name:        "Cooling System",
			speed:       0,
			icons:       map[Icon]int{IconCooling: 1},
			discardable: false,
			playable:    true,
			basic:       true,
			wantName:    "Cooling System",
			wantSpeed:   0,
			wantIcons:   map[Icon]int{IconCooling: 1},
		},
		{
			name:        "Empty Card",
			speed:       0,
			icons:       map[Icon]int{},
			discardable: false,
			playable:    false,
			basic:       true,
			wantName:    "Empty Card",
			wantSpeed:   0,
			wantIcons:   map[Icon]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := NewCard(tt.name, tt.speed, tt.icons, tt.discardable, tt.playable, tt.basic)

			// Test that the card implements the Card interface
			var _ Card = card

			// Test GetName
			if got := card.GetName(); got != tt.wantName {
				t.Errorf("NewCard().GetName() = %v, want %v", got, tt.wantName)
			}

			// Test GetSpeed
			if got := card.GetSpeed(); got != tt.wantSpeed {
				t.Errorf("NewCard().GetSpeed() = %v, want %v", got, tt.wantSpeed)
			}

			// Test GetIcons
			if got := card.GetIcons(); !reflect.DeepEqual(got, tt.wantIcons) {
				t.Errorf("NewCard().GetIcons() = %v, want %v", got, tt.wantIcons)
			}

			// Test IsDiscardable
			if got := card.IsDiscardable(); got != tt.discardable {
				t.Errorf("NewCard().IsDiscardable() = %v, want %v", got, tt.discardable)
			}

			// Test IsPlayable
			if got := card.IsPlayable(); got != tt.playable {
				t.Errorf("NewCard().IsPlayable() = %v, want %v", got, tt.playable)
			}

			// Test IsBasic
			if got := card.IsBasic(); got != tt.basic {
				t.Errorf("NewCard().IsBasic() = %v, want %v", got, tt.basic)
			}
		})
	}
}

func TestCard_GetName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Speed Boost", "Speed Boost"},
		{"Cooling System", "Cooling System"},
		{"", ""},
		{"Special Card with Spaces", "Special Card with Spaces"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := NewCard(tt.name, 0, map[Icon]int{}, false, false, false)
			if got := card.GetName(); got != tt.want {
				t.Errorf("Card.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_GetSpeed(t *testing.T) {
	tests := []struct {
		name  string
		speed int
		want  int
	}{
		{"Zero Speed", 0, 0},
		{"Low Speed", 1, 1},
		{"Medium Speed", 5, 5},
		{"High Speed", 10, 10},
		{"Negative Speed", -1, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := NewCard("Test Card", tt.speed, map[Icon]int{}, false, false, false)
			if got := card.GetSpeed(); got != tt.want {
				t.Errorf("Card.GetSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_GetIcons(t *testing.T) {
	tests := []struct {
		name  string
		icons map[Icon]int
		want  map[Icon]int
	}{
		{
			name:  "Empty Icons",
			icons: map[Icon]int{},
			want:  map[Icon]int{},
		},
		{
			name:  "Single Icon",
			icons: map[Icon]int{IconBoost: 1},
			want:  map[Icon]int{IconBoost: 1},
		},
		{
			name:  "Multiple Icons",
			icons: map[Icon]int{IconBoost: 2, IconCooling: 1},
			want:  map[Icon]int{IconBoost: 2, IconCooling: 1},
		},
		{
			name:  "Zero Count Icons",
			icons: map[Icon]int{IconBoost: 0, IconCooling: 0},
			want:  map[Icon]int{IconBoost: 0, IconCooling: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := NewCard("Test Card", 0, tt.icons, false, false, false)
			if got := card.GetIcons(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Card.GetIcons() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_IsDiscardable(t *testing.T) {
	tests := []struct {
		name        string
		discardable bool
		want        bool
	}{
		{"Discardable", true, true},
		{"Not Discardable", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := NewCard("Test Card", 0, map[Icon]int{}, tt.discardable, false, false)
			if got := card.IsDiscardable(); got != tt.want {
				t.Errorf("Card.IsDiscardable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_IsPlayable(t *testing.T) {
	tests := []struct {
		name     string
		playable bool
		want     bool
	}{
		{"Playable", true, true},
		{"Not Playable", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := NewCard("Test Card", 0, map[Icon]int{}, false, tt.playable, false)
			if got := card.IsPlayable(); got != tt.want {
				t.Errorf("Card.IsPlayable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_IsBasic(t *testing.T) {
	tests := []struct {
		name  string
		basic bool
		want  bool
	}{
		{"Basic Card", true, true},
		{"Special Card", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := NewCard("Test Card", 0, map[Icon]int{}, false, false, tt.basic)
			if got := card.IsBasic(); got != tt.want {
				t.Errorf("Card.IsBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_InterfaceCompliance(t *testing.T) {
	// Test that card struct properly implements Card interface
	var _ Card = (*card)(nil)

	// Test that NewCard returns a Card interface
	card := NewCard("Test", 0, map[Icon]int{}, false, false, false)
	var _ Card = card
}

func TestCard_ImmutableIcons(t *testing.T) {
	// Test that modifying the returned icons map doesn't affect the original card
	originalIcons := map[Icon]int{IconBoost: 1, IconCooling: 2}
	card := NewCard("Test Card", 0, originalIcons, false, false, false)

	// Get icons and modify them
	returnedIcons := card.GetIcons()
	returnedIcons[IconBoost] = 999
	delete(returnedIcons, IconCooling)

	// Check that original card is unchanged
	originalCardIcons := card.GetIcons()
	if originalCardIcons[IconBoost] != 1 {
		t.Errorf("Original card icons were modified: got %v, want %v", originalCardIcons[IconBoost], 1)
	}
	if originalCardIcons[IconCooling] != 2 {
		t.Errorf("Original card icons were modified: got %v, want %v", originalCardIcons[IconCooling], 2)
	}
}

func TestCard_EdgeCases(t *testing.T) {
	// Test with nil icons map
	card := NewCard("Nil Icons", 0, nil, false, false, false)
	icons := card.GetIcons()
	if icons == nil {
		t.Error("GetIcons() returned nil, expected empty map")
	}

	// Test with very large speed value
	largeSpeed := 999999
	card = NewCard("Large Speed", largeSpeed, map[Icon]int{}, false, false, false)
	if got := card.GetSpeed(); got != largeSpeed {
		t.Errorf("Card.GetSpeed() = %v, want %v", got, largeSpeed)
	}

	// Test with very long name
	longName := "This is a very long card name that might be used for testing purposes and should be handled correctly by the card system"
	card = NewCard(longName, 0, map[Icon]int{}, false, false, false)
	if got := card.GetName(); got != longName {
		t.Errorf("Card.GetName() = %v, want %v", got, longName)
	}
}

// Benchmark tests
func BenchmarkNewCard(b *testing.B) {
	icons := map[Icon]int{IconBoost: 2, IconCooling: 1}
	for i := 0; i < b.N; i++ {
		NewCard("Benchmark Card", 5, icons, true, true, false)
	}
}

func BenchmarkCard_GetName(b *testing.B) {
	card := NewCard("Benchmark Card", 5, map[Icon]int{}, true, true, false)
	for i := 0; i < b.N; i++ {
		card.GetName()
	}
}

func BenchmarkCard_GetSpeed(b *testing.B) {
	card := NewCard("Benchmark Card", 5, map[Icon]int{}, true, true, false)
	for i := 0; i < b.N; i++ {
		card.GetSpeed()
	}
}

func BenchmarkCard_GetIcons(b *testing.B) {
	icons := map[Icon]int{IconBoost: 2, IconCooling: 1}
	card := NewCard("Benchmark Card", 5, icons, true, true, false)
	for i := 0; i < b.N; i++ {
		card.GetIcons()
	}
}
