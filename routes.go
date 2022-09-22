package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getMeasurements returns all measurements
func getMeasurements(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, getAllMeasurements())
}

// postMeasurement adds a new measurement
func postMeasurement(c *gin.Context) {
	var newMeasurement measurement
	if err := c.BindJSON(&newMeasurement); err != nil {
		return
	}
	newMeasurement.ID = measurements[len(measurements)-1].ID + 1

	measurements = append(measurements, newMeasurement)
	c.IndentedJSON(http.StatusCreated, newMeasurement)
}
