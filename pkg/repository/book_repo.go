package repository

import (
	"errors"
	"log"
	

	"github.com/msegeya56/ecommerce.go.module/pkg/domains/entities"
	"github.com/msegeya56/ecommerce.go.module/pkg/domains/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	Post(data *models.Book) (replyData *entities.Book, replyError error)
	GetAll() (replyData []*entities.Book, replyError error)
	GetByID(param uint) (replyData *entities.Book, replyError error)
	Update(param uint, data *models.Book) (replyData *entities.Book, replyError error)
	GetByDate(dataParam string) (replyData []*entities.Book, replyError error)
	Delete(param uint) (replyError error)
}

type BookImpl struct {
	DB *gorm.DB
}

func NewBookRepository(Db *gorm.DB) *BookImpl {
	return &BookImpl{
		DB: Db,
	}
}

func (r *BookImpl) Post(data *models.Book) (replyData *entities.Book, replyError error) {
	// Validate each field before creating the Book

	Book := &entities.Book{
		ID:          data.ID,
		Title:       data.Title,
		Author:      data.Author,
		ISBN:        data.ISBN,
		PublishYear: data.PublishYear,
		Genre:       data.Genre,
	}

	result := r.DB.Create(&Book)
	if result.Error != nil {
		return nil, result.Error
	}

	replyData = Book
	return replyData, replyError
}

func (r *BookImpl) GetAll() (replyData []*entities.Book, replyError error) {
	var Books []entities.Book
	result := r.DB.Find(&Books)
	if result.Error != nil {
		replyError = result.Error
		return nil, replyError
	}

	// Convert entities.Book to []*entities.Book
	for i := range Books {
		replyData = append(replyData, &Books[i])
	}

	return replyData, nil
}
func (r *BookImpl) GetByID(param uint) (replyData *entities.Book, replyError error) {

	log.Println("Param value:", param)

	// Check if the provided ID is valid
	
	if param == 0 {
		return nil, errors.New("invalid ID: ID must be greater than 0")
	}

	var Book entities.Book
	result := r.DB.First(&Book, param)
	if result.Error != nil {
		return nil, result.Error
	}

	replyData = &Book
	return replyData, nil
}

func (r *BookImpl) Update(param uint, data *models.Book) (replyData *entities.Book, replyError error) {
	var Book entities.Book
	result := r.DB.First(&Book, param)
	if result.Error != nil {
		replyError = result.Error
		return nil, replyError
	}

	// Update Book fields based on data

	result = r.DB.Save(&Book)
	if result.Error != nil {
		replyError = result.Error
		return nil, replyError
	}

	replyData = &Book
	return replyData, nil
}

func (r *BookImpl) GetByDate(dataParam string) (replyData []*entities.Book, replyError error) {
	var Books []entities.Book
	result := r.DB.Where("date = ?", dataParam).Find(&Books)
	if result.Error != nil {
		replyError = result.Error
		return nil, replyError
	}

	// Convert entities.Book to []*entities.Book
	for i := range Books {
		replyData = append(replyData, &Books[i])
	}

	return replyData, nil
}

func (r *BookImpl) Delete(param uint) (replyError error) {
	result := r.DB.Delete(&entities.Book{}, param)
	if result.Error != nil {
		replyError = result.Error
		return
	}

	return replyError
}
