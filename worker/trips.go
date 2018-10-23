package worker

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type Trip struct {
	route_id     int
	service_id   int
	id           int
	headsign     string
	direction_id int
	block_id     int
	shape_id     int
}

func ReadTrips(filename string) []Trip {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}

	reader := csv.NewReader(file)
	reader.Read() // throw away the header

	var trips []Trip

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error parsing CSV file: ", err)
		}

		trip := parseTrip(row)
		trips = append(trips, trip)
	}

	return trips
}

func parseTrip(row []string) Trip {
	var trip Trip
	route_id, err := strconv.Atoi(row[0])
	if err != nil {
		log.Fatalln("Error parsing trip route_id: ", err, row)
	}
	trip.route_id = route_id

	service_id, err := strconv.Atoi(row[1])
	if err != nil {
		log.Fatalln("Error parsing trip service_id: ", err, row)
	}
	trip.service_id = service_id

	id, err := strconv.Atoi(row[2])
	if err != nil {
		log.Fatalln("Error parsing trip id: ", err, row)
	}
	trip.id = id

	shape_id, err := strconv.Atoi(row[6])
	if err != nil {
		log.Fatalln("Error parsing trip shape_id: ", err, row)
	}
	trip.shape_id = shape_id

	return trip
}
