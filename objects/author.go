package objects

import (
	"context"
	"github.com/d-sense/go-scalable-rest-api-example/password"
	"time"
)

type AuthorService interface {
	Create(ctx context.Context, data *Author) error
	Authors(ctx context.Context) ([]*Author, error)
	Author(ctx context.Context, id string) (*Author, error)
	FindAuthorByEmail(ctx context.Context, email string) (*Author, error)
}

type Author struct {
	ID        string         `gorm:"id"        json:"id,omitempty"`
	Counter   int            `gorm:"counter"   json:"counter"`
	FullName  string         `gorm:"full_name" json:"full_name,omitempty"`
	Email     string         `gorm:"email,unique;unique_index;not null" json:"email"`
	Password  *password.Hash `gorm:"password" json:"password,omitempty" gorm:"type: jsonb not null default '{}'::jsonb column:password"`
	Books     []Book         `json:"books"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt *time.Time     `json:"deleted_at,omitempty"`
}

type RegistrationVM struct {
	FullName string `json:"full_name,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginVM struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
