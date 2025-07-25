package data

import (
	"bike_store/central_store/data/repositories"
	"bike_store/configuration"
	"bike_store/database"
)

type Database struct {
	*database.BaseDatabase

	Users *repositories.UserRepository
}

func NewDB() database.IDatabase {
	return &Database{BaseDatabase: database.NewBaseDB()}
}

func (db *Database) Configure(config *configuration.Database) error {
	if err := db.BaseDatabase.Configure(config); err != nil {
		return err
	}

	db.Users = repositories.NewUserRepository(db.BaseDatabase.DB)

	return nil
}
