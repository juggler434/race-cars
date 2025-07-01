package repository

import (
	"database/sql"
	"fmt"
	"race-cars/internal/config"
	"race-cars/internal/models"
	"time"
)

// CarRepository handles database operations for cars
type GameRepository struct {
	db *sql.DB
}

// NewGameRepository creates a new game repository
func GameRepository() *CarRepository {
	return &CarRepository{
		db: config.DB,
	}
}
