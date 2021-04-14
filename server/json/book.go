package json

import (
	"fmt"
	"github.com/d-sense/go-scalable-rest-api-example/handlers/book"
	"github.com/d-sense/go-scalable-rest-api-example/objects"
	"github.com/d-sense/go-scalable-rest-api-example/responder"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookJsonServer struct {
	book *book.Handler
}

// NewBookServer is a Book JSON server constructor
func NewBookServer(book *book.Handler) *BookJsonServer {
	b := &BookJsonServer{
		book: book,
	}

	return b
}

func (bs BookJsonServer) CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var bookVM *objects.BookVM

		if err := c.ShouldBindJSON(&bookVM); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := bs.book.Create(c, bookVM)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourceCreated, nil)
	}
}

func (bs BookJsonServer) GetBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		pId := c.Query("productId")

		p, err := bs.book.Book(c, pId)
		if err != nil {
			responder.JsonResponse(c, false, fmt.Sprintf("ID %v %s ", pId, err.Error()), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourceFetched, p)
	}
}

func (ps BookJsonServer) GetBooks() gin.HandlerFunc {
	return func(c *gin.Context) {

		products, err := ps.book.Books(c)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourcesFetched, products)
	}
}
