package models

import (
	"reflect"
	"testing"
)

// Helper function to create test cards
func createDiscardPileTestCards() []Card {
	return []Card{
		NewCard("Card 1", 1, map[Icon]int{IconBoost: 1}, true, true, false),
		NewCard("Card 2", 2, map[Icon]int{IconCooling: 1}, true, true, false),
		NewCard("Card 3", 3, map[Icon]int{IconBoost: 2}, false, true, true),
		NewCard("Card 4", 4, map[Icon]int{IconCooling: 2}, true, false, false),
		NewCard("Card 5", 5, map[Icon]int{}, true, true, true),
	}
}

func TestNewDiscardPile(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Create new discard pile",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			discardPile := NewDiscardPile()

			// Test that the discard pile implements the DiscardPile interface
			var _ DiscardPile = discardPile

			// Test that the discard pile is initially empty
			// We can't directly access the cards slice, but we can test through ResetDeck
			deck := NewDeck([]Card{})
			discardPile.ResetDeck(deck)

			// After resetting an empty discard pile, the deck should still be empty
			if deck.DrawCard() != nil {
				t.Error("NewDiscardPile() should create an empty discard pile")
			}
		})
	}
}

func TestDiscardPile_AddCard(t *testing.T) {
	tests := []struct {
		name          string
		cardsToAdd    []Card
		expectedCount int
	}{
		{
			name:          "Add single card",
			cardsToAdd:    []Card{NewCard("Test Card", 1, map[Icon]int{}, true, true, false)},
			expectedCount: 1,
		},
		{
			name:          "Add multiple cards",
			cardsToAdd:    createDiscardPileTestCards(),
			expectedCount: 5,
		},
		{
			name:          "Add no cards",
			cardsToAdd:    []Card{},
			expectedCount: 0,
		},
		{
			name:          "Add nil card",
			cardsToAdd:    []Card{nil},
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			discardPile := NewDiscardPile()

			// Add cards to the discard pile
			for _, card := range tt.cardsToAdd {
				discardPile.AddCard(card)
			}

			// Test that the cards were added by resetting to a deck and counting
			deck := NewDeck([]Card{})
			discardPile.ResetDeck(deck)

			// Count cards in the deck
			cardCount := 0
			for deck.DrawCard() != nil {
				cardCount++
			}

			if cardCount != tt.expectedCount {
				t.Errorf("AddCard() added %d cards, expected %d", cardCount, tt.expectedCount)
			}
		})
	}
}

func TestDiscardPile_ResetDeck(t *testing.T) {
	tests := []struct {
		name              string
		discardPileCards  []Card
		deckCards         []Card
		expectedDeckCount int
	}{
		{
			name:              "Reset empty discard pile to empty deck",
			discardPileCards:  []Card{},
			deckCards:         []Card{},
			expectedDeckCount: 0,
		},
		{
			name:              "Reset cards to empty deck",
			discardPileCards:  createDiscardPileTestCards(),
			deckCards:         []Card{},
			expectedDeckCount: 5,
		},
		{
			name:              "Reset cards to existing deck",
			discardPileCards:  []Card{NewCard("Discard 1", 1, map[Icon]int{}, true, true, false)},
			deckCards:         []Card{NewCard("Deck 1", 2, map[Icon]int{}, true, true, false)},
			expectedDeckCount: 2,
		},
		{
			name: "Reset multiple cards to existing deck",
			discardPileCards: []Card{
				NewCard("Discard 1", 1, map[Icon]int{}, true, true, false),
				NewCard("Discard 2", 2, map[Icon]int{}, true, true, false),
			},
			deckCards:         []Card{NewCard("Deck 1", 3, map[Icon]int{}, true, true, false)},
			expectedDeckCount: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			discardPile := NewDiscardPile()
			deck := NewDeck(tt.deckCards)

			// Add cards to the discard pile
			for _, card := range tt.discardPileCards {
				discardPile.AddCard(card)
			}

			// Reset the deck
			discardPile.ResetDeck(deck)

			// Verify the deck has the expected number of cards
			cardCount := 0
			drawnCards := make([]Card, 0)
			for {
				card := deck.DrawCard()
				if card == nil {
					break
				}
				drawnCards = append(drawnCards, card)
				cardCount++
			}

			if cardCount != tt.expectedDeckCount {
				t.Errorf("ResetDeck() resulted in %d cards in deck, expected %d", cardCount, tt.expectedDeckCount)
			}

			// Verify we have the right cards (order may be different due to shuffling)
			if len(drawnCards) != tt.expectedDeckCount {
				t.Errorf("ResetDeck() resulted in %d cards, expected %d", len(drawnCards), tt.expectedDeckCount)
			}
		})
	}
}

