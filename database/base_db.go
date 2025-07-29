package database

import (
	"bike_store/configuration"
	"bike_store/log"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BaseDatabase struct {
	DB *gorm.DB
}

func NewBaseDB() *BaseDatabase {
	return &BaseDatabase{}
}

func getConnString(dbConfig *configuration.Database) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
}

func (db *BaseDatabase) Configure(dbConfig *configuration.Database) error {
	var err error

	for i := 0; i < 5; i++ {
		db.DB, err = gorm.Open(mysql.Open(getConnString(dbConfig)), &gorm.Config{})
		if err == nil {
			log.Info("connected to the database")
			return nil
		}
		log.Info("failed to initialize database connection (attempt %d): %v", i+1, err)
		time.Sleep(1 * time.Second)
	}

	log.Fatalf("failed to initialize database connection after 5 attempts: %v", err)
	return err
}
