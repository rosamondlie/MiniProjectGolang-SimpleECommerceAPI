package routes

import (
	"final-project/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine){
	apiRoutes := r.Group("")
	{
		adminRoutes := apiRoutes.Group("/admin")
		{
			adminRoutes.GET("/users", handlers.ListUser)
			adminRoutes.POST("/users", handlers.CreateUser)
			adminRoutes.GET("/users/:id", handlers.GetUserByID)
			adminRoutes.PUT("/users/:id", handlers.UpdateUser)
			adminRoutes.DELETE("/users/:id", handlers.DeleteUser)
		}
	}
}