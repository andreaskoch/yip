# Yip

Yip as in "your ip". Yip is a standalone web-server that responds with the users' IPv6 or IPv4 address.

[![Build Status](https://travis-ci.org/andreaskoch/yip.svg?branch=master)](https://travis-ci.org/andreaskoch/yip)

## Usage

Yip will start a HTTP server listening on tcp4 and tcp6 port 8080 that will do nothing but return the users' IP address:

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

If you want to use a different port or if you just want to listen on tcp4 or tcp6 ports you can specify the bind address as the first argument.

Start yip on port 80:

```bash
yip :80
```

Start yip on IPv4 port 80:

```bash
yip 0.0.0.0:80
```

Start yip on IPv6 port 80:

```bash
yip '[::]:80'
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
git clone git@github.com:andreaskoch/yip.git
cd $GOPATH/github.com/andreaskoch/yip
make crosscompile
```

## Docker

There is docker image with the latest yip binary:

```bash
docker run -ti -rm -p <your-interface-ip>:80:8080 andreaskoch/yip:latest
```

**Note**: When using the docker image it is important that you specify your IP address in the docker port binding, because otherwise yip will always return the IP of the docker bridge (172.17.0.1) instead of the users' ip.

## Roadmap

- Add command line options
- Add direct support for SSL via Let's Encrypt
