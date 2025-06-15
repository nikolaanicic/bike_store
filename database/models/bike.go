package models

type BikeType string

const (
	MountainBike BikeType = "MOUNTAIN"
	CityBike     BikeType = "CITY"
	RoadBike     BikeType = "ROAD"
)

type Bike struct {
	ID        int    `gorm:"column:ID;primaryKey;autoIncrement"`
	Type      string `gorm:"column:TYPE;size:10"`
	Available bool   `gorm:"column:AVAILABLE"`
	City      string `gorm:"column:CITY;size:10"`
}

func (Bike) TableName() string {
	return "Bikes"
}
