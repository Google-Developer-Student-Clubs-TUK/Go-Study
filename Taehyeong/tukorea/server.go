package tukorea

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server() {
	router := gin.Default()
	tukoreaGroup := router.Group("/tukorea")
	{
		tukoreaGroup.POST("/login", doLogin)
		tukoreaGroup.GET("/timetable", getTimetable)
	}
	router.Run("localhost:8080")
}

func getTimetable(c *gin.Context) {
	newTimetable := getTimetableService()
	c.IndentedJSON(http.StatusOK, newTimetable)
}

func doLogin(c *gin.Context) {
	var loginBody login

	if err := c.BindJSON(&loginBody); err != nil {
		return
	}

	tukoreaLoginBody := tukoreaLogin{usr_id: loginBody.Id, usr_pwd: loginBody.Pwd}

	resp := loginService(tukoreaLoginBody)
	c.IndentedJSON(http.StatusCreated, resp)
}
