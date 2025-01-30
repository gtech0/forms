package errs

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		var e CustomError
		switch {
		case errors.As(err.Err, &e):
			c.AbortWithStatusJSON(e.Status, e)
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
		}
	}
}
