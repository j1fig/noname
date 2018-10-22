package main

import (
	"flag"
	"log"
)

var stopsFilename = flag.String("stops", "data/stops.csv", "path to bus stops CSV file")
var routesFilename = flag.String("routes", "data/routes.csv", "path to bus routes CSV file")
var stopTimesFilename = flag.String("stop-times", "data/stop_times.csv", "path to bus stop times CSV file")
var tripsFilename = flag.String("trips", "data/trips.csv", "path to bus trips CSV file")

func main() {
	flag.Parse()
	log.SetPrefix("[noname] ")

	stops := readStops(*stopsFilename)
	log.Printf("read %v stops\n", len(stops))
	routes := readRoutes(*routesFilename)
	log.Printf("read %v routes\n", len(routes))
	stopTimes := readStopTimes(*stopTimesFilename)
	log.Printf("read %v stop times\n", len(stopTimes))
	trips := readTrips(*tripsFilename)
	log.Printf("read %v trips\n", len(trips))

	response := requestTime(4207)
	log.Println("request wait time status for 4207: ", response.Status)

	getMessages()
}
