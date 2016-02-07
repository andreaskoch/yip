// Copyright 2016 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net"
	"net/http"
	"strings"
	"testing"
)

func Test_getIPFromRequest_emptyRequest_ErrorIsReturned(t *testing.T) {
	// arrange
	request := &http.Request{}

	// act
	_, err := getIPFromRequest(request)

	// assert
	if err == nil || !strings.Contains(err.Error(), "Unable to determine IP address from the given request") {
		t.Fail()
		t.Logf("getIPFromRequest(%#v) should return an error", request)
	}

}

func Test_getIPFromRequest_OnlyRemoteAdressIsSet_RemoteAddressIsReturned(t *testing.T) {
	// arrange
	request := &http.Request{
		RemoteAddr: "[::1]:34576",
	}

	// act
	ip, _ := getIPFromRequest(request)

	// assert
	expected := net.ParseIP("::1")
	if ip.String() != expected.String() {
		t.Fail()
		t.Logf("getIPFromRequest(%#v) should return %q but returned %q instead", request, expected, ip)
	}

}

func Test_getIPFromRequest_XffAndRemoteAdressAreSet_XffIPIsReturned(t *testing.T) {
	// arrange
	request := &http.Request{
		RemoteAddr: "[::3]:34576",
		Header: map[string][]string{
			"Accept-Encoding":   []string{"gzip, deflate"},
			"If-Modified-Since": []string{"Sun, 7 Feb 2016 10:13:41 UTC"},
			"X-Forwarded-For":   []string{"8.8.8.8"},
		},
	}

	// act
	ip, _ := getIPFromRequest(request)

	// assert
	expected := net.ParseIP("8.8.8.8")
	if ip.String() != expected.String() {
		t.Fail()
		t.Logf("getIPFromRequest(%#v) should return %q but returned %q instead", request, expected, ip)
	}

}
