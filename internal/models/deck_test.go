package models

import (
	"reflect"
	"testing"
)

// Helper function to create test cards
func createTestCards() []Card {
	return []Card{
		NewCard("Card 1", 1, map[Icon]int{IconBoost: 1}, true, true, false),
		NewCard("Card 2", 2, map[Icon]int{IconCooling: 1}, true, true, false),
		NewCard("Card 3", 3, map[Icon]int{IconBoost: 2}, false, true, true),
		NewCard("Card 4", 4, map[Icon]int{IconCooling: 2}, true, false, false),
		NewCard("Card 5", 5, map[Icon]int{}, true, true, true),
	}
}

func TestNewDeck(t *testing.T) {
	tests := []struct {
		name  string
		cards []Card
		want  int
	}{
		{
			name:  "Empty Deck",
			cards: []Card{},
			want:  0,
		},
		{
			name:  "Single Card",
			cards: []Card{NewCard("Test", 1, map[Icon]int{}, true, true, false)},
			want:  1,
		},
		{
			name:  "Multiple Cards",
			cards: createTestCards(),
			want:  5,
		},
		{
			name:  "Nil Cards",
			cards: nil,
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deck := NewDeck(tt.cards)

			// Test that the deck implements the Deck interface
			var _ Deck = deck

			// Test that deck has correct number of cards
			cardCount := 0
			for deck.DrawCard() != nil {
				cardCount++
			}

			if cardCount != tt.want {
				t.Errorf("NewDeck() created deck with %d cards, want %d", cardCount, tt.want)
			}
		})
	}
}

func TestDeck_DrawCard(t *testing.T) {
	tests := []struct {
		name          string
		cards         []Card
		draws         int
		expectedCards []string
		expectedNil   bool
	}{
		{
			name:          "Draw from empty deck",
			cards:         []Card{},
			draws:         1,
			expectedCards: []string{},
			expectedNil:   true,
		},
		{
			name:          "Draw single card",
			cards:         []Card{NewCard("Test Card", 1, map[Icon]int{}, true, true, false)},
			draws:         1,
			expectedCards: []string{"Test Card"},
			expectedNil:   false,
		},
		{
			name:          "Draw multiple cards in order",
			cards:         createTestCards(),
			draws:         3,
			expectedCards: []string{"Card 1", "Card 2", "Card 3"},
			expectedNil:   false,
		},
		{
			name:          "Draw more cards than available",
			cards:         []Card{NewCard("Test Card", 1, map[Icon]int{}, true, true, false)},
			draws:         3,
			expectedCards: []string{"Test Card"},
			expectedNil:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deck := NewDeck(tt.cards)

			drawnCards := make([]Card, 0)
			for i := 0; i < tt.draws; i++ {
				card := deck.DrawCard()
				if card != nil {
					drawnCards = append(drawnCards, card)
				}
			}

			// Check number of cards drawn
			if len(drawnCards) != len(tt.expectedCards) {
				t.Errorf("Drew %d cards, expected %d", len(drawnCards), len(tt.expectedCards))
			}

			// Check card names
			for i, card := range drawnCards {
				if card.GetName() != tt.expectedCards[i] {
					t.Errorf("Card %d: got %s, want %s", i, card.GetName(), tt.expectedCards[i])
				}
			}

			// Check if we got nil when expected
			if len(drawnCards) < tt.draws && !tt.expectedNil {
				t.Error("Expected to draw more cards but got nil")
			}
		})
	}
}

func TestDeck_Shuffle(t *testing.T) {
	// Create a deck with known cards
	originalCards := createTestCards()
	deck := NewDeck(originalCards)

	// Get original order
	originalOrder := make([]string, len(originalCards))
	for i, card := range originalCards {
		originalOrder[i] = card.GetName()
	}

	// Shuffle the deck
	deck.Shuffle()

	// Draw all cards and check if order changed
	shuffledOrder := make([]string, 0)
	for {
		card := deck.DrawCard()
		if card == nil {
			break
		}
		shuffledOrder = append(shuffledOrder, card.GetName())
	}

	// Check that we have the same cards (order might be different)
	if len(shuffledOrder) != len(originalOrder) {
		t.Errorf("Shuffle changed number of cards: got %d, want %d", len(shuffledOrder), len(originalOrder))
	}

	// Create maps to compare card counts
	originalCounts := make(map[string]int)
	shuffledCounts := make(map[string]int)

	for _, name := range originalOrder {
		originalCounts[name]++
	}
	for _, name := range shuffledOrder {
		shuffledCounts[name]++
	}

	if !reflect.DeepEqual(originalCounts, shuffledCounts) {
		t.Errorf("Shuffle changed card composition: got %v, want %v", shuffledCounts, originalCounts)
	}

	// Note: We can't guarantee the order changed due to randomness, but we can test that shuffle doesn't break the deck
}

