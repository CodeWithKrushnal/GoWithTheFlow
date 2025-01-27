package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"bytes"
)


func TestAddWebsites(t *testing.T) {
	tests := []struct {
		name            string
		body            map[string][]string
		expectedStatus  int
		expectedMessage string
		expectedError   string
	}{
		{
			name: "Valid input",
			body: map[string][]string{
				"websites": {"example.com", "test.com"},
			},
			expectedStatus:  http.StatusAccepted,
			expectedMessage: "Websites Added Successfully",
			expectedError:   "",
		},
		{
			name:            "Missing 'websites' field",
			body:            map[string][]string{},
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: "",
			expectedError:   "Missing 'websites' field in request body",
		},
		{
			name:            "Invalid JSON format",
			body:            nil,
			expectedStatus:  http.StatusBadGateway,
			expectedMessage: "",
			expectedError:   "Error reading request body",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal the body to JSON
			var body []byte
			if tt.body != nil {
				var err error
				body, err = json.Marshal(tt.body)
				if err != nil {
					t.Fatalf("Error marshalling request body: %v", err)
				}
			}

			// Create the POST request
			req, err := http.NewRequest("POST", "/websites", bytes.NewBuffer(body))
			if err != nil {
				t.Fatalf("Error creating request: %v", err)
			}

			// Set content-type header
			req.Header.Set("Content-Type", "application/json")

			// Record the response
			rr := httptest.NewRecorder()

			// Call the handler
			handler := http.HandlerFunc(AddWebsites)
			handler.ServeHTTP(rr, req)

			// Parse the response JSON
			var resp regularResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
				t.Fatalf("Failed to parse response: %v", err)
			}

			// Check if the status code matches the expected value
			if status := resp.StstusCode; status != tt.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", tt.expectedStatus, status)
			}

			// Check if the message matches the expected message
			if tt.expectedMessage != "" && resp.Message != tt.expectedMessage {
				t.Errorf("Expected message %v, but got %v", tt.expectedMessage, resp.Message)
			}

			// Check for error in case of bad requests
			if tt.expectedError != "" && resp.Error != tt.expectedError {
				t.Errorf("Expected error message %v, but got %v", tt.expectedError, resp.Error)
			}
		})
	}
}

func TestWebStatus(t *testing.T) {
	// Set up initial data
	webDirectoryObj.writeDirectory("example.com", "UP")
	webDirectoryObj.writeDirectory("test.com", "DOWN")

	tests := []struct {
		name            string
		queryParam      string
		expectedStatus  int
		expectedMessage string
		expectedError   string
		expectedData    interface{}
	}{
		{
			name:            "Valid query parameter - website exists",
			queryParam:      "example.com",
			expectedStatus:  http.StatusAccepted,
			expectedMessage: "Request Completed Successfully",
			expectedError:   "",
			expectedData:    "example.com is UP",
		},
		{
			name:            "Valid query parameter - website does not exist",
			queryParam:      "nonexistent.com",
			expectedStatus:  http.StatusNotFound,
			expectedMessage: "",
			expectedError:   "Provided key nonexistent.com is not present in the database, kindly call the POST API prior to querying",
			expectedData:    nil,
		},
		{
			name:            "No query parameter and directory is not empty",
			queryParam:      "",
			expectedStatus:  http.StatusAccepted,
			expectedMessage: "Request Completed Successfully",
			expectedError:   "",
			expectedData: map[string]string{
				"example.com": "UP",
				"test.com":    "DOWN",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the GET request with query parameters if provided
			reqURL := "/websites"
			if tt.queryParam != "" {
				reqURL += "?param=" + tt.queryParam
			}
			req, err := http.NewRequest("GET", reqURL, nil)
			if err != nil {
				t.Fatalf("Error creating request: %v", err)
			}

			// Record the response
			rr := httptest.NewRecorder()

			// Call the handler
			handler := http.HandlerFunc(webStatus)
			handler.ServeHTTP(rr, req)

			// Parse the response JSON
			var resp regularResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
				t.Fatalf("Failed to parse response: %v", err)
			}

			// Check if the status code matches the expected value
			if status := resp.StstusCode; status != tt.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", tt.expectedStatus, status)
			}

			// Check if the message matches the expected message
			if tt.expectedMessage != "" && resp.Message != tt.expectedMessage {
				t.Errorf("Expected message %v, but got %v", tt.expectedMessage, resp.Message)
			}

			// Check if the error matches the expected error
			if tt.expectedError != "" && resp.Error != tt.expectedError {
				t.Errorf("Expected error message %v, but got %v", tt.expectedError, resp.Error)
			}

			// Check if the data matches the expected data
			if tt.expectedData != nil {
				expectedDataBytes, _ := json.Marshal(tt.expectedData)
				resultBytes, _ := json.Marshal(resp.Data)
				if !reflect.DeepEqual(expectedDataBytes, resultBytes) {
					t.Errorf("Expected data %v, but got %v", tt.expectedData, resp.Data)
				}
			} else if resp.Data != nil {
				t.Errorf("Expected no data, but got %v", resp.Data)
			}
		})
	}
}
