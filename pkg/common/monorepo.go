package common

import (
	"os"
	"sort"
)

// Return a list of service names
func ServiceNames(name string) ([]string, error) {
	// List the contents of the services directory
	files, err := os.ReadDir(name)
	if err != nil {
		return nil, err
	}

	var services []string
	for _, file := range files {
		if file.IsDir() {
			services = append(services, file.Name())
		}
	}

	// Sort the service names
	sort.Strings(services)

	return services, nil
}