func TestDeck_AddCardsToTop(t *testing.T) {
	tests := []struct {
		name          string
		initialCards  []Card
		cardsToAdd    []Card
		expectedOrder []string
		expectedCount int
	}{
		{
			name:          "Add to empty deck",
			initialCards:  []Card{},
			cardsToAdd:    []Card{NewCard("New Card", 1, map[Icon]int{}, true, true, false)},
			expectedOrder: []string{"New Card"},
			expectedCount: 1,
		},
		{
			name:          "Add single card to existing deck",
			initialCards:  []Card{NewCard("Original", 1, map[Icon]int{}, true, true, false)},
			cardsToAdd:    []Card{NewCard("New Card", 2, map[Icon]int{}, true, true, false)},
			expectedOrder: []string{"New Card", "Original"},
			expectedCount: 2,
		},
		{
			name:          "Add multiple cards to existing deck",
			initialCards:  []Card{NewCard("Original", 1, map[Icon]int{}, true, true, false)},
			cardsToAdd:    []Card{NewCard("New 1", 2, map[Icon]int{}, true, true, false), NewCard("New 2", 3, map[Icon]int{}, true, true, false)},
			expectedOrder: []string{"New 1", "New 2", "Original"},
			expectedCount: 3,
		},
		{
			name:          "Add empty slice",
			initialCards:  createTestCards(),
			cardsToAdd:    []Card{},
			expectedOrder: []string{"Card 1", "Card 2", "Card 3", "Card 4", "Card 5"},
			expectedCount: 5,
		},
		{
			name:          "Add nil slice",
			initialCards:  createTestCards(),
			cardsToAdd:    nil,
			expectedOrder: []string{"Card 1", "Card 2", "Card 3", "Card 4", "Card 5"},
			expectedCount: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deck := NewDeck(tt.initialCards)

			// Add cards to top
			deck.AddCardsToTop(tt.cardsToAdd)

			// Draw all cards and check order
			drawnCards := make([]string, 0)
			for {
				card := deck.DrawCard()
				if card == nil {
					break
				}
				drawnCards = append(drawnCards, card.GetName())
			}

			// Check count
			if len(drawnCards) != tt.expectedCount {
				t.Errorf("Deck has %d cards, expected %d", len(drawnCards), tt.expectedCount)
			}

			// Check order
			if !reflect.DeepEqual(drawnCards, tt.expectedOrder) {
				t.Errorf("Card order: got %v, want %v", drawnCards, tt.expectedOrder)
			}
		})
	}
}

func TestDeck_InterfaceCompliance(t *testing.T) {
	// Test that deck struct properly implements Deck interface
	var _ Deck = (*deck)(nil)

	// Test that NewDeck returns a Deck interface
	cards := createTestCards()
	deck := NewDeck(cards)
	var _ Deck = deck
}

func TestDeck_EmptyDeckOperations(t *testing.T) {
	// Test operations on empty deck
	deck := NewDeck([]Card{})

	// Draw from empty deck should return nil
	card := deck.DrawCard()
	if card != nil {
		t.Error("Drawing from empty deck should return nil")
	}

	// Shuffle empty deck should not panic
	deck.Shuffle()

	// Add cards to empty deck
	newCards := []Card{NewCard("Test", 1, map[Icon]int{}, true, true, false)}
	deck.AddCardsToTop(newCards)

	// Now should be able to draw
	card = deck.DrawCard()
	if card == nil {
		t.Error("Should be able to draw after adding cards")
	}
	if card.GetName() != "Test" {
		t.Errorf("Drew wrong card: got %s, want Test", card.GetName())
	}
}

func TestDeck_ConsecutiveOperations(t *testing.T) {
	// Test multiple operations in sequence
	cards := createTestCards()
	deck := NewDeck(cards)

	// Draw first card
	firstCard := deck.DrawCard()
	if firstCard == nil {
		t.Fatal("Failed to draw first card")
	}

	// Add card to top
	newCard := NewCard("Added Card", 10, map[Icon]int{}, true, true, false)
	deck.AddCardsToTop([]Card{newCard})

	// Shuffle
	deck.Shuffle()

	// Draw remaining cards
	drawnCards := make([]string, 0)
	for {
		card := deck.DrawCard()
		if card == nil {
			break
		}
		drawnCards = append(drawnCards, card.GetName())
	}

	// Should have 5 cards left (original 5 - 1 drawn + 1 added)
	if len(drawnCards) != 5 {
		t.Errorf("Expected 4 cards after operations, got %d", len(drawnCards))
	}

	// Should not contain the first drawn card
	for _, name := range drawnCards {
		if name == firstCard.GetName() {
			t.Errorf("First drawn card %s should not be in remaining cards", name)
		}
	}
}

func TestDeck_EdgeCases(t *testing.T) {
	// Test with very large deck
	largeDeck := make([]Card, 1000)
	for i := 0; i < 1000; i++ {
		largeDeck[i] = NewCard("Card "+string(rune(i)), i, map[Icon]int{}, true, true, false)
	}

	deck := NewDeck(largeDeck)

	// Shuffle large deck
	deck.Shuffle()

	// Draw all cards
	cardCount := 0
	for deck.DrawCard() != nil {
		cardCount++
	}

	if cardCount != 1000 {
		t.Errorf("Large deck lost cards: got %d, want 1000", cardCount)
	}

	// Test with cards that have nil icons
	nilIconsCard := NewCard("Nil Icons", 1, nil, true, true, false)
	deck = NewDeck([]Card{nilIconsCard})

	card := deck.DrawCard()
	if card == nil {
		t.Error("Failed to draw card with nil icons")
	}
	if card.GetName() != "Nil Icons" {
		t.Errorf("Drew wrong card: got %s, want Nil Icons", card.GetName())
	}
}

// Benchmark tests
func BenchmarkNewDeck(b *testing.B) {
	cards := createTestCards()
	for i := 0; i < b.N; i++ {
		NewDeck(cards)
	}
}

func BenchmarkDeck_DrawCard(b *testing.B) {
	cards := createTestCards()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Reset deck for each iteration
		deck := NewDeck(cards)
		for deck.DrawCard() != nil {
			// Draw all cards
		}
	}
}

func BenchmarkDeck_Shuffle(b *testing.B) {
	cards := createTestCards()
	deck := NewDeck(cards)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		deck.Shuffle()
	}
}

func BenchmarkDeck_AddCardsToTop(b *testing.B) {
	cards := createTestCards()
	newCards := []Card{NewCard("New", 1, map[Icon]int{}, true, true, false)}
	deck := NewDeck(cards)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		deck.AddCardsToTop(newCards)
	}
}
