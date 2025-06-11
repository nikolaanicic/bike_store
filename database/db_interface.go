package database

import "bike_store/configuration"

type IDatabase interface {
	Configure(configuration.Database) error
	Connect() error
}
