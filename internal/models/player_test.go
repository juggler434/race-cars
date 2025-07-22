package models

import (
	"testing"
)

// Helper function to create test cards
func createPlayerTestCards() []Card {
	return []Card{
		NewCard("Card 1", 1, map[Icon]int{IconBoost: 1}, true, true, false),
		NewCard("Card 2", 2, map[Icon]int{IconCooling: 1}, true, true, false),
		NewCard("Card 3", 3, map[Icon]int{IconBoost: 2}, false, true, true),
		NewCard("Card 4", 4, map[Icon]int{IconCooling: 2}, true, false, false),
		NewCard("Card 5", 5, map[Icon]int{}, true, true, true),
	}
}

func TestNewPlayer(t *testing.T) {
	tests := []struct {
		name        string
		playerName  string
		car         Car
		discardPile DiscardPile
		deck        Deck
		hand        Hand
	}{
		{
			name:        "Create player with all components",
			playerName:  "TestPlayer",
			car:         NewCar("red", 3),
			discardPile: NewDiscardPile(),
			deck:        NewDeck(createPlayerTestCards()),
			hand:        NewHand(),
		},
		{
			name:        "Create player with empty components",
			playerName:  "EmptyPlayer",
			car:         NewCar("blue", 2),
			discardPile: NewDiscardPile(),
			deck:        NewDeck([]Card{}),
			hand:        NewHand(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player := NewPlayer(tt.playerName, tt.car, tt.discardPile, tt.deck, tt.hand)

			// Test that the player implements the Player interface
			var _ Player = player

			// Test that all components are correctly set
			if player.GetName() != tt.playerName {
				t.Errorf("GetName() = %s, want %s", player.GetName(), tt.playerName)
			}
			if player.GetCar() != tt.car {
				t.Errorf("GetCar() = %v, want %v", player.GetCar(), tt.car)
			}
			if player.GetDiscardPile() != tt.discardPile {
				t.Errorf("GetDiscardPile() = %v, want %v", player.GetDiscardPile(), tt.discardPile)
			}
			if player.GetDeck() != tt.deck {
				t.Errorf("GetDeck() = %v, want %v", player.GetDeck(), tt.deck)
			}
			if player.GetHand() != tt.hand {
				t.Errorf("GetHand() = %v, want %v", player.GetHand(), tt.hand)
			}

			// Test that icons are initialized as empty map
			icons := player.GetIcons()
			if len(icons) != 0 {
				t.Errorf("Initial icons should be empty, got %v", icons)
			}
		})
	}
}

func TestPlayer_GetName(t *testing.T) {
	tests := []struct {
		name       string
		playerName string
		expected   string
	}{
		{
			name:       "Simple name",
			playerName: "Alice",
			expected:   "Alice",
		},
		{
			name:       "Empty name",
			playerName: "",
			expected:   "",
		},
		{
			name:       "Long name",
			playerName: "VeryLongPlayerNameWithSpaces",
			expected:   "VeryLongPlayerNameWithSpaces",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player := NewPlayer(tt.playerName, NewCar("red", 3), NewDiscardPile(), NewDeck([]Card{}), NewHand())

			if player.GetName() != tt.expected {
				t.Errorf("GetName() = %s, want %s", player.GetName(), tt.expected)
			}
		})
	}
}

func TestPlayer_GetCar(t *testing.T) {
	car := NewCar("blue", 4)
	player := NewPlayer("TestPlayer", car, NewDiscardPile(), NewDeck([]Card{}), NewHand())

	if player.GetCar() != car {
		t.Errorf("GetCar() = %v, want %v", player.GetCar(), car)
	}
}

func TestPlayer_GetDiscardPile(t *testing.T) {
	discardPile := NewDiscardPile()
	player := NewPlayer("TestPlayer", NewCar("red", 3), discardPile, NewDeck([]Card{}), NewHand())

	if player.GetDiscardPile() != discardPile {
		t.Errorf("GetDiscardPile() = %v, want %v", player.GetDiscardPile(), discardPile)
	}
}

func TestPlayer_GetDeck(t *testing.T) {
	deck := NewDeck(createPlayerTestCards())
	player := NewPlayer("TestPlayer", NewCar("red", 3), NewDiscardPile(), deck, NewHand())

	if player.GetDeck() != deck {
		t.Errorf("GetDeck() = %v, want %v", player.GetDeck(), deck)
	}
}

func TestPlayer_GetHand(t *testing.T) {
	hand := NewHand()
	player := NewPlayer("TestPlayer", NewCar("red", 3), NewDiscardPile(), NewDeck([]Card{}), hand)

	if player.GetHand() != hand {
		t.Errorf("GetHand() = %v, want %v", player.GetHand(), hand)
	}
}

