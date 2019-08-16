package main

import  "github.com/gin-gonic/gin"
func main() {
	g := gin.Default()
	g.GET("/gin", func(context *gin.Context) {
		context.JSON(200,gin.H{"message":"hello gin"})
	})
	g.Run()
}
