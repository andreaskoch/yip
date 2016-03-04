// Copyright 2016 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net"
	"testing"
)

func Test_getIPFromHeaders_NoHeader_ErrorIsReturned(t *testing.T) {

	// act
	_, err := getIPFromHeaders(nil)

	// assert
	if err == nil {
		t.Fail()
		t.Logf("getIPFromHeaders(%q) should return an error because the given header is empty", "nil")
	}
}

func Test_getIPFromHeaders_EmptyHeader_ErrorIsReturned(t *testing.T) {
	// arrange
	header := map[string][]string{}

	// act
	_, err := getIPFromHeaders(header)

	// assert
	if err == nil {
		t.Fail()
		t.Logf("getIPFromHeaders(%q) should return an error because the given header is empty", header)
	}
}

func Test_getIPFromHeaders_NoXffHeader_ErrorIsReturned(t *testing.T) {
	// arrange
	header := map[string][]string{
		"Accept-Encoding":   []string{"gzip, deflate"},
		"If-Modified-Since": []string{"Sun, 7 Feb 2016 10:13:41 UTC"},
	}

	// act
	_, err := getIPFromHeaders(header)

	// assert
	if err == nil {
		t.Fail()
		t.Logf("getIPFromHeaders(%q) should return an error because there is no IP in the given header", header)
	}
}

func Test_getIPFromHeaders_XffHeaderWithMultiplePublicIPsGiven_FirstPublicIPIsReturned(t *testing.T) {
	// arrange
	header := map[string][]string{
		"Accept-Encoding":   []string{"gzip, deflate"},
		"If-Modified-Since": []string{"Sun, 7 Feb 2016 10:13:41 UTC"},
		"X-Forwarded-For":   []string{"2001:0db8::0042::8a2e:0370:1,2001:0db8::0042::8a2e:0370:2,2001:0db8::0042::8a2e:0370:3"},
	}

	// act
	ip, _ := getIPFromHeaders(header)

	// assert
	expected := net.ParseIP("2001:0db8::0042::8a2e:0370:1").String()
	if ip.String() != expected {
		t.Fail()
		t.Logf("getIPFromHeaders(%q) should return %s", header, expected)
	}
}

func Test_getIPFromHeaders_XffHeaderWithDifferentCasing_PublicIPIsReturned(t *testing.T) {
	// arrange
	headerNames := []string{
		"X-Forwarded-For",
		"x-forwarded-for",
		"X-FORWARDED-FOR",
	}

	for _, headerName := range headerNames {

		header := map[string][]string{
			"Accept-Encoding":   []string{"gzip, deflate"},
			"If-Modified-Since": []string{"Sun, 7 Feb 2016 10:13:41 UTC"},
			headerName:          []string{"2001:0db8:0000:0042:0000:8a2e:0370:7334"},
		}

		// act
		ip, _ := getIPFromHeaders(header)

		// assert
		expected := net.ParseIP("2001:0db8:0000:0042:0000:8a2e:0370:7334").String()
		if ip.String() != expected {
			t.Fail()
			t.Logf("getIPFromHeaders(%q) should return %s", header, expected)
		}
	}
}
