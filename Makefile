.PHONY: build test

all: build

build: clean
	go build

test:
	go list ./... | grep -v vendor | xargs -I{} go test -v '{}' -check.v

clean:
	rm -f go-geoip-service
