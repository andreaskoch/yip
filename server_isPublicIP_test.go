// Copyright 2016 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net"
	"testing"
)

func Test_isPublicIP_PrivateIPIsGiven_ResultIsFalse(t *testing.T) {
	// arrange
	IPs := []net.IP{
		net.ParseIP("127.0.0.1"),
		net.ParseIP("127.1.1.1"),
		net.ParseIP("10.0.0.45"),
		net.ParseIP("10.2.0.45"),
		net.ParseIP("172.16.1.32"),
		net.ParseIP("172.17.1.32"),
		net.ParseIP("192.168.13.254"),
		net.ParseIP("fd00::1"),
		net.ParseIP("fd00::2000:1"),
	}

	for _, ip := range IPs {

		// act
		result := isPublicIP(ip)

		// assert
		if result == true {
			t.Fail()
			t.Logf("isPublicIP(%q) should return false.", ip)
		}
	}
}

func Test_isPublicIP_PublicIPIsGiven_ResultIsTrue(t *testing.T) {
	// arrange
	IPs := []net.IP{
		net.ParseIP("173.2.1.1"),
		net.ParseIP("8.8.8.8"),
		net.ParseIP("2001:0db8:0000:0042:0000:8a2e:0370:7334"),
	}

	for _, ip := range IPs {

		// act
		result := isPublicIP(ip)

		// assert
		if result == false {
			t.Fail()
			t.Logf("isPublicIP(%q) should return true.", ip)
		}
	}
}
