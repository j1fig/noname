package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "echoing %s!", r.URL.Path[1:])
}

func main() {
	flag.Parse()
	log.SetPrefix("[noname] ")

	if *request != 0 {
		response := requestTime(*request)
		log.Printf("request wait time status for %v: %v\n", *request, response.Status)
		return
	}

	if *_time != 0 {
		getMessages()
		return
	}

	if *cold {
		stops := readStops(STOPSFILE)
		log.Printf("read %v stops\n", len(stops))
		routes := readRoutes(ROUTESFILE)
		log.Printf("read %v routes\n", len(routes))
		stopTimes := readStopTimes(STOPTIMESFILE)
		log.Printf("read %v stop times\n", len(stopTimes))
		trips := readTrips(TRIPSFILE)
		log.Printf("read %v trips\n", len(trips))
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *port), nil))
}
