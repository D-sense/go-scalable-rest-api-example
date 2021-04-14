package objects

import (
	"context"
	"github.com/d-sense/go-scalable-rest-api-example/password"
	"time"
)

type CustomerService interface {
	Create(ctx context.Context, data Customer) error
	Customers(ctx context.Context) ([]*Customer, error)
	Customer(ctx context.Context, id string) (*Customer, error)
	FindCustomerByEmail(ctx context.Context, email string) (*Customer, error)
}

type Customer struct {
	ID        string         `gorm:"id"        json:"id,omitempty"`
	Counter   int            `gorm:"counter"   json:"counter"`
	FullName  string         `gorm:"full_name" json:"full_name,omitempty"`
	Email     string         `gorm:"email,unique;unique_index;not null" json:"email"`
	Password  *password.Hash `gorm:"password" json:"password,omitempty" gorm:"type: jsonb not null default '{}'::jsonb column:password"`
	Lends     []Book         `json:"books"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt *time.Time     `json:"deleted_at,omitempty"`
}
