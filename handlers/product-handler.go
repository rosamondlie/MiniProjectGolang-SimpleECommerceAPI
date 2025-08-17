package handlers

import (
	"bytes"
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
	"github.com/gin-gonic/gin/binding"
	"github.com/xuri/excelize/v2"
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
	if err := c.ShouldBindWith(&request, binding.FormMultipart); err != nil {
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
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
		Nama:            request.Nama,
		Harga:           request.Harga,
		Stok:            *request.Stok,
		PenanggungJawab: user.Nama,
		Photo:           &filename,
	}

	err = repositories.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error creating product"})
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
	if err := c.ShouldBindWith(&request, binding.FormMultipart); err != nil {
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

	if request.Stok != nil {
		product.Stok = *request.Stok
	}

	product.Nama = request.Nama
	product.Harga = request.Harga
	product.PenanggungJawab = user.Nama
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

func ExportProduct(c *gin.Context) {
	f := excelize.NewFile()
	sheetName := "Sheet1"
	f.SetCellValue(sheetName, "A1", "No")
	f.SetCellValue(sheetName, "B1", "Nama Produk")
	f.SetCellValue(sheetName, "C1", "Harga")
	f.SetCellValue(sheetName, "D1", "Penanggung Jawab")
	f.SetCellValue(sheetName, "E1", "Stok")

	products, err := repositories.ListProduct()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for index, p := range products {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", index+2), p.ID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", index+2), p.Nama)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", index+2), p.Harga)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", index+2), p.PenanggungJawab)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", index+2), p.Stok)
	}

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to write Excel file"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename-users.xlsx")
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buf.Bytes())
}
