package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	router.GET("/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.JSON(200, gin.H{
			"message": "Hola " + name,
		})
	})
	router.Run(":3000")

}
