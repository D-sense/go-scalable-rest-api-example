package book

import (
	"fmt"
	"github.com/d-sense/go-scalable-rest-api-example/objects"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Handler struct {
	bookService objects.BookService
	// add more services, such as Email Delivery service, Session Service, Third-parties services...as required.
}

// NewHandler is the Book handler constructor
func NewHandler(
	bookService objects.BookService,
) *Handler {
	h := &Handler{
		bookService: bookService,
	}
	return h
}

func (h *Handler) Create(ctx *gin.Context, input *objects.BookVM) error {

	book := objects.Book{
		Title:    input.Title,
		Category: input.Category,
		AuthorID: input.AuthorID,
	}

	err := h.bookService.CreateBook(ctx, book)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error creating a book"))
	}

	return nil
}

func (h *Handler) Book(ctx *gin.Context, id string) (*objects.Book, error) {
	ct, err := h.bookService.Book(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching an author"))
	}

	return ct, nil
}

func (h *Handler) Books(ctx *gin.Context) ([]*objects.Book, error) {
	bk, err := h.bookService.Books(ctx)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching all authors"))
	}

	return bk, nil
}
