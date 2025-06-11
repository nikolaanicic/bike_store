package repositories

import (
	"bike_store/database"
	"bike_store/database/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	database.BaseRepository[string, models.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{database.BaseRepository[string, models.User]{DB: db}}
}
