package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type team struct {
	Id         int `json:"id"`
	Developers int `json:"developers"`
}

var teams map[int]team

type project struct {
	Id          int `json:"id"`
	Devs_needed int `json:"devs_needed"`
	Team_id     int
	Index       int
}

var projects map[int]project

func main() {
	teams = make(map[int]team)
	projects = make(map[int]project)

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
	var newTeams []team
	if err := c.BindJSON(&newTeams); err != nil {
		return
	}
	tc := 0
	for _, ct := range newTeams {
		tc++
		teams[tc] = ct
	}
	c.IndentedJSON(http.StatusOK, teams)
}

func postProject(c *gin.Context) {
	var newProject project
	if err := c.BindJSON(&newProject); err != nil {
		return
	}
	projectFound := false
	for _, p := range projects {
		if p.Id == newProject.Id {
			p.Devs_needed = newProject.Devs_needed
			projectFound = true
		}
	}
	for _, t := range teams {
		if t.Developers == newProject.Devs_needed {
			newProject.Team_id = t.Id
		}
	}
	if !projectFound {
		newProject.Index = len(projects) + 1
		projects[newProject.Index] = newProject
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func postCompleted(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "wrong id")
		return
	}
	// search for a project
	for _, p := range projects {
		if p.Id == id {
			delete(projects, p.Index)
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, id)
}

func postAssigned(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "wrong id")
		return
	}
	// search for a project
	for _, p := range projects {
		if p.Id == id {
			if p.Team_id == 0 {
				c.IndentedJSON(http.StatusNoContent, nil)
				return
			}
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, id)
}
