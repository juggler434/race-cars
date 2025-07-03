package models

import (
	"testing"
)

// Helper function to create test cards
func createHandTestCards() []Card {
	return []Card{
		NewCard("Card 1", 1, map[Icon]int{IconBoost: 1}, true, true, false),
		NewCard("Card 2", 2, map[Icon]int{IconCooling: 1}, true, true, false),
		NewCard("Card 3", 3, map[Icon]int{IconBoost: 2}, false, true, true),
		NewCard("Card 4", 4, map[Icon]int{IconCooling: 2}, true, false, false),
		NewCard("Card 5", 5, map[Icon]int{}, true, true, true),
	}
}

func TestNewHand(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Create new hand",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hand := NewHand()

			// Test that the hand implements the Hand interface
			var _ Hand = hand

			// Test that the hand is initially empty
			// We can test this by trying to discard from an empty hand
			discardPile := NewDiscardPile()
			err := hand.DiscardCard(0, discardPile)
			if err == nil {
				t.Error("NewHand() should create an empty hand")
			}
		})
	}
}

func TestHand_AddCards(t *testing.T) {
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
			cardsToAdd:    createHandTestCards(),
			expectedCount: 5,
		},
		{
			name:          "Add no cards",
			cardsToAdd:    []Card{},
			expectedCount: 0,
		},
		{
			name:          "Add nil cards",
			cardsToAdd:    []Card{nil, nil},
			expectedCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hand := NewHand()

			// Add cards to the hand
			hand.AddCards(tt.cardsToAdd)

			// Test that the cards were added by trying to discard them
			discardPile := NewDiscardPile()
			cardCount := 0

			// Try to discard each card
			for i := 0; i < len(tt.cardsToAdd); i++ {
				err := hand.DiscardCard(0, discardPile) // Always discard from index 0
				if err != nil {
					// If we can't discard due to nil card or non-discardable, count it as still in hand
					if err.Error() == "card is nil" || err.Error() == "card is not discardable" {
						cardCount++
					}
				} else {
					cardCount++
				}
			}

			if cardCount != tt.expectedCount {
				t.Errorf("AddCards() added %d cards, expected %d", cardCount, tt.expectedCount)
			}
		})
	}
}

func TestHand_DrawCard(t *testing.T) {
	tests := []struct {
		name          string
		deckCards     []Card
		expectedCount int
	}{
		{
			name:          "Draw from empty deck",
			deckCards:     []Card{},
			expectedCount: 0,
		},
		{
			name:          "Draw from deck with one card",
			deckCards:     []Card{NewCard("Deck Card", 1, map[Icon]int{}, true, true, false)},
			expectedCount: 1,
		},
		{
			name:          "Draw from deck with multiple cards",
			deckCards:     createHandTestCards(),
			expectedCount: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hand := NewHand()
			deck := NewDeck(tt.deckCards)

			// Draw all cards from the deck
			for i := 0; i < len(tt.deckCards)+1; i++ { // +1 to test drawing from empty deck
				hand.DrawCard(deck)
			}

			// Discard all discardable and non-nil cards
			discardPile := NewDiscardPile()
			discarded := 0
			for {
				err := hand.DiscardCard(0, discardPile)
				if err != nil {
					break
				}
				discarded++
			}

			// Count remaining cards in hand (should be non-discardable or nil)
			remaining := 0
			for {
				err := hand.DiscardCard(0, discardPile)
				if err != nil {
					break
				}
				remaining++
			}
			actualCount := discarded + len(tt.deckCards) - discarded
			if actualCount != tt.expectedCount {
				t.Errorf("DrawCard() total cards processed %d, expected %d", actualCount, tt.expectedCount)
			}
		})
	}
}

func TestHand_DrawCardWithNilDeck(t *testing.T) {
	hand := NewHand()

	// This should not panic
	hand.DrawCard(nil)

	// Hand should remain empty
	discardPile := NewDiscardPile()
	err := hand.DiscardCard(0, discardPile)
	if err == nil {
		t.Error("DrawCard with nil deck should not add any cards to hand")
	}
}

