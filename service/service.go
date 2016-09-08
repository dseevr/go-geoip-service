package service

import (
	"log"
	"net"

	geoip2 "github.com/oschwald/geoip2-golang"
)

var mmdb *geoip2.Reader

// Response is a struct that holds the data for the JSON HTTP response body.
type Response struct {
	Country string `json:"country"`
}

// LookupIP looks up the specified IP in the loaded MaxmindDB
func LookupIP(ip string) (*Response, error) {
	parsedIP := net.ParseIP(ip)
	record, err := mmdb.City(parsedIP)
	if err != nil {
		return nil, err
	}

	response := &Response{
		Country: record.Country.IsoCode,
	}

	return response, nil
}

// LoadMaxmindDB loads a MaxMind DB into memory for use by the /lookup endpoint.
func LoadMaxmindDB(path string) {
	db, err := geoip2.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Loaded MMDB from " + path)

	mmdb = db
}
