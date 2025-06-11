package repositories

import (
	"bike_store/database"
	"bike_store/database/models"

	"gorm.io/gorm"
)

type RentalRepositories struct {
	database.BaseRepository[int, models.Rental]
}

func NewRentalRepository(db *gorm.DB) *RentalRepositories {
	return &RentalRepositories{database.BaseRepository[int, models.Rental]{DB: db}}
}
