package connection

import (
	"errors"

	"github.com/msegeya56/ecommerce.go.module/pkg/domains/models"
	"gorm.io/gorm"
)




func SyncDB(DB *gorm.DB) error {
	if DB == nil {
		return errors.New("gormdb is nil")
	}

	// Add migration logic here
	err := DB.AutoMigrate(&models.Book{})
	if err != nil {
		return err
	}

	// Add other model migrations here

	return nil
}

	// Add any other models and their associations here for automatic table creation






