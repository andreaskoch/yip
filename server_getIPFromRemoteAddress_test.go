// Copyright 2016 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"strings"
	"testing"
)

func Test_getIPFromRemoteAddress_RemoteAddressStringIsEmpty_ErrorIsReturned(t *testing.T) {
	// arrange
	remoteAddress := ""

	// act
	_, err := getIPFromRemoteAddress(remoteAddress)

	// assert
	if err == nil || !strings.Contains(err.Error(), "address cannot be empty") {
		t.Fail()
		t.Logf("getIPFromRemoteAddress(%q) should return an error because the given address is empty.", remoteAddress)
	}
}

func Test_getIPFromRemoteAddress_RemoteAddressStringIsTooLong_ErrorIsReturned(t *testing.T) {
	// arrange
	remoteAddress := "AbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyz"

	// act
	_, err := getIPFromRemoteAddress(remoteAddress)

	// assert
	if err == nil || !strings.Contains(err.Error(), "address is too long") {
		t.Fail()
		t.Logf("getIPFromRemoteAddress(%q) should return an error because the given address is too long to be an IP address.", remoteAddress)
	}
}

func Test_getIPFromRemoteAddress_InvalidIPPortFormat_ErrorIsReturned(t *testing.T) {
	// arrange
	remoteAddresses := []string{
		"127.0.0.1",
		"::1",
		"::1:64000",
		"[::1]",
		"[::1]::32",
	}
	for _, remoteAddress := range remoteAddresses {

		// act
		_, err := getIPFromRemoteAddress(remoteAddress)

		// assert
		if err == nil || !strings.Contains(err.Error(), "not in the expected format") {
			t.Fail()
			t.Logf("getIPFromRemoteAddress(%q) should return an error because the given address does not have the correct IP:port pattern.", remoteAddress)
		}
	}

}

func Test_getIPFromRemoteAddress_IPCannotBeParsed_ErrorIsReturned(t *testing.T) {
	// arrange
	remoteAddresses := []string{
		"127.0.0.1.1:80",
		"0.1.1:80",
		"[::100000]:80",
		"sadasdsa:80",
		"example.com:443",
	}
	for _, remoteAddress := range remoteAddresses {

		// act
		_, err := getIPFromRemoteAddress(remoteAddress)

		// assert
		if err == nil || !strings.Contains(err.Error(), "Failed to parse IP") {
			t.Fail()
			t.Logf("getIPFromRemoteAddress(%q) should return an error because the given address does not have the correct IP:port pattern.", remoteAddress)
		}
	}

}

func Test_getIPFromRemoteAddress_ValidRemoteAddress_IPIsReturned(t *testing.T) {
	// arrange
	remoteAddresses := []string{
		"127.0.0.1:80",
		"[::1]:80",
	}
	for _, remoteAddress := range remoteAddresses {

		// act
		ip, err := getIPFromRemoteAddress(remoteAddress)

		// assert
		if err != nil || ip == nil {
			t.Fail()
			t.Logf("getIPFromRemoteAddress(%q) should not return an error because the given address is valid.", remoteAddress)
		}
	}

}
