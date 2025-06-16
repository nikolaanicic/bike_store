package models

import "time"

type City string

const (
	NOVI_SAD   BikeType = "NOVISAD"
	KRAGUJEVAC BikeType = "KRAGUJEVAC"
	SUBOTICA   BikeType = "SUBOTICA"
)

type Rental struct {
	ID              int       `gorm:"column:ID;primaryKey;autoIncrement"`
	RentalDate      time.Time `gorm:"column:RENTAL_DATE;default:CURRENT_TIMESTAMP"`
	RentalCitizenID string    `gorm:"column:RENTAL_CITIZEN_ID;size:13"`
	RentalBikeID    int       `gorm:"column:RENTAL_BIKE_ID"`
	Bike            Bike      `gorm:"foreignKey:RentalBikeID;references:ID;OnDelete:CASCADE"`
	City            string    `gorm:"column:CITY;size:10"`
	ReturnedDate    time.Time `gorm:"column:RETURNED_DATE;default:NULL"`
	Completed       bool      `gorm:"column:COMPLETED;default:false"`
}

func (Rental) TableName() string {
	return "Rentals"
}
