package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResolvePort(t *testing.T) {
	t.Run("uses default", func(t *testing.T) {
		t.Setenv("PORT", "")

		port, err := resolvePort()
		if err != nil {
			t.Fatalf("resolvePort() error = %v", err)
		}
		if port != defaultPort {
			t.Fatalf("resolvePort() = %d, want %d", port, defaultPort)
		}
	})

	t.Run("uses configured port", func(t *testing.T) {
		t.Setenv("PORT", "9090")

		port, err := resolvePort()
		if err != nil {
			t.Fatalf("resolvePort() error = %v", err)
		}
		if port != 9090 {
			t.Fatalf("resolvePort() = %d, want 9090", port)
		}
	})

	t.Run("rejects invalid port", func(t *testing.T) {
		t.Setenv("PORT", "70000")

		if _, err := resolvePort(); err == nil {
			t.Fatal("resolvePort() error = nil, want an error")
		}
	})
}

func TestReadinessHandler(t *testing.T) {
	t.Run("reports ready", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/readiness", nil)
		response := httptest.NewRecorder()

		readinessHandler(response, request)

		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
		}
		if response.Body.String() != `{"status":"ready"}` {
			t.Fatalf("body = %q, want readiness JSON", response.Body.String())
		}
	})

	t.Run("rejects other methods", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/readiness", nil)
		response := httptest.NewRecorder()

		readinessHandler(response, request)

		if response.Code != http.StatusMethodNotAllowed {
			t.Fatalf("status = %d, want %d", response.Code, http.StatusMethodNotAllowed)
		}
		if response.Header().Get("Allow") != http.MethodGet {
			t.Fatalf("Allow header = %q, want %q", response.Header().Get("Allow"), http.MethodGet)
		}
	})
}
