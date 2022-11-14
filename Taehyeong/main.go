package main

import (
	"runner/album"
	"runner/gobyexample"
	"runner/todo"

	"github.com/gin-gonic/gin"
)

func main() {
	gobyexample.PrintExample("generics")

	router := gin.Default()
	album.Service(router.Group("/albums"))
	todo.Service(router.Group("/todo"))
	router.Run("localhost:8080")
}
