package main

import (
	"final-project/configs"
	"final-project/databases"
	"final-project/databases/seeders"
	"final-project/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi Kesalahan:", r)
		}
	}()

	//validasi env
	if err := godotenv.Load(); err != nil{
		panic(fmt.Sprintf("No .env file found: %s", err.Error()))
	}

	// koneksi ke db
	configs.SetupMySQL()

	// automigrate
	databases.AutoMigrate()

	//seeder
	seeders.SeederAdmin()
	seeders.SeederProducts()

	//router
	r := gin.Default()
	routes.SetupRoutes(r)
	
	port := os.Getenv("APP_PORT")
	fmt.Println("Running on port:", port)
	r.Run(port) // Jalankan server pada port default 8080

}