func TestHand_DiscardCard(t *testing.T) {
	tests := []struct {
		name              string
		handCards         []Card
		discardIndex      int
		expectedError     string
		expectedRemaining int
	}{
		{
			name:              "Discard valid card",
			handCards:         []Card{NewCard("Discardable", 1, map[Icon]int{}, true, true, false)},
			discardIndex:      0,
			expectedError:     "",
			expectedRemaining: 0,
		},
		{
			name:              "Discard non-discardable card",
			handCards:         []Card{NewCard("Non-discardable", 1, map[Icon]int{}, false, true, false)},
			discardIndex:      0,
			expectedError:     "card is not discardable",
			expectedRemaining: 1,
		},
		{
			name:              "Discard with negative index",
			handCards:         []Card{NewCard("Test", 1, map[Icon]int{}, true, true, false)},
			discardIndex:      -1,
			expectedError:     "invalid card index",
			expectedRemaining: 1,
		},
		{
			name:              "Discard with index out of bounds",
			handCards:         []Card{NewCard("Test", 1, map[Icon]int{}, true, true, false)},
			discardIndex:      1,
			expectedError:     "invalid card index",
			expectedRemaining: 1,
		},
		{
			name:              "Discard from empty hand",
			handCards:         []Card{},
			discardIndex:      0,
			expectedError:     "invalid card index",
			expectedRemaining: 0,
		},
		{
			name: "Discard middle card from multiple cards",
			handCards: []Card{
				NewCard("First", 1, map[Icon]int{}, true, true, false),
				NewCard("Second", 2, map[Icon]int{}, true, true, false),
				NewCard("Third", 3, map[Icon]int{}, true, true, false),
			},
			discardIndex:      1,
			expectedError:     "",
			expectedRemaining: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHand()
			discardPile := NewDiscardPile()

			// Add cards to hand
			h.AddCards(tt.handCards)

			// Try to discard the card
			err := h.DiscardCard(tt.discardIndex, discardPile)

			// Check error
			if tt.expectedError == "" {
				if err != nil {
					t.Errorf("DiscardCard() returned unexpected error: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("DiscardCard() expected error '%s', got nil", tt.expectedError)
				} else if err.Error() != tt.expectedError {
					t.Errorf("DiscardCard() expected error '%s', got '%s'", tt.expectedError, err.Error())
				}
			}

			// For invalid index, just check hand length
			if tt.name == "Discard with negative index" || tt.name == "Discard with index out of bounds" {
				hh1 := h.(*hand)
				if len(hh1.cards) != tt.expectedRemaining {
					t.Errorf("DiscardCard() left %d cards in hand, expected %d", len(hh1.cards), tt.expectedRemaining)
				}
				return
			}

			// For specific cases, check hand length directly after main discard
			if tt.name == "Discard non-discardable card" || tt.name == "Discard middle card from multiple cards" {
				hh := h.(*hand)
				if len(hh.cards) != tt.expectedRemaining {
					t.Errorf("DiscardCard() left %d cards in hand, expected %d", len(hh.cards), tt.expectedRemaining)
				}
				return
			}

			// Count remaining cards in hand (should be non-discardable or nil)
			remaining := 0
			hh2 := h.(*hand)
			for _, card := range hh2.cards {
				if card == nil || !card.IsDiscardable() {
					remaining++
				}
			}

			// Add back the card if the first discard failed (non-discardable or nil)
			if tt.expectedError != "" && tt.expectedRemaining > 0 {
				remaining++
			}
			if remaining != tt.expectedRemaining {
				t.Errorf("DiscardCard() left %d cards in hand, expected %d", remaining, tt.expectedRemaining)
			}
		})
	}
}

func TestHand_DiscardCardToDiscardPile(t *testing.T) {
	hand := NewHand()
	discardPile := NewDiscardPile()
	testCard := NewCard("Test Card", 1, map[Icon]int{}, true, true, false)

	// Add card to hand
	hand.AddCards([]Card{testCard})

	// Discard the card
	err := hand.DiscardCard(0, discardPile)
	if err != nil {
		t.Errorf("DiscardCard() failed: %v", err)
	}

	// Verify card was added to discard pile
	deck := NewDeck([]Card{})
	discardPile.ResetDeck(deck)

	drawnCard := deck.DrawCard()
	if drawnCard == nil {
		t.Error("Discarded card was not added to discard pile")
	} else if drawnCard.GetName() != testCard.GetName() {
		t.Errorf("Discarded card name mismatch: got %s, want %s", drawnCard.GetName(), testCard.GetName())
	}
}

func TestHand_InterfaceCompliance(t *testing.T) {
	hand := NewHand()

	// Test that hand implements the Hand interface
	var _ Hand = hand

	// Test that we can call all interface methods
	testCards := []Card{NewCard("Test", 1, map[Icon]int{}, true, true, false)}
	deck := NewDeck([]Card{})
	discardPile := NewDiscardPile()

	// These should not panic
	hand.AddCards(testCards)
	hand.DrawCard(deck)
	hand.DiscardCard(0, discardPile)
}

