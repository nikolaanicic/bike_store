package data

import (
	"bike_store/city_store/data/repositories"
	"bike_store/configuration"
	"bike_store/database"
)

type Database struct {
	*database.BaseDatabase

	Bikes   *repositories.BikeRepository
	Rentals *repositories.RentalRepositories
}

func NewDB() database.IDatabase {
	return &Database{BaseDatabase: database.NewBaseDB()}
}

func (db *Database) Configure(config *configuration.Database) error {
	if err := db.BaseDatabase.Configure(config); err != nil {
		return err
	}

	db.Bikes = repositories.NewBikeRepository(db.BaseDatabase.DB)
	db.Rentals = repositories.NewRentalRepository(db.BaseDatabase.DB)

	return nil
}
