// Copyright 2016 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"strings"
	"testing"
)

func Test_getIPFromHeaderValue_InvalidValues_ErrorIsReturned(t *testing.T) {
	// arrange
	ipHeaderValues := []string{
		"",
		" ",
		",",
		" , ",
		"127.555.1.1",
		"::90000",
	}

	for _, ipHeaderValue := range ipHeaderValues {

		// act
		_, err := getIPFromHeaderValue(ipHeaderValue)

		// assert
		if err == nil || !strings.Contains(err.Error(), "Unable to extract IP from") {
			t.Fail()
			t.Logf("getIPFromHeaderValue(%q) should return an error.", ipHeaderValue)
		}
	}
}

func Test_getIPFromHeaderValue_PrivateAddresses_ErrorIsReturned(t *testing.T) {
	// arrange
	ipHeaderValues := []string{
		"10.0.2.1",
		"127.0.0.1",
		"127.0.0.1,127.0.0.4,127.0.10.1",
		"172.16.54.4",
		"172.17.54.4",
		"192.168.1.3",
		"192.168.2.3",
		"fd00::1",
	}

	for _, ipHeaderValue := range ipHeaderValues {

		// act
		_, err := getIPFromHeaderValue(ipHeaderValue)

		// assert
		if err == nil || !strings.Contains(err.Error(), "Unable to extract IP from") {
			t.Fail()
			t.Logf("getIPFromHeaderValue(%q) should return an error.", ipHeaderValue)
		}
	}
}

func Test_getIPFromHeaderValue_PublicAddresses_IPIsReturned(t *testing.T) {
	// arrange
	ipHeaderValues := []string{
		"54.56.78.9",
		"8.8.8.8",
	}

	for _, ipHeaderValue := range ipHeaderValues {

		// act
		ip, err := getIPFromHeaderValue(ipHeaderValue)

		// assert
		if err != nil || ip == nil {
			t.Fail()
			t.Logf("getIPFromHeaderValue(%q) should return an IP address.", ipHeaderValue)
		}
	}
}

func Test_getIPFromHeaderValue_MixedPublicAndPrivateAddresses_PublicIPIsReturned(t *testing.T) {
	// arrange
	ipHeaderValue := "10.2.1.3 , 8.8.8.8"

	// act
	ip, _ := getIPFromHeaderValue(ipHeaderValue)

	// assert
	if ip.String() != "8.8.8.8" {
		t.Fail()
		t.Logf("getIPFromHeaderValue(%q) should return the public IP address.", ipHeaderValue)
	}
}