func TestPlayer_DrawCard(t *testing.T) {
	tests := []struct {
		name      string
		deckCards []Card
		draws     int
	}{
		{
			name:      "Draw from empty deck",
			deckCards: []Card{},
			draws:     1,
		},
		{
			name:      "Draw single card",
			deckCards: []Card{NewCard("Test", 1, map[Icon]int{}, true, true, false)},
			draws:     1,
		},
		{
			name:      "Draw multiple cards",
			deckCards: createPlayerTestCards(),
			draws:     3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player := NewPlayer("TestPlayer", NewCar("red", 3), NewDiscardPile(), NewDeck(tt.deckCards), NewHand())

			// Draw cards
			for i := 0; i < tt.draws; i++ {
				player.DrawCard(player.GetDeck())
			}

			// Verify cards were added to hand by checking hand's internal state
			hand := player.GetHand().(*hand)
			drawnCount := len(hand.cards)

			// Should have drawn at most the number of cards in the deck
			expectedDraws := tt.draws
			if tt.draws > len(tt.deckCards) {
				expectedDraws = len(tt.deckCards)
			}

			if drawnCount != expectedDraws {
				t.Errorf("Drew %d cards, expected %d", drawnCount, expectedDraws)
			}
		})
	}
}

func TestPlayer_DiscardCard(t *testing.T) {
	tests := []struct {
		name          string
		handCards     []Card
		discardIndex  int
		expectedError string
	}{
		{
			name:          "Discard valid card",
			handCards:     []Card{NewCard("Discardable", 1, map[Icon]int{}, true, true, false)},
			discardIndex:  0,
			expectedError: "",
		},
		{
			name:          "Discard non-discardable card",
			handCards:     []Card{NewCard("Non-discardable", 1, map[Icon]int{}, false, true, false)},
			discardIndex:  0,
			expectedError: "card is not discardable",
		},
		{
			name:          "Discard with invalid index",
			handCards:     []Card{NewCard("Test", 1, map[Icon]int{}, true, true, false)},
			discardIndex:  1,
			expectedError: "invalid card index",
		},
		{
			name:          "Discard from empty hand",
			handCards:     []Card{},
			discardIndex:  0,
			expectedError: "invalid card index",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hand := NewHand()
			hand.AddCards(tt.handCards)
			player := NewPlayer("TestPlayer", NewCar("red", 3), NewDiscardPile(), NewDeck([]Card{}), hand)

			err := player.DiscardCard(tt.discardIndex)

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
		})
	}
}

func TestPlayer_GetIcons(t *testing.T) {
	player := NewPlayer("TestPlayer", NewCar("red", 3), NewDiscardPile(), NewDeck([]Card{}), NewHand())

	// Initially should be empty
	icons := player.GetIcons()
	if len(icons) != 0 {
		t.Errorf("Initial icons should be empty, got %v", icons)
	}

	// Add some icons
	testIcons := map[Icon]int{IconBoost: 2, IconCooling: 1}
	player.AddIcons(testIcons)

	// Should now have the added icons
	icons = player.GetIcons()
	if icons[IconBoost] != 2 {
		t.Errorf("Boost icons = %d, want 2", icons[IconBoost])
	}
	if icons[IconCooling] != 1 {
		t.Errorf("Cooling icons = %d, want 1", icons[IconCooling])
	}
}

func TestPlayer_AddIcons(t *testing.T) {
	tests := []struct {
		name     string
		initial  map[Icon]int
		toAdd    map[Icon]int
		expected map[Icon]int
	}{
		{
			name:     "Add to empty icons",
			initial:  map[Icon]int{},
			toAdd:    map[Icon]int{IconBoost: 2, IconCooling: 1},
			expected: map[Icon]int{IconBoost: 2, IconCooling: 1},
		},
		{
			name:     "Add to existing icons",
			initial:  map[Icon]int{IconBoost: 1},
			toAdd:    map[Icon]int{IconBoost: 2, IconCooling: 1},
			expected: map[Icon]int{IconBoost: 3, IconCooling: 1},
		},
		{
			name:     "Add empty icons",
			initial:  map[Icon]int{IconBoost: 1},
			toAdd:    map[Icon]int{},
			expected: map[Icon]int{IconBoost: 1},
		},
		{
			name:     "Add nil icons",
			initial:  map[Icon]int{IconBoost: 1},
			toAdd:    nil,
			expected: map[Icon]int{IconBoost: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player := NewPlayer("TestPlayer", NewCar("red", 3), NewDiscardPile(), NewDeck([]Card{}), NewHand())

			// Set initial icons
			player.AddIcons(tt.initial)

			// Add new icons
			player.AddIcons(tt.toAdd)

			// Check result
			icons := player.GetIcons()
			for icon, expectedCount := range tt.expected {
				if icons[icon] != expectedCount {
					t.Errorf("Icon %v = %d, want %d", icon, icons[icon], expectedCount)
				}
			}
		})
	}
}

