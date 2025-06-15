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

func (r *RentalRepositories) GetByUserAndBike(citizenID string, bikeID int) (*models.Rental, error) {
	var rental models.Rental
	err := r.DB.Where(&models.Rental{RentalUserID: citizenID, RentalBikeID: bikeID, Completed: false}).First(&rental).Error

	return &rental, err
}
