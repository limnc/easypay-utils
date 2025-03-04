package dbconfig

import (
	"log"
	//Driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Initialize DB with pointers
var DB *gorm.DB

func InitDB(dsn string) (*gorm.DB, error) {
	//dsn = destination
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return DB, nil
}

// Disconnect function will accepting db as params
func Disconnect(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error disconnecting from database: ", err)
	}

	sqlDB.Close()
	log.Println("Disconnected from database")
}
