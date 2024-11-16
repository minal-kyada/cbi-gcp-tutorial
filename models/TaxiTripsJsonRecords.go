package models

////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

// The following is a sample record from the Taxi Trips dataset retrieved from the City of Chicago Data Portal

////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

// trip_id	"c354c843908537bbf90997917b714f1c63723785"
// trip_start_timestamp	"2021-11-13T22:45:00.000"
// trip_end_timestamp	"2021-11-13T23:00:00.000"
// trip_seconds	"703"
// trip_miles	"6.83"
// pickup_census_tract	"17031840300"
// dropoff_census_tract	"17031081800"
// pickup_community_area	"59"
// dropoff_community_area	"8"
// fare	"27.5"
// tip	"0"
// additional_charges	"1.02"
// trip_total	"28.52"
// shared_trip_authorized	false
// trips_pooled	"1"
// pickup_centroid_latitude	"41.8335178865"
// pickup_centroid_longitude	"-87.6813558293"
// pickup_centroid_location
// type	"Point"
// coordinates
// 		0	-87.6813558293
// 		1	41.8335178865
// dropoff_centroid_latitude	"41.8932163595"
// dropoff_centroid_longitude	"-87.6378442095"
// dropoff_centroid_location
// type	"Point"
// coordinates
// 		0	-87.6378442095
// 		1	41.8932163595
////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

type TaxiTripsJsonRecords []struct {
	Trip_id                    string `json:"trip_id"`
	Trip_start_timestamp       string `json:"trip_start_timestamp"`
	Trip_end_timestamp         string `json:"trip_end_timestamp"`
	Pickup_centroid_latitude   string `json:"pickup_centroid_latitude"`
	Pickup_centroid_longitude  string `json:"pickup_centroid_longitude"`
	Dropoff_centroid_latitude  string `json:"dropoff_centroid_latitude"`
	Dropoff_centroid_longitude string `json:"dropoff_centroid_longitude"`
}
