package main

import(
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

func getAllMeasurements() []string {
	db, err := connectToDb()
	result, err := db.Query("SELECT * FROM measurements")
    if err != nil { panic(err) }

	allMst := []string{}

    for result.Next() {
        var id int
        var temperature float32
		var heatIndex float32
		var humidity float32
		var timestamp string
         
        err = result.Scan(&id, &temperature, &heatIndex, &humidity, &timestamp)
        if err != nil {
            panic(err)
        }
        mst := fmt.Sprintf("%d$%f$%f$%f$%s\n", id, temperature, heatIndex, humidity, timestamp)
		
		allMst = append(allMst, mst)
    }
    defer db.Close()
	return allMst
}