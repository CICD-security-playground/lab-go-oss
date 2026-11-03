package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	appserver "github.com/CICD-security-playground/lab-go-oss/internal/server"
)

func TestHealthEndpoint(t *testing.T) {
	response := performRequest(t, "/health")

	if got, want := response.Code, http.StatusOK; got != want {
		t.Fatalf("status code = %d, want %d", got, want)
	}

	body := decodeResponse(t, response)
	if got, want := body["status"], "healthy"; got != want {
		t.Errorf("status = %q, want %q", got, want)
	}
}

func TestHelloEndpoint(t *testing.T) {
	response := performRequest(t, "/hello")

	if got, want := response.Code, http.StatusOK; got != want {
		t.Fatalf("status code = %d, want %d", got, want)
	}

	body := decodeResponse(t, response)
	if got := body["message"]; !strings.Contains(got, "Hello") {
		t.Errorf("message = %q, want a greeting containing %q", got, "Hello")
	}
}

func performRequest(t *testing.T, path string) *httptest.ResponseRecorder {
	t.Helper()

	request := httptest.NewRequest(http.MethodGet, path, nil)
	response := httptest.NewRecorder()
	appserver.NewHandler().ServeHTTP(response, request)

	if got, want := response.Header().Get("Content-Type"), "application/json; charset=utf-8"; got != want {
		t.Errorf("Content-Type = %q, want %q", got, want)
	}

	return response
}

func decodeResponse(t *testing.T, response *httptest.ResponseRecorder) map[string]string {
	t.Helper()

	var body map[string]string
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	return body
}
