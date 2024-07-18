package internal

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestEncode verifies the Encode function
func TestEncode(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/dummy-url", nil) // Create a dummy request
	user := struct {
		Name string `json:"name"`
	}{"John Doe"}

	err := Encode(w, req, http.StatusOK, user)
	if err != nil {
		t.Errorf("Encode returned an error: %v", err)
	}

	if w.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", w.Code)
	}

	expected := `{"name":"John Doe"}`
	if !strings.Contains(w.Body.String(), expected) {
		t.Errorf("expected body to contain %v; got %v", expected, w.Body.String())
	}
}

// TestDecode verifies the Decode function
func TestDecode(t *testing.T) {
	jsonBody := `{"name":"Jane Doe"}`
	r := httptest.NewRequest("POST", "/", bytes.NewBufferString(jsonBody))

	var user struct {
		Name string `json:"name"`
	}

	result, err := Decode[struct{ Name string }](r)
	if err != nil {
		t.Errorf("Decode returned an error: %v", err)
	}

	if result.Name != "Jane Doe" {
		t.Errorf("expected name to be Jane Doe; got %v", user.Name)
	}
}
