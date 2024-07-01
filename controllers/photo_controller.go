package controllers

import (
    "github.com/gin-gonic/gin"
    "myproject/database"
    "myproject/models"
    "net/http"
    "time"
)

func UploadPhoto(c *gin.Context) {
    var input models.Photo
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    userID := c.MustGet("userId").(uint)
    input.UserID = userID
    input.CreatedAt = time.Now()
    input.UpdatedAt = time.Now()
    database.DB.Create(&input)
    c.JSON(http.StatusOK, input)
}

func GetPhotos(c *gin.Context) {
    var photos []models.Photo
    database.DB.Find(&photos)
    c.JSON(http.StatusOK, photos)
}

func UpdatePhoto(c *gin.Context) {
    // Implementasi fungsi untuk mengupdate photo
}

func DeletePhoto(c *gin.Context) {
    // Implementasi fungsi untuk menghapus photo
}
