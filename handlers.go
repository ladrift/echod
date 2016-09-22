package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `echod(1): HTTP Request & Response Service, a Go fork of httpbin.

Endpoint
	/ This page.
	/ip Returns Origin IP.`)
}

func ip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ip := r.RemoteAddr[:strings.LastIndexByte(r.RemoteAddr, ':')]
	resp := map[string]interface{}{
		"origin": ip,
	}
	e := json.NewEncoder(w)
	e.Encode(resp)
}
