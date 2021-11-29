package middleware

import (
	"errors"
	"kuiz/app/helper"
	"kuiz/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helper.TokenValid(c.Request)
		if err != nil {
			controllers.ErrorResponse(c, http.StatusUnauthorized, "You need to authorized to access this URL", errors.New("Error"))
			c.Abort()
			return
		}
		c.Next()
	}
}
