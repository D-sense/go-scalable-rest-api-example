package author

import (
	"fmt"
	"github.com/d-sense/go-scalable-rest-api-example/objects"
	"github.com/d-sense/go-scalable-rest-api-example/password"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Handler struct {
	authorService objects.AuthorService
	bookService   objects.BookService
	// add more services, such as Email Delivery service, Session Service, Third-parties services...as required.
}

// NewHandler is the Author handler constructor
func NewHandler(
	authorService objects.AuthorService,
	bookService objects.BookService,
) *Handler {
	h := &Handler{
		authorService: authorService,
		bookService:   bookService,
	}
	return h
}

func (h *Handler) SignUp(ctx *gin.Context, input *objects.RegistrationVM) (*objects.Author, error) {
	// input validation should be handled first.
	// Handle it yourself
	//

	pHashed, err := password.NewHashedPassword(input.Password)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error hasshng a password"))
	}

	// do some checking such as if email address already exists
	// Handle it yourself
	//

	author := &objects.Author{
		FullName: input.FullName,
		Email:    input.Email,
		Password: pHashed,
	}

	err = h.authorService.Create(ctx, author)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error creating an author"))
	}

	return author, nil
}

func (h *Handler) AuthenticateUser(ctx *gin.Context, input *objects.LoginVM) (*objects.Author, error) {
	// input validation should be handled first.
	// Handle it yourself
	//

	user, err := h.authorService.FindAuthorByEmail(ctx, input.Email)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching user by email"))
	}

	if !user.Password.IsEqualTo(input.Password) {
		return nil, errors.Wrap(err, fmt.Sprintf("error confirming user passowrd"))
	}

	return user, nil
}

func (h *Handler) Author(ctx *gin.Context, id string) (*objects.Author, error) {
	author, err := h.authorService.Author(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching an author"))
	}

	return author, nil
}

func (h *Handler) Authors(ctx *gin.Context) ([]*objects.Author, error) {
	authors, err := h.authorService.Authors(ctx)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching all authors"))
	}

	return authors, nil
}
