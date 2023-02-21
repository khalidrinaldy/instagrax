package middlewares

import (
	"github.com/gin-gonic/gin"
	"instagrax/controllers"
	"net/http"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := controllers.ValidateToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": err.Error(),
				"data":    map[string]string{},
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
