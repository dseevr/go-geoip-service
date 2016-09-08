package service

import (
	"errors"
	"log"
	"net"
	"sync"

	geoip2 "github.com/oschwald/geoip2-golang"
)

var mmdb *geoip2.Reader
var lock sync.Mutex

// Response is a struct that holds the data for the JSON HTTP response body.
type Response struct {
	Country string `json:"country"`
}

// LookupIP looks up the specified IP in the loaded MaxmindDB
func LookupIP(ip string) (*Response, error) {
	parsedIP := net.ParseIP(ip) // nil result means error
	if nil == parsedIP {
		return nil, errors.New("failed to parse IP: " + ip)
	}

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
	lock.Lock()
	defer lock.Unlock()

	db, err := geoip2.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Loaded Maxmind DB from " + path)

	mmdb = db
}
