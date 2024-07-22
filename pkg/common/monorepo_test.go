package common

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestServiceNames(t *testing.T) {
	var (
		err error
	)

	t.Skip("Skipping TestServiceNames for now")

	// Setup: Create a temporary directory to simulate 'services'
	tempDir := os.TempDir()

	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Fatalf("Failed to remove temp dir: %v", err)
		}
	}(tempDir) // Cleanup

	// Create dummy directories to simulate services
	expectedServices := []string{"service1", "service2"}
	for _, serviceName := range expectedServices {
		err := os.Mkdir(filepath.Join(tempDir, serviceName), 0755)
		if err != nil {
			t.Fatalf("Failed to create service directory: %v", err)
		}
	}

	// Create a dummy file to test that only directories are listed
	_, err = os.Create(filepath.Join(tempDir, "file.txt"))
	if err != nil {
		t.Fatalf("Failed to create dummy file: %v", err)
	}

	// Override the services directory path to point to our temp directory
	// servicesDir := filepath.Join(tempDir) // Normally, this would be 'services'
	// Test
	gotServices, err := ServiceNames(tempDir)
	if err != nil {
		t.Errorf("ServiceNames() error = %v, wantErr false", err)
	}

	if !reflect.DeepEqual(gotServices, expectedServices) {
		t.Errorf("ServiceNames() = %v, want %v", gotServices, expectedServices)
	}
}
