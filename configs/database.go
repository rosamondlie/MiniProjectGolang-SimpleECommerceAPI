package configs

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupMySQL() {
	user := "root"
	password := ""
	host := "localhost"
	port := "3306"
	dbname := "finpro_golang"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		errorLog := fmt.Sprintf("Gagal koneksi ke database: %s", err.Error())
		panic(errorLog)
	}

	DB = db
}