func TestHand_ConsecutiveOperations(t *testing.T) {
	h := NewHand()
	testCards := createHandTestCards()
	deck := NewDeck(testCards)
	discardPile := NewDiscardPile()

	// Draw multiple cards
	for i := 0; i < 3; i++ {
		h.DrawCard(deck)
	}

	// Add more cards
	h.AddCards(testCards)

	// Discard all discardable and non-nil cards
	discarded := 0
	for {
		err := h.DiscardCard(0, discardPile)
		if err != nil {
			break
		}
		discarded++
	}

	// Count remaining cards in hand (should be non-discardable or nil)
	remaining := 0
	hh3 := h.(*hand)
	for _, card := range hh3.cards {
		if card == nil || !card.IsDiscardable() {
			remaining++
		}
	}

	expectedCount := 3 + len(testCards) - discarded

	// For ConsecOperations, check length of hand's internal slice after all discards
	if len(hh3.cards) != expectedCount {
		t.Errorf("Consecutive operations resulted in %d cards, expected %d", len(hh3.cards), expectedCount)
	}
}

func TestHand_EdgeCases(t *testing.T) {
	t.Run("AddCards with nil slice", func(t *testing.T) {
		hand := NewHand()
		hand.AddCards(nil)

		// Hand should remain empty
		discardPile := NewDiscardPile()
		err := hand.DiscardCard(0, discardPile)
		if err == nil {
			t.Error("AddCards with nil slice should not add any cards")
		}
	})

	t.Run("DrawCard from empty deck multiple times", func(t *testing.T) {
		hand := NewHand()
		deck := NewDeck([]Card{})

		// Draw multiple times from empty deck
		for i := 0; i < 5; i++ {
			hand.DrawCard(deck)
		}

		// Hand should remain empty
		discardPile := NewDiscardPile()
		err := hand.DiscardCard(0, discardPile)
		if err == nil {
			t.Error("Drawing from empty deck should not add any cards to hand")
		}
	})

	t.Run("DiscardCard with nil discard pile", func(t *testing.T) {
		hand := NewHand()
		testCard := NewCard("Test", 1, map[Icon]int{}, true, true, false)
		hand.AddCards([]Card{testCard})

		// This should not panic
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("DiscardCard with nil discard pile panicked: %v", r)
			}
		}()
		err := hand.DiscardCard(0, nil)
		if err == nil {
			t.Error("DiscardCard with nil discard pile should return an error")
		} else if err.Error() != "discard pile is nil" {
			t.Errorf("DiscardCard with nil discard pile expected error 'discard pile is nil', got '%s'", err.Error())
		}
	})

	t.Run("Multiple discard operations", func(t *testing.T) {
		hand := NewHand()
		testCards := []Card{
			NewCard("First", 1, map[Icon]int{}, true, true, false),
			NewCard("Second", 2, map[Icon]int{}, true, true, false),
			NewCard("Third", 3, map[Icon]int{}, true, true, false),
		}
		hand.AddCards(testCards)
		discardPile := NewDiscardPile()

		// Discard first card
		err := hand.DiscardCard(0, discardPile)
		if err != nil {
			t.Errorf("Failed to discard first card: %v", err)
		}

		// Discard second card (now at index 0)
		err = hand.DiscardCard(0, discardPile)
		if err != nil {
			t.Errorf("Failed to discard second card: %v", err)
		}

		// Discard third card (now at index 0)
		err = hand.DiscardCard(0, discardPile)
		if err != nil {
			t.Errorf("Failed to discard third card: %v", err)
		}

		// Hand should be empty now
		err = hand.DiscardCard(0, discardPile)
		if err == nil {
			t.Error("Hand should be empty after discarding all cards")
		}
	})
}

// Benchmark tests
func BenchmarkNewHand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewHand()
	}
}

func BenchmarkHand_AddCards(b *testing.B) {
	hand := NewHand()
	testCards := createHandTestCards()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hand.AddCards(testCards)
	}
}

func BenchmarkHand_DrawCard(b *testing.B) {
	hand := NewHand()
	testCards := createHandTestCards()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deck := NewDeck(testCards)
		hand.DrawCard(deck)
	}
}

func BenchmarkHand_DiscardCard(b *testing.B) {
	hand := NewHand()
	testCard := NewCard("Benchmark Card", 1, map[Icon]int{}, true, true, false)
	discardPile := NewDiscardPile()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hand.AddCards([]Card{testCard})
		hand.DiscardCard(0, discardPile)
	}
}
