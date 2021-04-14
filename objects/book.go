package objects

import (
	"context"
	"time"
)

type BookService interface {
	CreateBook(ctx context.Context, book Book) error
	Books(ctx context.Context) ([]*Book, error)
	Book(ctx context.Context, id string) (*Book, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, data *Book) error
}

type Book struct {
	ID        string     `gorm:"id"        json:"id,omitempty"`
	Counter   int        `gorm:"counter"   json:"counter"`
	Title     string     `gorm:"title"     json:"title,omitempty"`
	Category  string     `gorm:"category"  json:"category,omitempty"`
	AuthorID  string     `gorm:"author_id" json:"author_id"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type BookVM struct {
	Title    string `gorm:"title"     json:"title,omitempty"`
	Category string `gorm:"category"  json:"category,omitempty"`
	AuthorID string `gorm:"author_id" json:"author_id"`
}
