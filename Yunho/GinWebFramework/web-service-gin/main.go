package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	//"github.com/go-gorm/gorm"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

type (
	todoModel struct {
		gorm.Model
		Title     string `json:"title"`
		Completed int    `json:"completed"`
	}
	transformedTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)

var db *gorm.DB

func init() {

	var err error
	db, err = gorm.Open("mysql", "root:1234@/mysql?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&todoModel{})
}

func main() {
	router := gin.Default()
	//############게시판 만들기################
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	//###########ORM 실습###################
	//v1 = router.Group("/api/v1/todos")
	//{
	//	v1.POST("/", createTodo)
	//	v1.Get("/", fetchAllToDo)
	//	v1.GET("/:id", fetchSingleTodo)
	//	v1.PUT("/:id", updateTodo)
	//	v1.DELETE("/:id", deleteTodo)
	//
	//}
	router.Run("localhost:8080")

}

// ##################앨범 CRUD##########################
func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

//################## GORM ##########################
