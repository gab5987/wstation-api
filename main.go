package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type measurement struct {
	ID          int     `json:"id"`
	Temperature float32 `json:"temperature"`
	HeatIndex   float32 `json:"heatIndex"`
	Humidity    float32 `json:"humidity"`
	Timestamp   string  `json:"timestamp"`
}

var measurements = []measurement{
	{ID: 1, Temperature: 35.7, HeatIndex: 32.5, Humidity: 99.7, Timestamp: "2019-01-01"},
}

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("static/index.php", "static/err.html")
	router.Static("/static", "./static/")
	router.NoRoute(catchErrorPage)

	router.GET("/", getMainPage) // returns a page containing information about the api

	router.GET("/measurements", getMeasurements)
	router.GET("/measurements/last", getLastMeasurement)
	router.GET("/measurements/greaterthan", getMeasurementsGreaterthan)
	router.GET("/measurements/tempgreaterthan", getMeasurementsTempGreaterThan)

	router.POST("/measurements", postMeasurement)
	router.Run("localhost:8080")
}

func getDotEnvVAriable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
