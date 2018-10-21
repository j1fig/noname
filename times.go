package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type StopTime struct {
	trip_id             int
	arrival_time        time.Time
	departure_time      time.Time
	stop_id             int
	stop_sequence       int
	stop_headsign       int
	pickup_type         int
	drop_off_type       int
	shape_dist_traveled float64
}

func readStopTimes(filename *string) []StopTime {
	file, err := os.Open(*filename)

	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}

	reader := csv.NewReader(file)
	reader.Read() // throw away the header

	var stopTimes []StopTime

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error parsing CSV file: ", err)
		}

		stopTime := parseStopTime(row)
		stopTimes = append(stopTimes, stopTime)
	}

	return stopTimes
}

func parseStopTime(row []string) StopTime {
	var stopTime StopTime

	trip_id, err := strconv.Atoi(row[0])
	if err != nil {
		log.Fatalln("Error parsing trip_id: ", err, row)
	}

	// Date parsing code here!
	// hourMinuteSecond := "15:24:06"
	// arrival_time, err := time.Parse(hourMinuteSecond, row[1])
	// if err != nil {
	// 	log.Fatalln("Error parsing arrival_time: ", err, row)
	// }

	// departure_time, err := time.Parse(hourMinuteSecond, row[2])
	// if err != nil {
	// 	log.Fatalln("Error parsing departure_time: ", err, row)
	// }

	stop_id_tokens := strings.Split(row[3], "_")
	stop_id, err := strconv.Atoi(stop_id_tokens[1])
	if err != nil {
		log.Fatalln("Error parsing stop_id: ", err, row)
	}

	stop_sequence, err := strconv.Atoi(row[4])
	if err != nil {
		log.Fatalln("Error parsing stop_sequence: ", err, row)
	}

	stopTime.trip_id = trip_id
	// stopTime.arrival_time = arrival_time
	// stopTime.departure_time = departure_time
	stopTime.stop_id = stop_id
	stopTime.stop_sequence = stop_sequence

	return stopTime
}
