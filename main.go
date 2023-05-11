package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.NoRoute(func(ctx *gin.Context) {
		ctx.File("./public/index.html")
	})
	router.Static("/", "./public")
	router.Run(":3000")
}
