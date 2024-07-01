package router

import (
    "github.com/gin-gonic/gin"
    "myproject/controllers"
    "myproject/middlewares"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    userRoutes := r.Group("/users")
    {
        userRoutes.POST("/register", controllers.RegisterUser)
        userRoutes.POST("/login", controllers.LoginUser)
        userRoutes.PUT("/:userId", middlewares.Auth(), controllers.UpdateUser)
        userRoutes.DELETE("/:userId", middlewares.Auth(), controllers.DeleteUser)
    }

    photoRoutes := r.Group("/photos")
    {
        photoRoutes.POST("/", middlewares.Auth(), controllers.UploadPhoto)
        photoRoutes.GET("/", controllers.GetPhotos)
        photoRoutes.PUT("/:photoId", middlewares.Auth(), controllers.UpdatePhoto)
        photoRoutes.DELETE("/:photoId", middlewares.Auth(), controllers.DeletePhoto)
    }

    return r
}
