package worker

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type Route struct {
	id         int
	name       string
	short_name string
	long_name  string
	desc       string
	type_      int
	url        string
	color      string
	text_color string
	agency_id  int
}

func ReadRoutes(filename string) []Route {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}

	reader := csv.NewReader(file)
	reader.Read() // throw away the header

	var routes []Route

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error parsing CSV file: ", err)
		}

		route := parseRoute(row)
		routes = append(routes, route)
	}

	return routes
}

func parseRoute(row []string) Route {
	var route Route
	agency_id, err := strconv.Atoi(row[0])
	if err != nil {
		log.Fatalln("Error parsing route agency_id: ", err, row)
	}
	route.agency_id = agency_id

	id, err := strconv.Atoi(row[1])
	if err != nil {
		log.Fatalln("Error parsing route id: ", err, row)
	}
	route.id = id

	route.long_name = row[3]

	type_, err := strconv.Atoi(row[5])
	if err != nil {
		log.Fatalln("Error parsing route type: ", err, row)
	}
	route.type_ = type_

	return route
}
