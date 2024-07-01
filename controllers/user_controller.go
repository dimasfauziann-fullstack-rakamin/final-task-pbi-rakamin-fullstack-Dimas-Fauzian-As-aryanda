package controllers

import (
    "github.com/gin-gonic/gin"
    "myproject/database"
    "myproject/models"
    "myproject/helpers"
    "net/http"
    "time"
)

func RegisterUser(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    hashedPassword, _ := helpers.HashPassword(input.Password)
    input.Password = hashedPassword
    input.CreatedAt = time.Now()
    input.UpdatedAt = time.Now()
    database.DB.Create(&input)
    c.JSON(http.StatusOK, input)
}

func LoginUser(c *gin.Context) {
    var input models.User
    var user models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Where("email = ?", input.Email).First(&user)
    if user.ID == 0 || !helpers.CheckPasswordHash(input.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }
    token, _ := helpers.GenerateJWT(user.ID)
    c.JSON(http.StatusOK, gin.H{"token": token})
}

// Implementasikan fungsi UpdateUser dan DeleteUser serupa dengan RegisterUser
