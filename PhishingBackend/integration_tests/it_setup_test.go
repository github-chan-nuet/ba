//go:build integration

package integration_tests

import (
	"net/http/httptest"
	"phishing_backend/internal/adapters"
	"phishing_backend/internal/adapters/presentation"
	"testing"
)

var ts *httptest.Server

func TestMain(m *testing.M) {
	// Setup code
	d := adapters.ResolveDependencies()
	ts = httptest.NewServer(presentation.NewHttpHandler(d))
	defer ts.Close()

	// Run the tests
	m.Run()

	// Teardown code
}
