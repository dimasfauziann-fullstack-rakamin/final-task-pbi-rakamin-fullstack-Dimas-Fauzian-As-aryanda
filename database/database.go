package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    dsn := "myproject_user:myproject_password@tcp(127.0.0.1:3306)/myproject_db?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Println("Error connecting to the database:", err)
        return
    }
    fmt.Println("Database connection successful")
}
