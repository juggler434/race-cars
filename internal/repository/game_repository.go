package repository

import (
	"database/sql"
	"race-cars/internal/config"
)

// CarRepository handles database operations for cars
type GameRepository struct {
	db *sql.DB
}

// NewGameRepository creates a new game repository
func NewGameRepository() *GameRepository {
	return &GameRepository{
		db: config.DB,
	}
}
