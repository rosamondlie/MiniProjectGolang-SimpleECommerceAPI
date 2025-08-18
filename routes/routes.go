package routes

import (
	"final-project/handlers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	apiRoutes := r.Group("")
	{
		authRoutes := apiRoutes.Group("/admin")
		{
			authRoutes.POST("/login", handlers.Login)
		}
		publicRoutes := apiRoutes.Group("/products")
		{
			publicRoutes.GET("/latest", handlers.GetLatestProducts)
			publicRoutes.GET("/available", handlers.GetAvailableProducts)
		}
		adminRoutes := apiRoutes.Group("/admin", middlewares.AuthMiddleware())
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
				productRoutes.GET("/view/:filename", handlers.ViewPhotoProduct)
				productRoutes.GET("/export", handlers.ExportProduct)
			}
			dashboardRoutes := adminRoutes.Group("/dashboard")
			{
				dashboardRoutes.GET("/", handlers.GetDashboard)
			}

		}
	}
}
