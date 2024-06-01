package router

import (
	"golang-backend-template/controller"
	"golang-backend-template/middleware"

	"github.com/gin-gonic/gin"
)

func PostRoutes(router *gin.Engine, postController controller.IPostController) {
	postRoute := router.Group("/post", middleware.WithAuth("user", "admin"))
	{
		// SEE ALL POST ALSO CAN FILTER BY USERNAME '/post?username=Andi'
		postRoute.GET("/", postController.SeeAllPost)
		postRoute.GET("/:id", postController.SeePostById)
		postRoute.POST("/", postController.AddPost)
		postRoute.PUT("/:id", postController.EditPost)
		postRoute.DELETE("/:id", postController.DeletePost)
	}
}
