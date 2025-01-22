package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpleserver/internal/sse"
	"simpleserver/models"
)

var lines = []models.Line{
	{Points: []int{0, 0}},
}

func ClearAllLines(c *gin.Context) {
	lines = []models.Line{{
		Points: []int{0, 0},
	}}
	c.Done()
}

func GetLines(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, lines)
}

func SetLine(c *gin.Context) {
	var newLine models.Line

	if err := c.ShouldBindJSON(&newLine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	sse.Instance.Broadcast("UpdateLinesEvent")

	lines = append(lines, newLine)

	c.IndentedJSON(http.StatusCreated, newLine)
}
