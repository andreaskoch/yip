// Copyright 2016 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

// GitInfo is either the empty string (the default)
// or is set to the git hash of the most recent commit
// using the -X linker flag (Example: "2016-02-06-284c030+")
var GitInfo string

// version returns the git version of this binary (e.g. "2016-02-06-284c030+").
// If the linker flags were not provided, the return value is "unknown".
func version() string {
	if GitInfo != "" {
		return GitInfo
	}

	return "unknown"
}

// privateNetworkMasks contains a list of all private network masks
var privateNetworkMasks = func() []net.IPNet {
	masks := []net.IPNet{}
	for _, cidr := range []string{"10.0.0.0/8", "172.16.0.0/12", "172.17.0.0/12", "192.168.0.0/16", "fd00::/8"} {
		_, net, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(err)
		}
		masks = append(masks, *net)
	}

	return masks
}()

func main() {
	ipServer(":8080")
}

// ipServer starts the IP server on the given bind address (e.g. ":80").
func ipServer(bindAddress string) {

	// server version
	if version() != "unknown" {
		log.Println("[INFO]", "Starting server", version())
	} else {
		log.Println("[INFO]", "Starting server")
	}

	http.HandleFunc("/", ipResponder)
	log.Fatal(http.ListenAndServe(bindAddress, nil))

}

// ipResponder is a HTTP-handler func that parses the request
// and writes the requests' address into the response.
func ipResponder(w http.ResponseWriter, r *http.Request) {

	// try to get IP from http headers
	ip, err := getIPFromRequest(r)
	if err != nil {
		log.Printf("[ERROR] %s\n", err.Error())
		http.Error(w, "Unable to determine remote address from request", 400)
		return
	}

	log.Println("[INFO]", ip.String(), r.UserAgent())

	w.WriteHeader(200)
	fmt.Fprintf(w, "%s\n", ip.String())
}

// getIPFromRequest tries to return the remote IP address from the given request.
// It will return an error if it fails to extract the IP address.
func getIPFromRequest(r *http.Request) (net.IP, error) {

	// try to get IP from http headers
	if ip, err := getIPFromHeaders(r.Header); err == nil {
		return ip, nil
	}

	// try to get IP from remote address
	if ip, err := getIPFromRemoteAddress(r.RemoteAddr); err == nil {
		return ip, nil
	}

	return nil, fmt.Errorf("Unable to determine IP address from the given request")
}

// getIPFromHeaders tries to extract the remote IP address from the
// given HTTP Header. Returns an error if it cannot locate the
// remote IP address.
func getIPFromHeaders(header http.Header) (net.IP, error) {

	xffHeaderValue, err := getXForwardedForHeader(header)
	if err != nil {
		return nil, fmt.Errorf("No IP found in the given header")
	}

	ip, err := getIPFromHeaderValue(xffHeaderValue)
	if err != nil {
		return nil, err
	}

	return ip, nil
}

// getXForwardedForHeader returns the value of the X-Forwarded-For header
// if there is one; otherwise returns an error.
func getXForwardedForHeader(header http.Header) (string, error) {
	for key, value := range header {
		if strings.ToLower(key) == "x-forwarded-for" && len(value) > 0 {
			return value[0], nil
		}
	}

	return "", fmt.Errorf("X-Forwared-For header was not found")
}

// isPublicIP returns true if the given IP can be routed on the Internet
func isPublicIP(ip net.IP) bool {
	if !ip.IsGlobalUnicast() {
		return false
	}

	for _, mask := range privateNetworkMasks {
		if mask.Contains(ip) {
			return false
		}
	}

	return true
}

// getIPFromHeaderValue parses the value of the X-Forwarded-For Header and returns the IP address.
func getIPFromHeaderValue(ipHeaderValue string) (net.IP, error) {
	for _, entry := range strings.Split(ipHeaderValue, ",") {

		// cleanup the value
		entry = strings.TrimSpace(entry)

		// skip empty values
		if entry == "" {
			continue
		}

		// try to parse the IP
		if ip := net.ParseIP(entry); ip != nil && isPublicIP(ip) {
			return ip, nil
		}
	}

	return nil, fmt.Errorf("Unable to extract IP from %q", ipHeaderValue)
}

// getIPFromRemoteAddress returns the IP address of the given remote address string
// given it is in the format `IP:port`. Otherwise it will return an error.
func getIPFromRemoteAddress(remoteAddress string) (net.IP, error) {

	// don't let garbage in
	if remoteAddress == "" {
		return nil, fmt.Errorf("The given remote address cannot be empty")
	}

	if len(remoteAddress) > 50 {
		return nil, fmt.Errorf("The given remote address is too long to be valid")
	}

	// extract the IP address from the given IP:port string
	ipAddress, _, err := net.SplitHostPort(remoteAddress)
	if err != nil {
		return nil, fmt.Errorf("The given remote address (%q) was not in the expected format of \"IP:port\".", remoteAddress)
	}

	// parse the IP address
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return nil, fmt.Errorf("Failed to parse IP %q", ipAddress)
	}

	return ip, nil
}
