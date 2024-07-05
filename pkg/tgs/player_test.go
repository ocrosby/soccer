// Author: Omar Crosby
// File: player_test.go

package tgs

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetAllStates(t *testing.T) {
	// Mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		states := []State{
			{ID: 1, Name: "California"},
			{ID: 2, Name: "New York"},
		}

		if err := json.NewEncoder(w).Encode(states); err != nil {
			return
		}
	}))
	defer mockServer.Close()

	// Expected states
	expectedStates := []State{
		{ID: 1, Name: "California"},
		{ID: 2, Name: "New York"},
	}

	// Test GetAllStates
	states, err := GetAllStates(mockServer.URL)
	if err != nil {
		t.Fatalf("GetAllStates failed: %v", err)
	}

	// Assert
	if !reflect.DeepEqual(states, expectedStates) {
		t.Errorf("Expected %v, got %v", expectedStates, states)
	}
}
