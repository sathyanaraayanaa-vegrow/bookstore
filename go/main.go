package main

import (
	"BookStore/initializers"
	"BookStore/routers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.ConnectToLoggerDB()
	routers.Routes()
}
func main() {
	// r := gin.Default()
	// // r.POST("/", )
	// r.POST("/createbooks", controllers.CreateBooks)
	// r.GET("/books", controllers.ReadBooks)
	// r.Run() // listen and serve on 0.0.0.0:8080
	// result := map[string]interface{}{}
	// initializers.DB.Model(&model.Book{}).Last(&result)
	// fmt.Println(result)
}