func TestDiscardPile_ResetDeckShuffles(t *testing.T) {
	// Create a discard pile with known cards
	testCards := createDiscardPileTestCards()
	discardPile := NewDiscardPile()

	for _, card := range testCards {
		discardPile.AddCard(card)
	}

	// Create an empty deck
	deck := NewDeck([]Card{})

	// Reset the deck (this should shuffle it)
	discardPile.ResetDeck(deck)

	// Draw all cards and get their order
	drawnCards := make([]Card, 0)
	for {
		card := deck.DrawCard()
		if card == nil {
			break
		}
		drawnCards = append(drawnCards, card)
	}

	// Verify we have the same number of cards
	if len(drawnCards) != len(testCards) {
		t.Errorf("ResetDeck() changed number of cards: got %d, want %d", len(drawnCards), len(testCards))
	}

	// Create maps to compare card counts (order might be different due to shuffling)
	originalCounts := make(map[string]int)
	drawnCounts := make(map[string]int)

	for _, card := range testCards {
		originalCounts[card.GetName()]++
	}
	for _, card := range drawnCards {
		drawnCounts[card.GetName()]++
	}

	if !reflect.DeepEqual(originalCounts, drawnCounts) {
		t.Errorf("ResetDeck() changed card composition: got %v, want %v", drawnCounts, originalCounts)
	}
}

func TestDiscardPile_ResetDeckEmptiesPile(t *testing.T) {
	discardPile := NewDiscardPile()
	testCards := createDiscardPileTestCards()

	// Add cards to the discard pile
	for _, card := range testCards {
		discardPile.AddCard(card)
	}

	// Reset to a deck
	deck := NewDeck([]Card{})
	discardPile.ResetDeck(deck)

	// Verify the discard pile is now empty by resetting again
	emptyDeck := NewDeck([]Card{})
	discardPile.ResetDeck(emptyDeck)

	// The empty deck should remain empty
	if emptyDeck.DrawCard() != nil {
		t.Error("ResetDeck() should empty the discard pile after resetting")
	}
}

func TestDiscardPile_InterfaceCompliance(t *testing.T) {
	discardPile := NewDiscardPile()

	// Test that discardPile implements the DiscardPile interface
	var _ DiscardPile = discardPile

	// Test that we can call all interface methods
	testCard := NewCard("Test", 1, map[Icon]int{}, true, true, false)
	deck := NewDeck([]Card{})

	// These should not panic
	discardPile.AddCard(testCard)
	discardPile.ResetDeck(deck)
}

func TestDiscardPile_ConsecutiveOperations(t *testing.T) {
	discardPile := NewDiscardPile()
	testCards := createDiscardPileTestCards()

	// Add cards multiple times
	for i := 0; i < 3; i++ {
		for _, card := range testCards {
			discardPile.AddCard(card)
		}
	}

	// Reset to deck
	deck := NewDeck([]Card{})
	discardPile.ResetDeck(deck)

	// Count total cards
	cardCount := 0
	for deck.DrawCard() != nil {
		cardCount++
	}

	expectedCount := len(testCards) * 3
	if cardCount != expectedCount {
		t.Errorf("Consecutive AddCard operations resulted in %d cards, expected %d", cardCount, expectedCount)
	}
}

func TestDiscardPile_EdgeCases(t *testing.T) {
	t.Run("AddCard with nil card", func(t *testing.T) {
		discardPile := NewDiscardPile()
		discardPile.AddCard(nil)

		deck := NewDeck([]Card{})
		discardPile.ResetDeck(deck)

		card := deck.DrawCard()
		if card != nil {
			t.Error("AddCard with nil should not add anything to the discard pile")
		}
	})

	t.Run("ResetDeck with nil deck", func(t *testing.T) {
		discardPile := NewDiscardPile()
		testCard := NewCard("Test", 1, map[Icon]int{}, true, true, false)
		discardPile.AddCard(testCard)

		// This should not panic - we need to handle nil deck gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("ResetDeck with nil deck panicked: %v", r)
			}
		}()
		discardPile.ResetDeck(nil)
	})

	t.Run("Multiple ResetDeck calls", func(t *testing.T) {
		discardPile := NewDiscardPile()
		testCard := NewCard("Test", 1, map[Icon]int{}, true, true, false)
		discardPile.AddCard(testCard)

		deck1 := NewDeck([]Card{})
		deck2 := NewDeck([]Card{})

		// Reset to first deck
		discardPile.ResetDeck(deck1)

		// Reset to second deck (should be empty)
		discardPile.ResetDeck(deck2)

		if deck2.DrawCard() != nil {
			t.Error("Second ResetDeck should not add any cards since discard pile was emptied")
		}
	})
}

// Benchmark tests
func BenchmarkNewDiscardPile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewDiscardPile()
	}
}

func BenchmarkDiscardPile_AddCard(b *testing.B) {
	discardPile := NewDiscardPile()
	testCard := NewCard("Benchmark Card", 1, map[Icon]int{}, true, true, false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		discardPile.AddCard(testCard)
	}
}

func BenchmarkDiscardPile_ResetDeck(b *testing.B) {
	discardPile := NewDiscardPile()
	testCards := createDiscardPileTestCards()

	// Add cards to discard pile
	for _, card := range testCards {
		discardPile.AddCard(card)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deck := NewDeck([]Card{})
		discardPile.ResetDeck(deck)

		// Re-add cards for next iteration
		for _, card := range testCards {
			discardPile.AddCard(card)
		}
	}
}
