package service

import (
	"context"

	"go/src/github.om/msegeya56/library_platform/book/pkg/repository"

	"github.com/msegeya56/ecommerce.go.module/pkg/domains/entities"
	"github.com/msegeya56/ecommerce.go.module/pkg/domains/models"
)

// BookService describes the service.
type BookService interface {
	Post(ctx context.Context, data *models.Book) (replyData *entities.Book, replyError error)
	GetAll(ctx context.Context) (replyData []*entities.Book, replyError error)
	GetByID(ctx context.Context, ID uint) (replyData *entities.Book, replyError error)
	Update(ctx context.Context, ID uint, data *models.Book) (replyData *entities.Book, replyError error)
	GetByDate(ctx context.Context, param string) (replyData []*entities.Book, replyError error)
	Delete(ctx context.Context, ID uint) (replyError error)
}

type basicBookService struct{
	repo repository.BookRepository
}

func (b *basicBookService) Post(ctx context.Context, data *models.Book) (replyData *entities.Book, replyError error) {
	// TODO implement the business logic of Post
	replyData, replyError = b.repo.Post(data)
	return replyData, replyError
}

func (b *basicBookService) GetAll(ctx context.Context) (replyData []*entities.Book, replyError error) {
	// TODO implement the business logic of GetAll
	replyData, replyError = b.repo.GetAll()

	return replyData, replyError
}

func (b *basicBookService) GetByID(ctx context.Context, param uint) (replyData *entities.Book, replyError error) {
	// TODO implement the business logic of GetByID
	replyData, replyError = b.repo.GetByID(param)
	return replyData, replyError
}

func (b *basicBookService) Update(ctx context.Context, param uint, data *models.Book) (replyData *entities.Book, replyError error) {
	// TODO implement the business logic of Update
	replyData, replyError = b.repo.Update(param, data)
	return replyData, replyError
}

func (b *basicBookService) GetByDate(ctx context.Context, dataParam string) (replyData []*entities.Book, replyError error) {
	// TODO implement the business logic of GetByDate
	replyData, replyError = b.repo.GetByDate(dataParam)
	return replyData, replyError
}

func (b *basicBookService) Delete(ctx context.Context, param uint) (replyError error) {
	// TODO implement the business logic of Delete
	replyError = b.repo.Delete(param)
	return replyError
}

// NewBasicBookService returns a naive, stateless implementation of BookService.
func NewBasicBookService(repo repository.BookRepository) BookService {
	return &basicBookService{
		repo: repo,
	}
}

// New returns a BookService with all of the expected middleware wired in.
func New(repo repository.BookRepository,middleware []Middleware) BookService {
	var svc BookService = NewBasicBookService(repo)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