func TestPlayer_PlayCard(t *testing.T) {
	tests := []struct {
		name          string
		handCards     []Card
		playIndex     int
		expectedError string
	}{
		{
			name:          "Play valid card",
			handCards:     []Card{NewCard("Playable", 1, map[Icon]int{}, true, true, false)},
			playIndex:     0,
			expectedError: "",
		},
		{
			name:          "Play non-playable card",
			handCards:     []Card{NewCard("Non-playable", 1, map[Icon]int{}, true, false, false)},
			playIndex:     0,
			expectedError: "card is not playable",
		},
		{
			name:          "Play with invalid index",
			handCards:     []Card{NewCard("Test", 1, map[Icon]int{}, true, true, false)},
			playIndex:     1,
			expectedError: "invalid card index",
		},
		{
			name:          "Play from empty hand",
			handCards:     []Card{},
			playIndex:     0,
			expectedError: "invalid card index",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hand := NewHand()
			hand.AddCards(tt.handCards)
			player := NewPlayer("TestPlayer", NewCar("red", 3), NewDiscardPile(), NewDeck([]Card{}), hand)

			err := player.PlayCard(tt.playIndex)

			if tt.expectedError == "" {
				if err != nil {
					t.Errorf("PlayCard() returned unexpected error: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("PlayCard() expected error '%s', got nil", tt.expectedError)
				} else if err.Error() != tt.expectedError {
					t.Errorf("PlayCard() expected error '%s', got '%s'", tt.expectedError, err.Error())
				}
			}
		})
	}
}

func TestPlayer_ResolvePlayedCards(t *testing.T) {
	tests := []struct {
		name          string
		playedCards   []Card
		expectedSpeed int
		expectedIcons map[Icon]int
	}{
		{
			name: "Resolve basic cards",
			playedCards: []Card{
				NewCard("Speed 1", 1, map[Icon]int{IconBoost: 1}, true, true, false),
				NewCard("Speed 2", 2, map[Icon]int{IconCooling: 1}, true, true, false),
			},
			expectedSpeed: 3,
			expectedIcons: map[Icon]int{IconBoost: 1, IconCooling: 1},
		},
		{
			name: "Resolve cards with complex icons",
			playedCards: []Card{
				NewCard("Complex", 3, map[Icon]int{IconBoost: 2, IconCooling: 1}, true, true, false),
			},
			expectedSpeed: 3,
			expectedIcons: map[Icon]int{IconBoost: 2, IconCooling: 1},
		},
		{
			name:          "Resolve no cards",
			playedCards:   []Card{},
			expectedSpeed: 0,
			expectedIcons: map[Icon]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player := NewPlayer("TestPlayer", NewCar("red", 3), NewDiscardPile(), NewDeck([]Card{}), NewHand())

			// Add cards to hand and play them one by one
			hand := NewHand()
			hand.AddCards(tt.playedCards)
			player = NewPlayer("TestPlayer", NewCar("red", 3), NewDiscardPile(), NewDeck([]Card{}), hand)

			// Play all cards
			for i := 0; i < len(tt.playedCards); i++ {
				err := player.PlayCard(0) // Always play from index 0 since cards are removed after playing
				if err != nil {
					t.Errorf("Failed to play card %d: %v", i, err)
				}
			}

			// Resolve played cards
			player.ResolvePlayedCards()

			// Check car speed
			if player.GetCar().GetSpeed() != tt.expectedSpeed {
				t.Errorf("Car speed = %d, want %d", player.GetCar().GetSpeed(), tt.expectedSpeed)
			}

			// Check icons
			icons := player.GetIcons()
			for icon, expectedCount := range tt.expectedIcons {
				if icons[icon] != expectedCount {
					t.Errorf("Icon %v = %d, want %d", icon, icons[icon], expectedCount)
				}
			}
		})
	}
}

func TestPlayer_ResolvePlayedCards_StressCard(t *testing.T) {
	// Create a deck with basic cards for stress resolution
	basicCards := []Card{
		NewCard("Non-basic", 0, map[Icon]int{}, false, true, false),
		NewCard("Basic", 2, map[Icon]int{}, true, true, true),
	}
	deck := NewDeck(basicCards)
	discardPile := NewDiscardPile()

	// Create stress card
	stressCard := NewCard("Stress", 0, map[Icon]int{}, true, true, false)

	// Create player and play stress card
	hand := NewHand()
	hand.AddCards([]Card{stressCard})
	player := NewPlayer("TestPlayer", NewCar("red", 3), discardPile, deck, hand)

	// Play the stress card
	player.PlayCard(0)

	// Resolve played cards
	player.ResolvePlayedCards()

	// The stress card should have been resolved by drawing until a basic card was found
	// The car speed should be the speed of the basic card (2)
	if player.GetCar().GetSpeed() != 2 {
		t.Errorf("Car speed after stress card = %d, want 2", player.GetCar().GetSpeed())
	}
}

