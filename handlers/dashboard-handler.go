package handlers

import (
	"final-project/models"
	"final-project/repositories"
	"final-project/schemas"
	"final-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDashboard(c *gin.Context) {
	var totalUsers, activeUsers, totalProducts, availableProducts int64
	var latestProducts []models.Product

	totalUsers, err := repositories.CountUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	activeUsers, err = repositories.CountActiveUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalProducts, err = repositories.CountProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	availableProducts, err = repositories.CountAvailableProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	latestProducts, err = repositories.GetLatestProducts(3)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ID := 1
	var latestProductsResponse []schemas.ListLatestProductResponse
	for _, p := range latestProducts{
		urlPath := "localhost:8080/admin/products/view/" + *p.Photo
		p.Photo = &urlPath


		latestProductsResponse = append(latestProductsResponse, schemas.ListLatestProductResponse{
			ID:    ID,
			Nama:  p.Nama,
			Harga: p.Harga,
			Photo: *p.Photo, 
			Date: utils.FormatTanggal(p.CreatedAt),
		})
		ID++
	}

	response := schemas.DashboardResponse{
		TotalUsers:          totalUsers,
		ActiveUsers:         activeUsers,
		TotalProducts:       totalProducts,
		TotalAvailProducts:   availableProducts,
		LatestProducts:      latestProductsResponse,
	}
	
	c.JSON(200, response)

}
