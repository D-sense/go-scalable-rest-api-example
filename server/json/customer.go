package json

import (
	"github.com/d-sense/go-scalable-rest-api-example/handlers/customer"
	"github.com/d-sense/go-scalable-rest-api-example/objects"
	"github.com/d-sense/go-scalable-rest-api-example/responder"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomerJsonServer struct {
	customer *customer.Handler
}

// NewCustomerServer is a Customer JSON server constructor
func NewCustomerServer(customer *customer.Handler) *CustomerJsonServer {
	c := &CustomerJsonServer{
		customer: customer,
	}

	return c
}

func (cs CustomerJsonServer) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var registrationVM *objects.RegistrationVM

		if err := c.ShouldBindJSON(&registrationVM); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := cs.customer.SignUp(c, registrationVM)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourceCreated, user)
	}
}

func (cs CustomerJsonServer) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginVM *objects.LoginVM

		if err := c.ShouldBindJSON(&loginVM); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := cs.customer.AuthenticateUser(c, loginVM)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourceCreated, user)
	}
}

func (cs CustomerJsonServer) GetCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {

		customerId := c.Query("customer_id")
		cust, err := cs.customer.Customer(c, customerId)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourcesFetched, cust)
	}
}

func (cs CustomerJsonServer) GetCustomers() gin.HandlerFunc {
	return func(c *gin.Context) {

		products, err := cs.customer.Customers(c)
		if err != nil {
			responder.JsonResponse(c, false, err.Error(), nil)
			return
		}

		responder.JsonResponse(c, true, responder.ResourcesFetched, products)
	}
}
