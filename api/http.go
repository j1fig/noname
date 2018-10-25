package api

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var rc *redis.Client


func getRedisClient () *redis.Client {
	if rc == nil {
		rc = redis.NewClient(&redis.Options{
		    Addr:     "redis:6379",
		    Password: "", // no password set
		    DB:       0,  // use default DB
		})
	}
	return rc
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	c := getRedisClient()

	pong, err := c.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}
	fmt.Fprintf(w, "echoing %s! redis: %v %v", r.URL.Path[1:], pong, err)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "echoing %s!", r.URL.Path[1:])
}

func stopsHandler(w http.ResponseWriter, r*http.Request) {
	fmt.Fprintf(w, "echoing stops %s!", r.URL.Path[1:])
}

func ListenAndServe(port int) {
	r := mux.NewRouter()

	// health endpoint. weird name but is the container standard.
	r.HandleFunc("/healthz", healthHandler)

	// HTTP API serving.
	a := r.PathPrefix("/api/").Subrouter()
	a.HandleFunc("/stops/", stopsHandler)
	a.HandleFunc("/", apiHandler)

	// static file serving.
	dir := http.Dir("static/dist")
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
