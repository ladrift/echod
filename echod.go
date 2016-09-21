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
	http.HandleFunc("/", index)
	http.HandleFunc("/ip", ip)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatalf("start echod failed: %v", err)
		flag.Usage()
	}
}
