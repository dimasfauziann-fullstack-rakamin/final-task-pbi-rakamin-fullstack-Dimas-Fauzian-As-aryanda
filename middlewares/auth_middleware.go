package middlewares

import (
    "github.com/gin-gonic/gin"
    "myproject/helpers"
    "net/http"
)

func Auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        userId, err := helpers.ValidateJWT(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            c.Abort()
            return
        }
        c.Set("userId", userId)
        c.Next()
    }
}
