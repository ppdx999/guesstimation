package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		path             string
		expectedStatus   int
		expectedResponse string
	}{
		{path: "/hello", expectedStatus: http.StatusOK, expectedResponse: "Hello, World!"},
		{path: "/path1", expectedStatus: http.StatusOK, expectedResponse: "Response from /path1"},
		{path: "/path2", expectedStatus: http.StatusOK, expectedResponse: "Response from /path2"},
		{path: "/unknown", expectedStatus: http.StatusNotFound, expectedResponse: "Not Found"},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			request := events.APIGatewayProxyRequest{Path: tt.path}
			response, err := handler(context.Background(), request)

			if err != nil {
				t.Fatalf("handler returned an error: %v", err)
			}

			if response.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, response.StatusCode)
			}

			if response.Body != tt.expectedResponse {
				t.Errorf("expected body %q, got %q", tt.expectedResponse, response.Body)
			}
		})
	}
}
