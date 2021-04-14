package postgres

import (
	"context"
	"github.com/d-sense/go-scalable-rest-api-example/objects"
	"time"
)

type BookService struct {
	client *Client
}

// NewBookService is an Book service constructor
func NewBookService(ctx context.Context, client *Client) *BookService {
	bs := &BookService{client}
	return bs
}

func (bs BookService) CreateBook(ctx context.Context, book objects.Book) error {
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	adminDetail := objects.Book{
		Title:    book.Title,
		Category: book.Category,
		AuthorID: book.AuthorID,
	}

	return bs.client.db.Create(&adminDetail).Error
}

func (bs BookService) Books(ctx context.Context) ([]*objects.Book, error) {
	var books []*objects.Book
	return books, bs.client.db.Order("created_at desc").Find(&books).Error
}

func (bs BookService) Book(ctx context.Context, id string) (*objects.Book, error) {
	var team objects.Book
	return &team, bs.client.db.Set("gorm:auto_preload", true).Where("id = ?", id).Find(&team).Error
}

func (bs BookService) Update(ctx context.Context, book *objects.Book) error {
	book.UpdatedAt = time.Now()
	return bs.client.db.Model(&book).Where("id = ?", book.ID).
		Updates(map[string]interface{}{"name": book.Title, "category": book.Category}).Error
}

func (bs BookService) Delete(ctx context.Context, id string) error {
	var book objects.Book
	return bs.client.db.Where("id = ?", id).Delete(&book).Error
}
