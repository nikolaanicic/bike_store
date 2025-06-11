package repositories

import (
	"bike_store/database"
	"bike_store/database/models"

	"gorm.io/gorm"
)

type BikeRepository struct {
	database.BaseRepository[int, models.Bike]
}

func NewBikeRepository(db *gorm.DB) *BikeRepository {
	return &BikeRepository{database.BaseRepository[int, models.Bike]{DB: db}}
}
