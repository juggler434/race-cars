package repository

import (
	"database/sql"
	"fmt"
	"race-cars/internal/config"
	"race-cars/internal/models"
	"time"
)

// CarRepository handles database operations for cars
type CarRepository struct {
	db *sql.DB
}

// NewCarRepository creates a new car repository
func NewCarRepository() *CarRepository {
	return &CarRepository{
		db: config.DB,
	}
}
