package main

import (
	"flag"
	"gitlab.com/j1fig/noname/worker"
	"gitlab.com/j1fig/noname/api"
	"log"
)

const (
	STOPSFILE     = "data/stops.csv"
	ROUTESFILE    = "data/routes.csv"
	STOPTIMESFILE = "data/stop_times.csv"
	TRIPSFILE     = "data/trips.csv"
)

var cold = flag.Bool("cold", false, "cold starts the app by reloading data files")
var request = flag.Int("request", 0, "request wait time for a given stop ID")
var _time = flag.Int("time", 0, "displays waiting time (min) for a given stop ID")
var port = flag.Int("port", 5000, "HTTP web server port")


func main() {
	flag.Parse()
	log.SetPrefix("[noname] ")

	if *request != 0 {
		response := worker.RequestTime(*request)
		log.Printf("request wait time status for %v: %v\n", *request, response.Status)
		return
	}

	if *_time != 0 {
		worker.GetMessages()
		return
	}

	if *cold {
		stops := worker.ReadStops(STOPSFILE)
		log.Printf("read %v stops\n", len(stops))
		routes := worker.ReadRoutes(ROUTESFILE)
		log.Printf("read %v routes\n", len(routes))
		stopTimes := worker.ReadStopTimes(STOPTIMESFILE)
		log.Printf("read %v stop times\n", len(stopTimes))
		trips := worker.ReadTrips(TRIPSFILE)
		log.Printf("read %v trips\n", len(trips))
	}

	api.ListenAndServe(*port)
}
