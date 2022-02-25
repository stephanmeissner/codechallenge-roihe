package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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
	router.Use(RequestLogger())
	router.GET("/status", getStatus)
	router.PUT("/teams", putTeams)
	router.POST("/project", postProject)
	router.POST("/completed", postCompleted)
	router.POST("/assigned", postAssigned)
	// router.Run(":3000")
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
	for k := range teams {
		delete(teams, k)
	}
	tc := 0
	for _, ct := range newTeams {
		tc++

		if ct.Developers == 0 || ct.Id == 0 {
			c.IndentedJSON(http.StatusBadRequest, teams)
			return
		}
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

		if p.Devs_needed == 0 || p.Id == 0 {
			c.IndentedJSON(http.StatusBadRequest, teams)
			return
		}

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
	c.IndentedJSON(http.StatusNotFound, nil)
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
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest, id)
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

		fmt.Println(readBody(rdr1)) // Print request body

		c.Request.Body = rdr2
		c.Next()
	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}
