package main

import (
	"flag"
)

var stopsFilename = flag.String("stops", "data/stops.csv", "path to bus stops CSV file")
var routesFilename = flag.String("routes", "data/routes.csv", "path to bus routes CSV file")

// var stopTimesFilename = flag.String("stop-times", "data/stop_times.csv", "path to bus stop times CSV file")
// var tripsFilename = flag.String("trips", "data/trips.csv", "path to bus trips CSV file")

func main() {
	flag.Parse()

	readStops(stopsFilename)
	readRoutes(routesFilename)
	// readStopTimes(stopTimesFilename)
	// readTrips(tripsFilename)
}
