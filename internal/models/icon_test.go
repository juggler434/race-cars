package models

import (
	"testing"
)

func TestIconConstants(t *testing.T) {
	tests := []struct {
		name     string
		icon     Icon
		expected int
	}{
		{
			name:     "IconBoost should be 0",
			icon:     IconBoost,
			expected: 0,
		},
		{
			name:     "IconCooling should be 1",
			icon:     IconCooling,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if int(tt.icon) != tt.expected {
				t.Errorf("Icon %s has value %d, expected %d", tt.name, int(tt.icon), tt.expected)
			}
		})
	}
}

func TestIcon_String(t *testing.T) {
	tests := []struct {
		name     string
		icon     Icon
		expected string
	}{
		{
			name:     "IconBoost should return 'Boost'",
			icon:     IconBoost,
			expected: "Boost",
		},
		{
			name:     "IconCooling should return 'Cooling'",
			icon:     IconCooling,
			expected: "Cooling",
		},
		{
			name:     "Unknown icon should return empty string",
			icon:     Icon(999),
			expected: "",
		},
		{
			name:     "Negative icon should return empty string",
			icon:     Icon(-1),
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.icon.String()
			if result != tt.expected {
				t.Errorf("Icon.String() = %s, expected %s", result, tt.expected)
			}
		})
	}
}

func TestIconTypeConversion(t *testing.T) {
	tests := []struct {
		name     string
		icon     Icon
		expected string
	}{
		{
			name:     "Convert IconBoost to string",
			icon:     IconBoost,
			expected: "Boost",
		},
		{
			name:     "Convert IconCooling to string",
			icon:     IconCooling,
			expected: "Cooling",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test direct String() method
			result := tt.icon.String()
			if result != tt.expected {
				t.Errorf("Icon.String() = %s, expected %s", result, tt.expected)
			}

			// Test fmt.Sprintf conversion
			formatted := tt.icon.String()
			if formatted != tt.expected {
				t.Errorf("fmt.Sprintf(\"%%s\", %s) = %s, expected %s", tt.name, formatted, tt.expected)
			}
		})
	}
}

func TestIconComparison(t *testing.T) {
	tests := []struct {
		name     string
		icon1    Icon
		icon2    Icon
		expected bool
	}{
		{
			name:     "IconBoost equals IconBoost",
			icon1:    IconBoost,
			icon2:    IconBoost,
			expected: true,
		},
		{
			name:     "IconCooling equals IconCooling",
			icon1:    IconCooling,
			icon2:    IconCooling,
			expected: true,
		},
		{
			name:     "IconBoost not equals IconCooling",
			icon1:    IconBoost,
			icon2:    IconCooling,
			expected: false,
		},
		{
			name:     "IconCooling not equals IconBoost",
			icon1:    IconCooling,
			icon2:    IconBoost,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.icon1 == tt.icon2
			if result != tt.expected {
				t.Errorf("Icon comparison %s == %s = %t, expected %t",
					tt.icon1.String(), tt.icon2.String(), result, tt.expected)
			}
		})
	}
}

func TestIconUniqueness(t *testing.T) {
	// Test that all icons have unique values
	icons := []Icon{IconBoost, IconCooling}
	seen := make(map[Icon]bool)

	for _, icon := range icons {
		if seen[icon] {
			t.Errorf("Icon %s has duplicate value %d", icon.String(), int(icon))
		}
		seen[icon] = true
	}
}

func TestIconStringUniqueness(t *testing.T) {
	// Test that all icons have unique string representations
	icons := []Icon{IconBoost, IconCooling}
	seen := make(map[string]bool)

	for _, icon := range icons {
		name := icon.String()
		if seen[name] {
			t.Errorf("Icon %s has duplicate string representation '%s'", icon.String(), name)
		}
		seen[name] = true
	}
}

func TestIconEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		icon     Icon
		expected string
	}{
		{
			name:     "Zero value icon",
			icon:     Icon(0),
			expected: "Boost",
		},
		{
			name:     "Large positive icon",
			icon:     Icon(1000),
			expected: "",
		},
		{
			name:     "Negative icon",
			icon:     Icon(-5),
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.icon.String()
			if result != tt.expected {
				t.Errorf("Icon.String() for %s = %s, expected %s", tt.name, result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkIcon_String(b *testing.B) {
	icon := IconBoost
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = icon.String()
	}
}

func BenchmarkIconBoost_String(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = IconBoost.String()
	}
}

func BenchmarkIconCooling_String(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = IconCooling.String()
	}
}

func BenchmarkIconComparison(b *testing.B) {
	icon1 := IconBoost
	icon2 := IconCooling
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = icon1 == icon2
	}
}
