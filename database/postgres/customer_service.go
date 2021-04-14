package postgres

import (
	"context"
	"github.com/d-sense/go-scalable-rest-api-example/objects"
	"time"
)

type CustomerService struct {
	client *Client
}

// NewCustomerService is a Customer service constructor
func NewCustomerService(ctx context.Context, client *Client) *CustomerService {
	cs := &CustomerService{client}
	return cs
}

func (cs CustomerService) Create(ctx context.Context, customer objects.Customer) error {
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()

	adminDetail := objects.Customer{
		FullName: customer.FullName,
	}

	return cs.client.db.Create(&adminDetail).Error
}

func (cs CustomerService) Customers(ctx context.Context) ([]*objects.Customer, error) {
	var customers []*objects.Customer
	return customers, cs.client.db.Order("created_at desc").Find(&customers).Error
}

func (cs CustomerService) Customer(ctx context.Context, id string) (*objects.Customer, error) {
	var customer objects.Customer
	return &customer, cs.client.db.Set("gorm:auto_preload", true).Where("id = ?", id).Find(&customer).Error
}

func (cs CustomerService) FindCustomerByEmail(ctx context.Context, email string) (*objects.Customer, error) {
	c := &objects.Customer{}
	return c, cs.client.db.Set("gorm:auto_preload", true).Find(c, "emil = ?", email).Error
}
