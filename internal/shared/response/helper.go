package response

import (
	"github.com/gin-gonic/gin"
)

func Success[T any](c *gin.Context, httpStatus int, code string, message string, data T) {
	c.JSON(httpStatus, SuccessResponse[T]{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, httpStatus int, code string, message string) {
	c.JSON(httpStatus, ErrorResponse{
		Code:    code,
		Message: message,
	})
}

func NewFieldError(field, message string) FieldError {
	return FieldError{
		Field:   field,
		Message: message,
	}
}

func ValidateError(c *gin.Context, httpStatus int, code string, message string, errors []FieldError) {
	c.JSON(httpStatus, ErrorResponse{
		Code:    code,
		Message: message,
		Errors:  errors,
	})
}
