package config

import "testing"

func TestGetBaseURL(t *testing.T) {
	cfg := &Config{}
	if cfg.GetBaseURL() != defaultBaseURL {
		t.Fatalf("bad default base url: %s", cfg.GetBaseURL())
	}

	cfg.BaseURL = "http://example.com"
	if cfg.GetBaseURL() != "http://example.com" {
		t.Fatalf("bad custom base url: %s", cfg.GetBaseURL())
	}
}
