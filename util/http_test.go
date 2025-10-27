package util

import (
	"net/http"
	"testing"
)

// TestHttpWithTLS_NilTransport tests the scenario where DefaultHTTPClient.Transport is nil
func TestHttpWithTLS_NilTransport(t *testing.T) {
	// Save original transport
	originalTransport := DefaultHTTPClient.Transport
	defer func() {
		DefaultHTTPClient.Transport = originalTransport
	}()

	// Set Transport to nil to simulate the bug scenario
	DefaultHTTPClient.Transport = nil

	// This should not panic after the fix
	// Note: This will fail due to invalid cert path, but shouldn't panic on type assertion
	_, err := httpWithTLS("./testdata/invalid_cert.p12", "password")

	// We expect an error (cert file not found), but NOT a panic
	if err == nil {
		t.Error("Expected error due to invalid cert path, but got nil")
	}
}

// TestHttpWithTLS_CustomTransport tests the scenario where DefaultHTTPClient has a custom Transport
func TestHttpWithTLS_CustomTransport(t *testing.T) {
	// Save original transport
	originalTransport := DefaultHTTPClient.Transport
	defer func() {
		DefaultHTTPClient.Transport = originalTransport
	}()

	// Set a custom http.Transport
	customTransport := &http.Transport{
		MaxIdleConns: 100,
	}
	DefaultHTTPClient.Transport = customTransport

	// This should not panic
	_, err := httpWithTLS("./testdata/invalid_cert.p12", "password")

	// We expect an error (cert file not found), but NOT a panic
	if err == nil {
		t.Error("Expected error due to invalid cert path, but got nil")
	}
}

// CustomRoundTripper is a custom implementation of http.RoundTripper
type CustomRoundTripper struct{}

func (c *CustomRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return http.DefaultTransport.RoundTrip(req)
}

// TestHttpWithTLS_CustomRoundTripper tests the edge case where DefaultHTTPClient has a custom RoundTripper
// that is NOT *http.Transport
func TestHttpWithTLS_CustomRoundTripper(t *testing.T) {
	// Save original transport
	originalTransport := DefaultHTTPClient.Transport
	defer func() {
		DefaultHTTPClient.Transport = originalTransport
	}()

	// Set a custom RoundTripper that is NOT *http.Transport
	customRoundTripper := &CustomRoundTripper{}
	DefaultHTTPClient.Transport = customRoundTripper

	// Create a recovery handler to catch potential panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("httpWithTLS panicked with custom RoundTripper: %v", r)
		}
	}()

	// This might panic if the code doesn't handle non-*http.Transport RoundTripper properly
	_, _ = httpWithTLS("./testdata/invalid_cert.p12", "password")
}
