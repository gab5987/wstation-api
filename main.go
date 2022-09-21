package main

import "github.com/gin-gonic/gin"

type measurement struct {
	ID    string  `json:"id"`
	Value float32 `json:"value"`
	Date  string  `json:"date"`
}

var measurements = []measurement{
	{ID: "1", Value: 1.1, Date: "2019-01-01"},
}

func main() {
	router := gin.Default()
	router.GET("/measurements", getMeasurements)

	router.Run("localhost:8080")
}
