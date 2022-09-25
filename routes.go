package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getMainPage returns the main page of the api
func getMainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.php", nil)
}

// catchErrorPage returns the error page in case of a 404 or smth
func catchErrorPage(c *gin.Context) {
	c.HTML(http.StatusNotFound, "err.html", nil)
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

func getMeasurementsGreaterthan(c *gin.Context) {
	id := c.Query("id")
	c.IndentedJSON(http.StatusOK, getMeasurementsGreaterthanFromDB(id))
}

func getMeasurementsTempGreaterThan(c *gin.Context) {
	temp := c.Query("temp")
	c.IndentedJSON(http.StatusOK, getMeasurementsTempGreaterThanFromDB(temp))
}
