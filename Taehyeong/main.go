package main

import (
	"runner/album"
	"runner/gobyexample"

	"github.com/gin-gonic/gin"
)

func main() {
	gobyexample.PrintExample("generics")

	router := gin.Default()
	album.Service(router.Group("/albums"))
	router.Run("localhost:8080")
}
