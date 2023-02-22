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

	protected.GET("/users/posts/:id", controllers.GetUsersAllPosts)
	protected.POST("/posts", controllers.CreatePost)
	protected.PUT("/posts/:id", controllers.EditPost)
	protected.DELETE("/posts/:id", controllers.DeletePost)

	protected.POST("/posts/like/:id", controllers.Like)

	protected.GET("/posts/comment/:id", controllers.GetAllComment)
	protected.POST("/posts/comment/:id", controllers.AddComment)
	protected.DELETE("/posts/comment/:id", controllers.DeleteComment)
	return router
}
