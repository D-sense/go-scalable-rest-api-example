package responder

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserError is an error that doesn't expose internal data
// and can be devileved directly to user
type UserError struct {
	code string
	err  string
}

func (e *UserError) Code() string {
	if e == nil {
		return ""
	}
	return e.code
}

func (e *UserError) Error() string {
	if e == nil {
		return ""
	}
	return e.err
}

var (
	ResourceCreated  = "Resource has been posted successfully"
	ResourceFetched  = "Resource is fetched successfully"
	ResourcesFetched = "Resources are fetched successfully"
)

func (e UserError) MarshalJSON() ([]byte, error) {
	j := struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	}{
		Code:    e.code,
		Message: e.Error(),
	}
	return json.Marshal(&j)
}

type ResponseBody struct {
	Code    int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Success struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

var response ResponseBody

func JsonResponse(c *gin.Context, status bool, message string, data interface{}) {
	response.Success = status
	response.Message = message
	response.Data = data

	c.JSON(http.StatusOK, response)
	return
}
