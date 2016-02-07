// Copyright 2016 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_ipResponder_EmptyRequest_StatusCode400(t *testing.T) {
	// arrange
	responseWriter := httptest.NewRecorder()
	request := &http.Request{}

	// act
	ipResponder(responseWriter, request)

	// assert
	if responseWriter.Code != 400 {
		t.Fail()
		t.Logf("ipResponder() should not return the status code %d", responseWriter.Code)
	}
}

func Test_ipResponder_EmptyRequest_ErrorIsWrittenToResponse(t *testing.T) {
	// arrange
	responseWriter := httptest.NewRecorder()
	request := &http.Request{}

	// act
	ipResponder(responseWriter, request)

	// assert
	expectedResponse := "Unable to determine remote address from request"
	if !strings.Contains(responseWriter.Body.String(), expectedResponse) {
		t.Fail()
		t.Logf("ipResponder() wrote %q to the response", responseWriter.Body)
	}
}

func Test_ipResponder_RemoteAddressIsSet_StatusCode200(t *testing.T) {
	// arrange
	responseWriter := httptest.NewRecorder()
	request := &http.Request{
		RemoteAddr: "8.8.8.8:46484",
	}

	// act
	ipResponder(responseWriter, request)

	// assert
	if responseWriter.Code != 200 {
		t.Fail()
		t.Logf("ipResponder() should not return the status code %d", responseWriter.Code)
	}
}

func Test_ipResponder_RemoteAddressIsSet_RemoteAddressIsWrittenToResponse(t *testing.T) {
	// arrange
	responseWriter := httptest.NewRecorder()
	request := &http.Request{
		RemoteAddr: "8.8.8.8:46484",
	}

	// act
	ipResponder(responseWriter, request)

	// assert
	expectedResponse := "8.8.8.8"
	if !strings.Contains(responseWriter.Body.String(), expectedResponse) {
		t.Fail()
		t.Logf("ipResponder() wrote %q to the response", responseWriter.Body)
	}
}

func Test_ipResponder_XffHeaderIsSet_StatusCode200(t *testing.T) {
	// arrange
	responseWriter := httptest.NewRecorder()
	request := &http.Request{
		RemoteAddr: "8.8.8.8:46484",
		Header: map[string][]string{
			"Accept-Encoding":   []string{"gzip, deflate"},
			"If-Modified-Since": []string{"Sun, 7 Feb 2016 10:13:41 UTC"},
			"X-Forwarded-For":   []string{"4.4.4.4"},
		},
	}

	// act
	ipResponder(responseWriter, request)

	// assert
	if responseWriter.Code != 200 {
		t.Fail()
		t.Logf("ipResponder() should not return the status code %d", responseWriter.Code)
	}
}

func Test_ipResponder_XffHeaderIsSet_XffAddressIsWrittenToResponse(t *testing.T) {
	// arrange
	responseWriter := httptest.NewRecorder()
	request := &http.Request{
		RemoteAddr: "8.8.8.8:46484",
		Header: map[string][]string{
			"Accept-Encoding":   []string{"gzip, deflate"},
			"If-Modified-Since": []string{"Sun, 7 Feb 2016 10:13:41 UTC"},
			"X-Forwarded-For":   []string{"4.4.4.4"},
		},
	}

	// act
	ipResponder(responseWriter, request)

	// assert
	expectedResponse := "4.4.4.4"
	if !strings.Contains(responseWriter.Body.String(), expectedResponse) {
		t.Fail()
		t.Logf("ipResponder() wrote %q to the response", responseWriter.Body)
	}
}
