package main

import (
	"database/sql"
	"fmt"
	"log"
	"main/services"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	// Establish connection to Postgres Database
	// db_connection := "user=postgres dbname=chicago_business_intelligence password=root host=localhost sslmode=disable"  //local
	// db_connection := "user=postgres dbname=chicago_business_intelligence password=root host=postgresdb sslmode=disable port = 5433" //docker

	//Database application running on Google Cloud Platform.
	db_connection := "user=postgres dbname=chicago_business_intelligence password=root host=/cloudsql/cbi-gcp-tutorial-a20544029:us-central1:mypostgres sslmode=disable port = 5432"

	db, err := sql.Open("postgres", db_connection)
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	}()

	err = db.Ping()
	if err != nil {
		fmt.Println("Couldn't Connect to database")
		panic(err)
	}

	// Spin in a loop and pull data from the city of chicago data portal
	// Once every hour, day, week, etc.
	// Though, please note that Not all datasets need to be pulled on daily basis
	// fine-tune the following code-snippet as you see necessary
	for {
		// build and fine-tune functions to pull data from different data sources
		// This is a code snippet to show you how to pull data from different data sources.

		fmt.Println(("Calling CBI microservices"))

		services.GetUnemploymentRates(db)
		fmt.Println("Done with unemployement rates")

		services.GetBuildingPermits(db)
		fmt.Println("Done with building permits")

		services.GetTaxiTrips(db)
		fmt.Println("Done with taxi trips")

		// Pull the data once a day
		// You might need to pull Taxi Trips and COVID data on daily basis
		// but not the unemployment dataset becasue its dataset doesn't change every day
		time.Sleep(24 * time.Hour)
	}

}
