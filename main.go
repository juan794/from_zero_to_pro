package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.NoRoute(func(ctx *gin.Context) {
		ctx.File("./public/index.html")
	})
	router.Static("/", "./public")
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	server.ListenAndServe()
}
