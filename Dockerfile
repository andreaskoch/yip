FROM golang:latest
MAINTAINER Andreas Koch <andy@ak7.io>

# Add sources
ADD . /go/src/github.com/andreaskoch/yip
WORKDIR /go/src/github.com/andreaskoch/yip

# Build
RUN go run make.go -install && mv bin/yip /usr/bin/yip

# Cross-compile
RUN go run make.go -crosscompile

EXPOSE 8080

ENTRYPOINT go run server.go
