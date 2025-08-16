package handlers

import (
	"final-project/models"
	"final-project/repositories"
	"final-project/schemas"
	"fmt"
	"net/http"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ListProduct(c *gin.Context) {
	product, err := repositories.ListProduct()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	product, err := repositories.GetProductByID(uint(idInt))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	urlPath := "localhost:8080/admin/products/view/" + *product.Photo
	product.Photo = &urlPath

	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var request schemas.CreateProductRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repositories.GetUserByID(request.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	allowedExt := []string{".jpg", ".jpeg", ".png"}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !slices.Contains(allowedExt, ext) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not Valid Extension"})
		return
	}

	// id := c.Param("id")
	// _, err = strconv.Atoi(id)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
	// 	return
	// }

	uploadDir := "./uploads"
	filename := fmt.Sprintf("%s%s", request.Nama, ext)

	filepath := filepath.Join(uploadDir, filename)
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gagal menyimpan file"})
		return
	}

	product := models.Product{
		Nama:   request.Nama,
		Harga:  request.Harga,
		Stok:   request.Stok,
		UserID: user.ID,
		Photo:  &filename,
	}

	err = repositories.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var request schemas.CreateProductRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := repositories.GetProductByID(uint(idInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	user, err := repositories.GetUserByID(request.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	allowedExt := []string{".jpg", ".jpeg", ".png"}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !slices.Contains(allowedExt, ext) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not Valid Extension"})
		return
	}

	// id := c.Param("id")
	// _, err = strconv.Atoi(id)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
	// 	return
	// }

	uploadDir := "./uploads"
	filename := fmt.Sprintf("%s%s", request.Nama, ext)

	filepath := filepath.Join(uploadDir, filename)
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gagal menyimpan file"})
		return
	}

	product.Nama = request.Nama
	product.Harga = request.Harga
	product.Stok = request.Stok
	product.UserID = user.ID
	product.Photo = &filename

	err = repositories.UpdateProduct(product.ID, product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = repositories.DeleteProduct(uint(idInt))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func ViewPhotoProduct(c *gin.Context) {
	filename := c.Param("filename")
	filePath := fmt.Sprintf("./uploads/%s", filename)

	c.File(filePath)
}
