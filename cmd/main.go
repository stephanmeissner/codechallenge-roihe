package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type team struct {
	ID         int `json:"id"`
	Developers int `json:"developers"`
}

var teams = []team{}

type project struct {
	ID          int `json:"id"`
	Devs_needed int `json:"devs_needed"`
}

var projects = []project{}

func main() {
	router := gin.Default()
	router.GET("/status", getStatus)
	router.PUT("/teams", putTeams)
	router.POST("/project", postProject)
	router.POST("/completed", postCompleted)
	router.POST("/assigned", postAssigned)
	router.Run("localhost:3000")
}

func getStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

func putTeams(c *gin.Context) {
	if err := c.BindJSON(&teams); err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, teams)
}

func postProject(c *gin.Context) {
	var newProject project

	if err := c.BindJSON(&newProject); err != nil {
		return
	}

	projects = append(projects, newProject)

	c.IndentedJSON(http.StatusOK, newProject)
}

func postCompleted(c *gin.Context) {

	id := c.Query("id")

	c.IndentedJSON(http.StatusOK, id)
}

func postAssigned(c *gin.Context) {

	id := c.Query("id")

	c.IndentedJSON(http.StatusOK, id)
}
