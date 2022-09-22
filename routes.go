package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getMainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// getMeasurements returns all measurements
func getMeasurements(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, getAllMeasurementsFromDB())
}

// postMeasurement adds a new measurement
func postMeasurement(c *gin.Context) {
	var newMeasurement measurement
	if err := c.BindJSON(&newMeasurement); err != nil {
		return
	}
	postMeasurementToDB(newMeasurement)
	c.IndentedJSON(http.StatusCreated, newMeasurement)
}

// returns all measurements that are newer than the given index
func getLastMeasurement(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, getLastMeasurementFromDB())
}
