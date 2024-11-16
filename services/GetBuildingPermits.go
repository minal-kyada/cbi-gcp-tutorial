package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/models"
	"net/http"
	"strconv"
)

func GetBuildingPermits(db *sql.DB) {
	fmt.Println("Collecting Building Permits Data")

	// Drop table if it exists
	// dropTable := `DROP TABLE IF EXISTS building_permits`
	// _, err := db.Exec(dropTable)
	// if err != nil {
	// 	fmt.Println("Error dropping table: ", err)
	// }

	// Create the table for building permits
	createTable := `
    CREATE TABLE IF NOT EXISTS building_permits (
        id SERIAL PRIMARY KEY,
        permit_id VARCHAR(255) UNIQUE,
        permit_code VARCHAR(255),
        permit_type VARCHAR(255),
        review_type VARCHAR(255),
        application_start_date VARCHAR(255),
        issue_date VARCHAR(255),
        processing_time VARCHAR(255),
        street_number VARCHAR(255),
        street_direction VARCHAR(255),
        street_name VARCHAR(255),
        suffix VARCHAR(255),
        work_description TEXT,
        building_fee_paid VARCHAR(255),
        zoning_fee_paid VARCHAR(255),
        other_fee_paid VARCHAR(255),
        subtotal_paid VARCHAR(255),
        building_fee_unpaid VARCHAR(255),
        zoning_fee_unpaid VARCHAR(255),
        other_fee_unpaid VARCHAR(255),
        subtotal_unpaid VARCHAR(255),
        building_fee_waived VARCHAR(255),
        zoning_fee_waived VARCHAR(255),
        other_fee_waived VARCHAR(255),
        subtotal_waived VARCHAR(255),
        total_fee VARCHAR(255),
        contact_1_type VARCHAR(255),
        contact_1_name VARCHAR(255),
        contact_1_city VARCHAR(255),
        contact_1_state VARCHAR(255),
        contact_1_zipcode VARCHAR(255),
        reported_cost VARCHAR(255),
        pin1 VARCHAR(255),
        pin2 VARCHAR(255),
        community_area VARCHAR(255),
        census_tract VARCHAR(255),
        ward VARCHAR(255),
        xcoordinate DOUBLE PRECISION,
        ycoordinate DOUBLE PRECISION,
        latitude DOUBLE PRECISION,
        longitude DOUBLE PRECISION
    );`

	_, err = db.Exec(createTable)
	if err != nil {
		panic(err)
	}
	fmt.Println("Created Table for Building Permits")

	// While doing unit-testing keep the limit value to 500
	// later you could change it to 1000, 2000, 10,000, etc.
	var url = "https://data.cityofchicago.org/resource/building-permits.json?$limit=100"

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Received data from SODA REST API for Building Permits")

	body, _ := ioutil.ReadAll(res.Body)

	var building_data_list models.BuildingPermitsJsonRecords
	json.Unmarshal(body, &building_data_list)

	// prettyJSON, _ := json.MarshalIndent(building_data_list, "", "    ")
	// fmt.Println("Building Permit Data: ", string(prettyJSON))

	sql := `INSERT INTO building_permits ("permit_id", "permit_code", "permit_type","review_type",
	"application_start_date",
	"issue_date",
	"processing_time",
	"street_number",
	"street_direction",
	"street_name",
	"suffix",
	"work_description",
	"building_fee_paid",
	"zoning_fee_paid",
	"other_fee_paid",
	"subtotal_paid",
	"building_fee_unpaid",
	"zoning_fee_unpaid",
	"other_fee_unpaid",
	"subtotal_unpaid",
	"building_fee_waived",
	"zoning_fee_waived",
	"other_fee_waived",
	"subtotal_waived",
	"total_fee",
	"contact_1_type",
	"contact_1_name",
	"contact_1_city",
	"contact_1_state",
	"contact_1_zipcode",
	"reported_cost",
	"pin1",
	"pin2",
	"community_area",
	"census_tract",
	"ward",
	"xcoordinate",
	"ycoordinate",
	"latitude",
	"longitude" )
	values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,$11, $12, $13, $14, $15,$16, $17, $18, $19, $20,$21, $22, $23, $24, $25,$26, $27, $28, $29,$30,$31, $32, $33, $34, $35,$36, $37, $38, $39, $40)`

	for i := 0; i < len(building_data_list); i++ {

		// We will execute defensive coding to check for messy/dirty/missing data values
		// Any record that has messy/dirty/missing data we don't enter it in the data lake/table

		_, err = db.Exec(
			sql,
			building_data_list[i].Id,
			building_data_list[i].Permit_Code,
			building_data_list[i].Permit_type,
			building_data_list[i].Review_type,
			building_data_list[i].Application_start_date,
			building_data_list[i].Issue_date,
			building_data_list[i].Processing_time,
			building_data_list[i].Street_number,
			building_data_list[i].Street_direction,
			building_data_list[i].Street_name,
			building_data_list[i].Suffix,
			building_data_list[i].Work_description,
			building_data_list[i].Building_fee_paid,
			building_data_list[i].Zoning_fee_paid,
			building_data_list[i].Other_fee_paid,
			building_data_list[i].Subtotal_paid,
			building_data_list[i].Building_fee_unpaid,
			building_data_list[i].Zoning_fee_unpaid,
			building_data_list[i].Other_fee_unpaid,
			building_data_list[i].Subtotal_unpaid,
			building_data_list[i].Building_fee_waived,
			building_data_list[i].Zoning_fee_waived,
			building_data_list[i].Other_fee_waived,
			building_data_list[i].Subtotal_waived,
			building_data_list[i].Total_fee,
			building_data_list[i].Contact_1_type,
			building_data_list[i].Contact_1_name,
			building_data_list[i].Contact_1_city,
			building_data_list[i].Contact_1_state,
			building_data_list[i].Contact_1_zipcode,
			building_data_list[i].Reported_cost,
			building_data_list[i].Pin1,
			building_data_list[i].Pin2,
			building_data_list[i].Community_area,
			building_data_list[i].Census_tract,
			building_data_list[i].Ward,
			parseCoordinate(building_data_list[i].Xcoordinate),
			parseCoordinate(building_data_list[i].Ycoordinate),
			parseCoordinate(building_data_list[i].Latitude),
			parseCoordinate(building_data_list[i].Longitude))

		if err != nil {
			panic(err)
		}
	}
}

func parseCoordinate(coord string) float64 {
	value, _ := strconv.ParseFloat(coord, 64)
	return value
}
