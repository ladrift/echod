package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "echod(1): HTTP Request & Response Service, a Go fork of httpbin.")
}

func ip(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Printf("%s\n", r.RemoteAddr)
	ip := r.RemoteAddr[:strings.LastIndexByte(r.RemoteAddr, ':')]
	fmt.Printf("%s\n", ip)
	resp := map[string]interface{}{
		"origin": ip,
	}
	e := json.NewEncoder(w)
	e.Encode(resp)
}
