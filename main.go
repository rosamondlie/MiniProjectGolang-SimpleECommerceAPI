package main

import (
	"final-project/configs"
	"final-project/databases"
	"final-project/databases/seeders"
	"final-project/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi Kesalahan:", r)
		}
	}()

	// koneksi ke db
	configs.SetupMySQL()

	// automigrate
	databases.AutoMigrate()

	//seeder
	seeders.SeederAdmin()

	//router
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run() // Jalankan server pada port default 8080

}