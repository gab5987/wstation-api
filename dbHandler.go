package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func connectToDb() (*sql.DB, error) {
	dbCredentials := getDotEnvVAriable("DB_CREDENTIALS")
	db, err := sql.Open("mysql", dbCredentials)
	if err != nil {
		panic(err)
	}
	return db, err
}

func getAllMeasurementsFromDB() []string {
	db, err := connectToDb()
	response, err := db.Query("SELECT * FROM measurements")
	if err != nil {
		panic(err)
	}

	allMst := []string{}

	for response.Next() {
		var id int
		var temperature float32
		var heatIndex float32
		var humidity float32
		var timestamp string

		err = response.Scan(&id, &temperature, &heatIndex, &humidity, &timestamp)
		if err != nil {
			panic(err)
		}
		mst := fmt.Sprintf("%d$%f$%f$%f$%s\n", id, temperature, heatIndex, humidity, timestamp)

		allMst = append(allMst, mst)
	}
	defer db.Close()
	return allMst
}

func postMeasurementToDB(newMeasurement measurement) {
	db, err := connectToDb()
	if err != nil {
		panic(err)
	}
	stmt := "VALUES (NULL," + fmt.Sprintf("%v", newMeasurement.Temperature) + "," + fmt.Sprintf("%v", newMeasurement.HeatIndex) + "," + fmt.Sprintf("%v", newMeasurement.Humidity) + ",'" + fmt.Sprintf("%v", newMeasurement.Timestamp) + "')"

	_, err = db.Query("INSERT INTO `measurements`" + stmt)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func getLastMeasurementFromDB() string {
	db, err := connectToDb()
	if err != nil {
		panic(err)
	}

	response, err := db.Query("SELECT * FROM measurements ORDER BY id DESC LIMIT 1")
	mst := ""
	for response.Next() {
		var id int
		var temperature float32
		var heatIndex float32
		var humidity float32
		var timestamp string

		err = response.Scan(&id, &temperature, &heatIndex, &humidity, &timestamp)
		if err != nil {
			panic(err)
		}
		mst = fmt.Sprintf("%d$%f$%f$%f$%s\n", id, temperature, heatIndex, humidity, timestamp)
	}
	return fmt.Sprintf("%v", mst)
}
