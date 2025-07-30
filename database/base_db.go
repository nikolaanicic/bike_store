package database

import (
	"bike_store/configuration"
	"bike_store/log"
	"time"

	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v3"
	"github.com/golang-migrate/migrate/v3/database/mysql"
	_ "github.com/golang-migrate/migrate/v3/source/file"
	gormmysql "gorm.io/driver/mysql"
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

func createDatabaseIfNotExists(user, password, host, dbName string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", user, password, host) // no DB specified
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	return err
}

func (db *BaseDatabase) runMigrations(user, password, host, dbName, migrationsPath string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbName)
	ddb, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer ddb.Close()

	driver, err := mysql.WithInstance(ddb, &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		dbName, driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func (db *BaseDatabase) Configure(dbConfig *configuration.Database) error {
	var err error

	if err := createDatabaseIfNotExists(dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.DbName); err != nil {
		log.Fatalf("failed to create the database: %v", err)
		return err
	}

	if err := db.runMigrations(dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.DbName, dbConfig.MigrationsPath); err != nil {
		log.Fatalf("failed to create the database: %v", err)
		return err
	}

	for i := 0; i < 5; i++ {
		db.DB, err = gorm.Open(gormmysql.Open(getConnString(dbConfig)), &gorm.Config{})
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
