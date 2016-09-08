package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"

	geoip2 "github.com/oschwald/geoip2-golang"
)

// Response is a struct that holds the data for the JSON HTTP response body.
type Response struct {
	Country string `json:"country"`
}

func lookupIP(ip string) *Response {
	parsedIP := net.ParseIP(ip)
	record, err := mmdb.City(parsedIP)
	if err != nil {
		log.Fatal(err)
	}

	return &Response{
		Country: record.Country.IsoCode,
	}
}

func lookupHandler(w http.ResponseWriter, req *http.Request) {
	ip := req.URL.Query().Get("ip")

	record := lookupIP(ip)

	bytes, err := json.Marshal(record)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(bytes))
}

func loadMaxmindDB(path string) *geoip2.Reader {
	db, err := geoip2.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

var mmdb *geoip2.Reader

func init() {
	mmdb = loadMaxmindDB("GeoLite2-Country.mmdb")
}

func main() {
	http.HandleFunc("/lookup", lookupHandler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
