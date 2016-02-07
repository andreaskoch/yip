// Copyright 2016 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"strings"
	"testing"
)

func Test_getXForwardedForHeader_EmptyHeader_ErrorIsReturned(t *testing.T) {
	// arrange
	header := map[string][]string{}

	// act
	_, err := getXForwardedForHeader(header)

	// assert
	if err == nil || !strings.Contains(err.Error(), "X-Forwared-For header was not found") {
		t.Fail()
		t.Logf("getXForwardedForHeader(%q) should return an error if no XFF header was found.", header)
	}

}

func Test_getXForwardedForHeader_HeaderWithoutXff_ErrorIsReturned(t *testing.T) {
	// arrange
	header := map[string][]string{
		"Accept-Encoding":   []string{"gzip, deflate"},
		"If-Modified-Since": []string{"Sun, 7 Feb 2016 10:13:41 UTC"},
	}

	// act
	_, err := getXForwardedForHeader(header)

	// assert
	if err == nil || !strings.Contains(err.Error(), "X-Forwared-For header was not found") {
		t.Fail()
		t.Logf("getXForwardedForHeader(%q) should return an error if no XFF header was found.", header)
	}

}

func Test_getXForwardedForHeader_HeaderWithSingleXff_XffIsReturned(t *testing.T) {
	// arrange
	header := map[string][]string{
		"Accept-Encoding":   []string{"gzip, deflate"},
		"If-Modified-Since": []string{"Sun, 7 Feb 2016 10:13:41 UTC"},
		"X-Forwarded-For":   []string{"::1"},
	}

	// act
	result, _ := getXForwardedForHeader(header)

	// assert
	if result != "::1" {
		t.Fail()
		t.Logf("getXForwardedForHeader(%q) should return %s.", header, "::1")
	}

}

func Test_getXForwardedForHeader_HeaderWithMultipleXff_FirstXffIsReturned(t *testing.T) {
	// arrange
	header := map[string][]string{
		"Accept-Encoding":   []string{"gzip, deflate"},
		"If-Modified-Since": []string{"Sun, 7 Feb 2016 10:13:41 UTC"},
		"X-Forwarded-For":   []string{"::1", "::2", "::3"},
	}

	// act
	result, _ := getXForwardedForHeader(header)

	// assert
	if result != "::1" {
		t.Fail()
		t.Logf("getXForwardedForHeader(%q) should return %s.", header, "::1")
	}

}

func Test_getXForwardedForHeader_XffHeaderWithDifferentCasing_XffIsReturned(t *testing.T) {
	// arrange
	headerNames := []string{
		"X-Forwarded-For",
		"x-forwarded-for",
		"X-FORWARDED-FOR",
	}

	for _, headerName := range headerNames {

		header := map[string][]string{
			headerName:          []string{"::1", "::2", "::3"},
			"Accept-Encoding":   []string{"gzip, deflate"},
			"If-Modified-Since": []string{"Sun, 7 Feb 2016 10:13:41 UTC"},
		}

		// act
		result, _ := getXForwardedForHeader(header)

		// assert
		if result != "::1" {
			t.Fail()
			t.Logf("getXForwardedForHeader(%q) should return %s.", header, "::1")
		}

	}
}
