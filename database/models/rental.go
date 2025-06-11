package models

import "time"

type Rental struct {
	ID           int       `gorm:"column:ID;primaryKey;autoIncrement"`
	RentalDate   time.Time `gorm:"column:RENTAL_DATE"`
	RentalUserID string
	RentalBikeID int
	User         User `gorm:"foreignKey:RentalUserID;references:ID;OnDelete:CASCADE"`
	Bike         Bike `gorm:"foreignKey:RentalBikeID;references:ID;OnDelete:CASCADE"`
}
