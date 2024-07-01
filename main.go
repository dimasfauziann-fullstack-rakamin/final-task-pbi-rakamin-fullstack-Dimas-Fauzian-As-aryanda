// main.go
package main

import (
	"myproject/database"
	"myproject/router"

	"github.com/gin-gonic/gin"
)

func main() {
    database.Connect()
    database.Migrate()
    r := gin.Default()
    router.SetupRoutes(r)
    r.Run(":8080") 
}
