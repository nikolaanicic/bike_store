package models

import "time"

type City string

const (
	NOVI_SAD   BikeType = "NOVISAD"
	KRAGUJEVAC BikeType = "KRAGUJEVAC"
	SUBOTICA   BikeType = "SUBOTICA"
)

type Rental struct {
	ID           int       `gorm:"column:ID;primaryKey;autoIncrement"`
	RentalDate   time.Time `gorm:"column:RENTAL_DATE"`
	RentalUserID string    `gorm:"column:RENTAL_USER_ID;size:13"`
	RentalBikeID int       `gorm:"column:RENTAL_BIKE_ID"`
	Bike         Bike      `gorm:"foreignKey:RentalBikeID;references:ID;OnDelete:CASCADE"`
	City         string    `gorm:"column:CITY;size:10"`
	ReturnedData time.Time `gorm:"column:RETURNED_DATE"`
	Completed    bool      `gorm:"column:COMPLETED;default:false"`
}
