package database

import (
	"bike_store/database/models"

	"gorm.io/gorm"
)

type Key interface {
	~int | ~string
}

type Model interface {
	models.User | models.Rental | models.Bike
}

type BaseRepository[K Key, T Model] struct {
	DB *gorm.DB
}

func NewBaseRepository[K Key, T Model](db *gorm.DB) BaseRepository[K, T] {
	return BaseRepository[K, T]{DB: db}
}

func (r *BaseRepository[K, T]) GetById(id K, model *T) error {
	return r.DB.First(model, id).Error
}

func (r *BaseRepository[K, T]) GetAll() ([]T, error) {
	var models []T

	if err := r.DB.Find(&models).Error; err != nil {
		return nil, err
	}

	return models, nil
}

func (r *BaseRepository[K, T]) Create(model *T) error {
	return r.DB.Create(model).Error
}

func (r *BaseRepository[K, T]) Update(model *T) error {
	return r.DB.Save(model).Error
}
