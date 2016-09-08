package service

import (
	"errors"
	"log"
	"net"
	"sync"

	geoip2 "github.com/oschwald/geoip2-golang"
)

var loaded bool
var mmdb *geoip2.Reader
var lock sync.Mutex

func init() {
	loaded = false
}

// Response is a struct that holds the data for the JSON HTTP response body.
type Response struct {
	Country string `json:"country"`
}

// LookupIP looks up the specified IP in the loaded Maxmind DB
func LookupIP(ip string) (*Response, error) {
	lock.Lock()
	defer lock.Unlock()

	if !loaded {
		return nil, errors.New("MaxMind DB not loaded")
	}

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

	if loaded {
		mmdb.Close()
		loaded = false
	}

	db, err := geoip2.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Loaded Maxmind DB from " + path)

	mmdb = db
	loaded = true
}

// UnloadMaxmindDB unloads the MaxMind DB from memory.  This is just for testing.
func UnloadMaxmindDB() {
	lock.Lock()
	defer lock.Unlock()

	if !loaded {
		return
	}

	log.Println("Unloaded MaxMind DB")

	mmdb.Close()
	loaded = false
}
