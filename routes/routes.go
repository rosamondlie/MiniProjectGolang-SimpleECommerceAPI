package routes

import (
	"final-project/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	apiRoutes := r.Group("")
	{
		adminRoutes := apiRoutes.Group("/admin")
		{
			userRoutes := adminRoutes.Group("/users")
			{
				userRoutes.GET("/", handlers.ListUser)
				userRoutes.POST("/", handlers.CreateUser)
				userRoutes.GET("/:id", handlers.GetUserByID)
				userRoutes.PUT("/:id", handlers.UpdateUser)
				userRoutes.DELETE("/:id", handlers.DeleteUser)
			}
			productRoutes := adminRoutes.Group("/products")
			{
				productRoutes.GET("/", handlers.ListProduct)
				productRoutes.POST("/", handlers.CreateProduct)
				productRoutes.GET("/:id", handlers.GetProductByID)
				productRoutes.PUT("/:id", handlers.UpdateProduct)
				productRoutes.DELETE("/:id", handlers.DeleteProduct)
			}

		}
	}
}
