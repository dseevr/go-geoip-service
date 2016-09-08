.PHONY: test

all: test

test:
	go list ./... | grep -v vendor | xargs -I{} go test -v '{}' -check.v
