package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"

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

	log.Println("Loaded MMDB from " + path)

	return db
}

// -------------------------------------------------------------------------------------------------

var mmdb *geoip2.Reader

var (
	dbPath string
	port   int
)

func main() {
	flag.StringVar(&dbPath, "db-path", "", "path to MaxMind GeoLite2 database")
	flag.IntVar(&port, "port", 12345, "http port to listen on")
	flag.Parse()

	if 0 == len(dbPath) {
		log.Fatalln("you must specify a --db-path")
	}

	// TODO: allow port 0? not sure if it's worth it
	if port < 1 || port > 65535 {
		log.Fatalln("--port must be >= 1 and <= 65535")
	}

	mmdb = loadMaxmindDB(dbPath)

	stringPort := strconv.Itoa(port)

	log.Println("Listening on 0.0.0.0:" + stringPort)

	http.HandleFunc("/lookup", lookupHandler)
	log.Fatal(http.ListenAndServe(":"+stringPort, nil))
}
