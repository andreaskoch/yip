FROM golang:latest
MAINTAINER Andreas Koch <andy@ak7.io>

# Add sources
ADD . /go/src/github.com/andreaskoch/yip
WORKDIR /go/src/github.com/andreaskoch/yip

EXPOSE 8080

ENTRYPOINT go run server.go
