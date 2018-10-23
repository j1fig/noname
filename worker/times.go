package worker

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type StopTime struct {
	tripId             int
	arrivalTime        time.Time
	departureTime      time.Time
	stopId             int
	stopSequence       int
	stop_headsign       int
	pickup_type         int
	drop_off_type       int
	shape_dist_traveled float64
}

func ReadStopTimes(filename string) []StopTime {
	file, err := os.Open(filename)

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

	tripId, err := strconv.Atoi(row[0])
	if err != nil {
		log.Fatalln("Error parsing trip_id: ", err, row)
	}

	arrivalTime := parseTimeOfDay(row[1])
	departureTime := parseTimeOfDay(row[2])

	stopIdTokens := strings.Split(row[3], "_")
	stopId, err := strconv.Atoi(stopIdTokens[1])
	if err != nil {
		log.Fatalln("Error parsing stop_id: ", err, row)
	}

	stopSequence, err := strconv.Atoi(row[4])
	if err != nil {
		log.Fatalln("Error parsing stop_sequence: ", err, row)
	}

	stopTime.tripId = tripId
	stopTime.arrivalTime = arrivalTime
	stopTime.departureTime = departureTime
	stopTime.stopId = stopId
	stopTime.stopSequence = stopSequence

	return stopTime
}

func parseTimeOfDay(t string) time.Time {
	hourMinuteSecond := "15:04:05"
	pt, err := time.Parse(hourMinuteSecond, t)
	if err != nil {
		// time can span over the 24h period if still on the same service day.
		// see https://developers.google.com/transit/gtfs/reference/#stop_timestxt
		nextDay := strings.Compare("23:59:59", t) < 0
		// we can do this because we'll still know it's in the next day as the data
		// pertains to the same trip_id and we have stop_sequence to enforce the order.
		if nextDay {
			var hour, minute, second int
			fmt.Sscanf(t, "%d:%d:%d", &hour, &minute, &second)
			hour -= 24
			return time.Date(0, 0, 0, hour, minute, second, 0, time.UTC)
		}
	}

	return pt
}