func TestPlayer_InterfaceCompliance(t *testing.T) {
	player := NewPlayer("TestPlayer", NewCar("red", 3), NewDiscardPile(), NewDeck([]Card{}), NewHand())

	// Test that player implements the Player interface
	var _ Player = player

	// Test that we can call all interface methods
	deck := NewDeck([]Card{})

	// These should not panic
	_ = player.GetName()
	_ = player.GetCar()
	_ = player.GetDiscardPile()
	_ = player.GetDeck()
	_ = player.GetHand()
	player.DrawCard(deck)
	player.DiscardCard(0)
	_ = player.GetIcons()
	player.AddIcons(map[Icon]int{IconBoost: 1})
	player.PlayCard(0)
	player.ResolvePlayedCards()
}

func TestPlayer_EdgeCases(t *testing.T) {
	t.Run("DrawCard with nil deck", func(t *testing.T) {
		player := NewPlayer("TestPlayer", NewCar("red", 3), NewDiscardPile(), NewDeck([]Card{}), NewHand())

		// This should not panic
		player.DrawCard(nil)
	})

	t.Run("AddIcons with nil map", func(t *testing.T) {
		player := NewPlayer("TestPlayer", NewCar("red", 3), NewDiscardPile(), NewDeck([]Card{}), NewHand())

		// This should not panic
		player.AddIcons(nil)

		// Icons should remain unchanged
		icons := player.GetIcons()
		if len(icons) != 0 {
			t.Error("Icons should remain empty after adding nil")
		}
	})

	t.Run("Multiple icon additions", func(t *testing.T) {
		player := NewPlayer("TestPlayer", NewCar("red", 3), NewDiscardPile(), NewDeck([]Card{}), NewHand())

		// Add icons multiple times
		player.AddIcons(map[Icon]int{IconBoost: 1})
		player.AddIcons(map[Icon]int{IconBoost: 2})
		player.AddIcons(map[Icon]int{IconCooling: 1})

		// Check final state
		icons := player.GetIcons()
		if icons[IconBoost] != 3 {
			t.Errorf("Boost icons = %d, want 3", icons[IconBoost])
		}
		if icons[IconCooling] != 1 {
			t.Errorf("Cooling icons = %d, want 1", icons[IconCooling])
		}
	})
}

// Benchmark tests
func BenchmarkNewPlayer(b *testing.B) {
	car := NewCar("red", 3)
	discardPile := NewDiscardPile()
	deck := NewDeck([]Card{})
	hand := NewHand()

	for i := 0; i < b.N; i++ {
		NewPlayer("BenchmarkPlayer", car, discardPile, deck, hand)
	}
}

func BenchmarkPlayer_GetName(b *testing.B) {
	player := NewPlayer("BenchmarkPlayer", NewCar("red", 3), NewDiscardPile(), NewDeck([]Card{}), NewHand())

	for i := 0; i < b.N; i++ {
		player.GetName()
	}
}

func BenchmarkPlayer_DrawCard(b *testing.B) {
	player := NewPlayer("BenchmarkPlayer", NewCar("red", 3), NewDiscardPile(), NewDeck(createPlayerTestCards()), NewHand())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		player.DrawCard(player.GetDeck())
	}
}

func BenchmarkPlayer_AddIcons(b *testing.B) {
	player := NewPlayer("BenchmarkPlayer", NewCar("red", 3), NewDiscardPile(), NewDeck([]Card{}), NewHand())
	icons := map[Icon]int{IconBoost: 2, IconCooling: 1}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		player.AddIcons(icons)
	}
}

func BenchmarkPlayer_PlayCard(b *testing.B) {
	hand := NewHand()
	testCard := NewCard("Benchmark", 1, map[Icon]int{}, true, true, false)
	hand.AddCards([]Card{testCard})
	player := NewPlayer("BenchmarkPlayer", NewCar("red", 3), NewDiscardPile(), NewDeck([]Card{}), hand)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Reset hand for each iteration
		hand := NewHand()
		hand.AddCards([]Card{testCard})
		player = NewPlayer("BenchmarkPlayer", NewCar("red", 3), NewDiscardPile(), NewDeck([]Card{}), hand)
		player.PlayCard(0)
	}
}
