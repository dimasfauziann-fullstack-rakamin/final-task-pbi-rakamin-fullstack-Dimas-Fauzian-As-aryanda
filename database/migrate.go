package database

import "myproject/models"

func Migrate() {
    DB.AutoMigrate(&models.User{})
}
