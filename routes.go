package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getMeasurements returns all measurements
func getMeasurements(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, measurements)
}
