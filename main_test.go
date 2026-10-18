package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	handler(rec, req)

	if !strings.Contains(rec.Body.String(), "Hello from Go Cloud API!") {
		t.Fatalf("unexpected response: %s", rec.Body.String())
	}
}
