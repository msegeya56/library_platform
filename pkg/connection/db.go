package connection

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() (*gorm.DB, error) {
	if DB != nil {
		// Return the existing database connection if it's already created
		return DB, nil
	}

	// Connect to SQLite database
	var err error
	DB, err = gorm.Open(sqlite.Open("Books.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}
