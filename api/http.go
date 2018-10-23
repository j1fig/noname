package api

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// probably inefficient to allocate a new connection pool *every time* there is a request
	// which I think is what the library does here.
	rc := redis.NewClient(&redis.Options{
	    Addr:     "redis:6379",
	    Password: "", // no password set
	    DB:       0,  // use default DB
	})

	pong, err := rc.Ping().Result()
	fmt.Println(pong, err)
	fmt.Fprintf(w, "echoing %s! redis: %v %v", r.URL.Path[1:], pong, err)
}

func ListenAndServe(port int) {
	fs := http.FileServer(http.Dir("static/dist"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/healthz/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
