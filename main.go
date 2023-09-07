package main
 
import (
   "fmt"
 
   "project/controllers"
   "project/database"
   "github.com/gin-gonic/gin"
)
 
func main() {
   fmt.Println("Starting application ...")
   database.DatabaseConnection()
 
   r := gin.Default()
   r.GET("/preman/:id", controllers.ReadPreman)
   r.GET("/premans", controllers.ReadPremans)
   r.POST("/premans", controllers.CreatePreman)
   r.PUT("/premans/:id", controllers.UpdatePreman)
   r.DELETE("/premans/:id", controllers.DeletePreman)
   r.Run(":5000")
}