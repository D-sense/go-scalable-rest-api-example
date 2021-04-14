package json

import (
	"github.com/d-sense/go-scalable-rest-api-example/handlers/author"
	"github.com/d-sense/go-scalable-rest-api-example/objects"
	"github.com/d-sense/go-scalable-rest-api-example/responder"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthorJsonServer struct {
	author *author.Handler
}

// NewAuthorServer is an Author JSON server constructor
func NewAuthorServer(author *author.Handler) *AuthorJsonServer {
	a := &AuthorJsonServer{
		author: author,
	}

	return a
}

func (as AuthorJsonServer) SignUpAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var registrationVM *objects.RegistrationVM

		if err := c.ShouldBindJSON(&registrationVM); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := as.author.SignUp(c, registrationVM)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourceCreated, user)
	}
}

func (as AuthorJsonServer) LoginAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginVM *objects.LoginVM

		if err := c.ShouldBindJSON(&loginVM); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := as.author.AuthenticateUser(c, loginVM)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourceCreated, user)
	}
}

func (as AuthorJsonServer) GetAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorId := ""
		user, err := as.author.Author(c, authorId)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourcesFetched, user)
	}
}

func (as AuthorJsonServer) GetAuthors() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := as.author.Authors(c)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourcesFetched, users)
	}
}
