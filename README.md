# Yip

Yip as in "your ip". Yip is a standalone web-server that responds with the users' IPv6 or IPv4 address.

[![Build Status](https://travis-ci.org/andreaskoch/yip.svg?branch=master)](https://travis-ci.org/andreaskoch/yip)

## Usage

Yip will start a HTTP server listening on port 8080 that will do nothing but return the users' IP address:

**Using go**

```bash
go run server.go
```

**Using the binaries**

```bash
yip
```

**Test the server**

```bash
curl http://<your-servers-ip>:8080
```

## Download

You can download the latest binaries of yip for your platform (Linux, Windows or Mac OS) from the releases-section:

[github.com/andreaskoch/yip/releases](https://github.com/andreaskoch/yip/releases)

## Installation

**If you have [go](https://golang.org) installed**

```bash
go get github.com/andreaskoch/yip
```

or

```bash
git clone git@github.com:andreaskoch/yip.git && cd yip && make install
```

## Cross Compilation

Yip comes with Make file that allows you to easily cross-compile yip for Linux (64bit, ARM, ARM5, ARM6, ARM7), Mac OS (64bit) and Windows (64bit).

```bash
make crosscompile
```

## Docker

Yip has a trusted build on docker hub ([andreaskoch/yip:stable](https://hub.docker.com/r/andreaskoch/yip/)) that you can directly use without having to install yip or go:

```bash
docker run -ti -rm -p <your-interface-ip>:80:8080 andreaskoch/yip
```

**Note**: When using the docker image it is important that you specify your IP address in the docker port binding, because otherwise yip will always return the IP of the docker bridge (172.17.0.1) instead of the users' ip.

## Roadmap

- Add command line options
- Add direct support for SSL via Let's Encrypt
