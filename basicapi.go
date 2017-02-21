package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

var count = 0

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9010", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	count++
	var encoder = json.NewEncoder(w)
	encoder.Encode(map[string]string{"count": strconv.Itoa(count), "ip": r.RemoteAddr})
}
