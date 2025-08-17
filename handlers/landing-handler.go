package handlers

import (
	"final-project/repositories"
	"final-project/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLatestProducts(c *gin.Context){
	latestProducts, err := repositories.GetLatestProducts(10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response []schemas.ListProductLandingResponse
	for _, p := range latestProducts {
		urlPath := "localhost:8080/admin/products/view/" + *p.Photo
		p.Photo = &urlPath


		response = append(response, schemas.ListProductLandingResponse{
			Nama:  p.Nama,
			Harga: p.Harga,
			Photo: *p.Photo, 
		})
	}
	
	c.JSON(200, response)
}

func GetAvailableProducts(c *gin.Context) {
	availProducts, err := repositories.GetAvailableProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response []schemas.ListProductLandingResponse
	for _, p := range availProducts {
		urlPath := "localhost:8080/admin/products/view/" + *p.Photo
		p.Photo = &urlPath


		response = append(response, schemas.ListProductLandingResponse{
			Nama:  p.Nama,
			Harga: p.Harga,
			Photo: *p.Photo, 
		})
	}
	
	c.JSON(200, response)
}