package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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
	location_ype   string
	parent_station string
}

var stopsFilename = flag.String("stops", "data/stops.csv", "path to bus stops CSV file")

func main() {
	flag.Parse()

	file, err := os.Open(*stopsFilename)

	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}

	reader := csv.NewReader(file)

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
	id, _ := strconv.Atoi(row[0])
	// pretty sure theres a conversion pro here
	stop.id = id

	return stop
}
