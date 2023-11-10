package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dbURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	//Migrate DB

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Ticket{})

	return db, nil
}
