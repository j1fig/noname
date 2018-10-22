package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stop struct {
	id             int
	code           int
	name           string
	desc           string
	lat            float64
	lon            float64
	zone_id        int
	url            string
	location_type  string
	parent_station string
}

func readStops(filename string) []Stop {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}

	reader := csv.NewReader(file)
	reader.Read() // throw away the header

	var stops []Stop

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error parsing CSV file: ", err)
		}

		stop := parseStop(row)
		stops = append(stops, stop)
	}

	return stops
}

func parseStop(row []string) Stop {
	var stop Stop
	id_tokens := strings.Split(row[0], "_")
	id, err := strconv.Atoi(id_tokens[1])
	if err != nil {
		log.Fatalln("Error parsing stop id: ", err, row)
	}
	stop.id = id

	stop.name = row[2]
	lat, err := strconv.ParseFloat(row[4], 64)
	if err != nil {
		log.Fatalln("Error parsing stop latitude: ", err, row)
	}
	lon, err := strconv.ParseFloat(row[5], 64)
	if err != nil {
		log.Fatalln("Error parsing stop longitude: ", err, row)
	}

	stop.lat = lat
	stop.lon = lon

	return stop
}
