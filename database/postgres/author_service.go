package postgres

import (
	"context"
	"github.com/d-sense/go-scalable-rest-api-example/objects"
	"time"
)

type AuthorService struct {
	client *Client
}

// NewAuthorService is an Author service constructor
func NewAuthorService(ctx context.Context, client *Client) *AuthorService {
	as := &AuthorService{client}
	return as
}

func (as AuthorService) Create(ctx context.Context, input *objects.Author) error {
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	userDetail := objects.Author{
		FullName: input.FullName,
		Email:    input.Email,
		Password: input.Password,
	}

	return as.client.db.Create(&userDetail).Error
}

func (as AuthorService) Authors(ctx context.Context) ([]*objects.Author, error) {
	var authors []*objects.Author
	return authors, as.client.db.Order("created_at desc").Find(&authors).Error
}

func (as AuthorService) Author(ctx context.Context, id string) (*objects.Author, error) {
	var author objects.Author
	return &author, as.client.db.Set("gorm:auto_preload", true).Where("id = ?", id).Find(&author).Error
}

func (as AuthorService) Update(ctx context.Context, author *objects.Author) error {
	author.UpdatedAt = time.Now()
	return as.client.db.Model(&author).Where("id = ?", author.ID).
		Updates(map[string]interface{}{"full_name": author.FullName}).Error
}

func (as AuthorService) Delete(ctx context.Context, id string) error {
	var author objects.Author
	return as.client.db.Where("id = ?", id).Delete(&author).Error
}

func (as AuthorService) FindAuthorByEmail(ctx context.Context, email string) (*objects.Author, error) {
	p := &objects.Author{}
	return p, as.client.db.Set("gorm:auto_preload", true).Find(p, "emil = ?", email).Error
}
