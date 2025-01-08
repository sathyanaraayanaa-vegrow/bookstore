package routers

import (
	"BookStore/controllers"

	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func Routes() {
	r := gin.Default()
	r.Use(corsMiddleware())
	r.POST("/books", controllers.AuthMiddleware(), controllers.CreateBooks)
	r.GET("/books", controllers.AuthMiddleware(), controllers.ReadAllBooks)
	r.GET("/books/:id", controllers.AuthMiddleware(), controllers.ReadBook)
	r.PUT("/books/:id", controllers.AuthMiddleware(), controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.AuthMiddleware(), controllers.DeleteBook)
	r.POST("/records", controllers.AuthMiddleware(), controllers.CreateRecord)
	r.GET("/records", controllers.AuthMiddleware(), controllers.ViewRecords)
	r.GET("/records/:id", controllers.AuthMiddleware(), controllers.ViewRecords)
	r.GET("/record/:id", controllers.AuthMiddleware(), controllers.ViewRecord)
	r.DELETE("/record/:id", controllers.AuthMiddleware(), controllers.DeleteRecord)
	r.DELETE("/record", controllers.AuthMiddleware(), controllers.DeleteRecordWithBook)
	r.GET("/authors", controllers.AuthMiddleware(), controllers.FetchAuthors)
	r.POST("/authors", controllers.AuthMiddleware(), controllers.AddAuthors)
	r.Run()
}
