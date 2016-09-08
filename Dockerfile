FROM golang:1.7.1-alpine

RUN apk add --update git
RUN go get github.com/dseevr/go-geoip-service

ENTRYPOINT ["bin/go-geoip-service"]
CMD [""]
