package customer

import (
	"fmt"
	"github.com/d-sense/go-scalable-rest-api-example/objects"
	"github.com/d-sense/go-scalable-rest-api-example/password"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Handler struct {
	customerService objects.CustomerService
	bookService     objects.BookService
	// add more services as you need, such as Email Delivery service, Session Service, Third-parties services...etc.
}

// NewHandler is the Customer handler constructor
func NewHandler(
	customerService objects.CustomerService,
	bookService objects.BookService,
) *Handler {
	h := &Handler{
		customerService: customerService,
		bookService:     bookService,
	}
	return h
}

func (h *Handler) SignUp(ctx *gin.Context, input *objects.RegistrationVM) (*objects.Customer, error) {
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

	cust := objects.Customer{
		FullName: input.FullName,
		Email:    input.Email,
		Password: pHashed,
	}

	err = h.customerService.Create(ctx, cust)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error creating an author"))
	}

	return &cust, nil
}

func (h *Handler) AuthenticateUser(ctx *gin.Context, input *objects.LoginVM) (*objects.Customer, error) {
	// input validation should be handled first.
	// Handle it yourself
	//

	ct, err := h.customerService.FindCustomerByEmail(ctx, input.Email)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching user by email"))
	}

	if !ct.Password.IsEqualTo(input.Password) {
		return nil, errors.Wrap(err, fmt.Sprintf("error confirming user passowrd"))
	}

	ct.Password = nil
	return ct, nil
}

func (h *Handler) Customers(ctx *gin.Context) ([]*objects.Customer, error) {
	ct, err := h.customerService.Customers(ctx)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching all authors"))
	}

	return ct, nil
}

func (h *Handler) Customer(ctx *gin.Context, id string) (*objects.Customer, error) {
	ct, err := h.customerService.Customer(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching an author"))
	}

	return ct, nil
}
