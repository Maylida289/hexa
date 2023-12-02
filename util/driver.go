//driver ini gunanya untuk mengirimkan koneksi database
package util

import (
	"hexa/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm/dialects/postgres"
)

type DatabaseDriver string

const (
	Postgres DatabaseDriver = "postgres"
)

type DatabaseConnection struct { // sebagai identifier teknologi yang kita pakai
	Driver DatabaseDriver

	Postgres *gorm.DB
}

func NewConnectionDatabase(config *config.AppConfig) *DatabaseConnection {
	var db DatabaseConnection
	db.Driver = Postgres
	db.Postgres = newPostgres(config)

	return &db
}

func newPostgres(config *config.AppConfig) *gorm.DB {
	dbURL := config.Database.DBURL
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func (db *DatabaseConnection) CloseConnection() {
	if db.Postgres != nil {
		db, _ := db.Postgres.DB()
		db.Close()
	}
}
