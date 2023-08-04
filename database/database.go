package database

import (
	"log"
	"time"

	"bitbucket.org/adson14/api_golang_ii.git/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func StartDb() {
	str := "host=localhost user=postgres password=123456 dbname=api_go port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})

	if err != nil {
		log.Fatal("error: ", err)
	}

	db = database

	config, _ := db.DB()

	config.SetConnMaxIdleTime(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}
