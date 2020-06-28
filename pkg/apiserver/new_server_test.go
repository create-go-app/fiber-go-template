package apiserver

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIRoutes(t *testing.T) {
	// Define a structure for specifying input and output
	// data of a single test case. This structure is then used
	// to create a so called test map, which contains all test
	// cases, that should be run for testing this function
	tests1 := []struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "API routes",
			route:         "/api/posts",
			expectedError: false,
			expectedCode:  200,
		},
		{
			description:   "Static route (with prefix)",
			route:         "/public/index.html",
			expectedError: false,
			expectedCode:  200,
		},
	}

	tests2 := []struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "Static route (without prefix)",
			route:         "/index.html",
			expectedError: false,
			expectedCode:  200,
		},
	}

	// Start the app as it is done in the main function
	app1 := NewServer(&Config{
		Static: static{Prefix: "/public", Path: "../../static"},
	}).Start()

	app2 := NewServer(&Config{
		Static: static{Prefix: "", Path: "../../static"},
	}).Start()

	// Iterate through test single test cases
	for _, test := range tests1 {
		// Create a new http request with the route
		// from the test case
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res1, err1 := app1.Test(req, -1)

		// verify that no error occurred, that is not expected
		assert.Equalf(t, test.expectedError, err1 != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res1.StatusCode, test.description)
	}

	for _, test := range tests2 {
		// Create a new http request with the route
		// from the test case
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res2, err2 := app2.Test(req, -1)

		// verify that no error occurred, that is not expected
		assert.Equalf(t, test.expectedError, err2 != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res2.StatusCode, test.description)
	}
}
