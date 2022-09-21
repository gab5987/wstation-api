package main

import "github.com/gin-gonic/gin"

type measurement struct {
	ID          int     `json:"id"`
	Temperature float32 `json:"temperature"`
	HeatIndex   float32 `json:"heatIndex"`
	Humidity    float32 `json:"humidity"`
	Date        string  `json:"date"`
}

var measurements = []measurement{
	{ID: 1, Temperature: 35.7, HeatIndex: 32.5, Humidity: 99.7, Date: "2019-01-01"},
}

func main() {
	router := gin.Default()
	router.GET("/measurements", getMeasurements)
	router.POST("/measurements", postMeasurement)

	router.Run("localhost:8080")
}
