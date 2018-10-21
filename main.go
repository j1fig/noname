package main

import (
	"encoding/csv"
	"flag"
	"fmt"
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
	location_type   string
	parent_station string
}

var stopsFilename = flag.String("stops", "data/stops.csv", "path to bus stops CSV file")

func main() {
	flag.Parse()

	readStops(stopsFilename)
}

func readStops(filename *string) {
	file, err := os.Open(*filename)

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
		fmt.Println(stop)
		stops = append(stops, stop)
	}

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
