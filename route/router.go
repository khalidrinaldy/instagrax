package route

import (
	"github.com/gin-gonic/gin"
	"instagrax/controllers"
	"instagrax/middlewares"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/users", controllers.GetAllUsers)
	router.POST("/users/login", controllers.Login)
	router.POST("/users/register", controllers.Register)

	protected := router.Group("")
	protected.Use(middlewares.JWTMiddleware())
	protected.PUT("/users/edit", controllers.EditProfile)
	return router
}
