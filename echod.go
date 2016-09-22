// echod is a HTTP request & response service, written in Go.
package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("http", "localhost:2333", "HTTP service address (e.g., ':2333')")

func main() {
	flag.Parse()
	// Index page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		index(w, r)
	})
	// Returns origin IP
	http.HandleFunc("/ip", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		ip(w, r)
	})

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatalf("start echod failed: %v", err)
		flag.Usage()
	}
}
