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

	protected.GET("/users/:id/posts", controllers.GetUsersAllPosts)
	protected.POST("/users/posts", controllers.CreatePost)
	protected.PUT("/users/posts/:id", controllers.EditPost)
	protected.DELETE("/users/posts/:id", controllers.DeletePost)

	protected.POST("/users/posts/like/:id", controllers.AddLike)
	protected.DELETE("/users/posts/like/:id", controllers.DeleteLike)

	protected.POST("/users/posts/comment/:id", controllers.AddComment)
	protected.DELETE("/users/posts/comment/:id", controllers.DeleteComment)
	return router
}
