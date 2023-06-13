package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type person struct {
	ID       string `json:"id"`
	FullName string `json:"fullname"`
	Age      int    `json:"age"`
}

var persons = []person{
	{ID: "1", FullName: "Farrel Nikoson", Age: 27},
	{ID: "2", FullName: "Didik Arsayh", Age: 35},
}

func getPersons(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, persons)
}

func postPersons(c *gin.Context) {
	var newPerson person

	if err := c.BindJSON(&newPerson); err != nil {
		return
	}

	persons = append(persons, newPerson)
	c.IndentedJSON(http.StatusCreated, newPerson)
}

func getPersonByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range persons {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
}

func deletePerson(c *gin.Context) {
	id := c.Param("id")

	for i, a := range persons {
		if a.ID == id {
			persons = append(persons[:i], persons[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
}

func main() {
	router := gin.Default()
	router.GET("/persons", getPersons)
	router.POST("/persons", postPersons)
	router.GET("/persons/:id", getPersonByID)
	router.DELETE("/persons/:id", deletePerson)

	router.Run("localhost:8080")
}
