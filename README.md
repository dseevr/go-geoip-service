[![Build Status](https://travis-ci.org/dseevr/go-geoip-service.svg?branch=master)](https://travis-ci.org/dseevr/go-geoip-service) ![goreportcard](https://goreportcard.com/badge/github.com/dseevr/go-geoip-service "goreportcard")

# Go GeoIP Service

A Go (golang) HTTP service which will tell you the country associated with a given IP.

## Usage

Go get:
```sh
go get github.com/dseevr/go-geoip-service
$GOPATH/bin/go-geoip-service --db-path=/mm.db --port=1234
```

Build it yourself:
```sh
git clone git@github.com:dseevr/go-geoip-service.git
cd go-geoip-service
make
./go-geoip-service --db-path=/mm.db --port=1234
```

Run from Docker:

```sh
docker run --rm -it \
  -p 1234:1234 \
  -v /foo/bar/GeoLite2-Country.mmdb:/mm.db \
  billrobinson/go-geoip-service --db-path=/mm.db --port=1234
```

## Looking up an IP

```
curl -w "\n" "http://localhost:12345/lookup?ip=1.2.3.4"
```

Response:

```json
{"country":"AU"}
```

## Prerequisites

You will need a GeoLite2 database from MaxMind.  Specify it with the `--db-path` option.

You can get a database here:  https://dev.maxmind.com/geoip/geoip2/geolite2/

## Running the tests

You will need to have a MaxMind database in the base folder called `test.db`.

Then just run `make test`!

## License

BSD
