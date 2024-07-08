package main

// detect_changed_services.go is a utility script designed for use in a CI/CD pipeline within a Go monorepo structure.
// Its primary function is to identify which services under the `cmd` directory have changed by comparing the current
// branch with the main branch. This allows for optimized CI/CD processes by ensuring that only the services affected
// by recent changes are built and tested. This script utilizes git commands to fetch the latest changes, determine the
// merge base, and list changed files, ultimately extracting and deduplicating service names to identify the changed
// services.

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func fetchLatestChanges() error {
	cmd := exec.Command("git", "fetch", "origin", "main")
	return cmd.Run()
}

func getMergeBase() (string, error) {
	cmd := exec.Command("git", "merge-base", "FETCH_HEAD", "main")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return strings.TrimSpace(out.String()), err
}

func getChangedServices(mergeBase string) ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-only", "FETCH_HEAD", mergeBase, "--", "cmd/*")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	changedFiles := strings.Split(strings.TrimSpace(out.String()), "\n")
	serviceMap := make(map[string]bool)
	for _, file := range changedFiles {
		parts := strings.Split(file, "/")
		if len(parts) > 1 {
			serviceMap[parts[1]] = true
		}
	}

	var services []string
	for service := range serviceMap {
		services = append(services, service)
	}

	return services, nil
}

func main() {
	if err := fetchLatestChanges(); err != nil {
		fmt.Println("Error fetching latest changes:", err)
		return
	}

	mergeBase, err := getMergeBase()
	if err != nil {
		fmt.Println("Error getting merge base:", err)
		return
	}

	services, err := getChangedServices(mergeBase)
	if err != nil {
		fmt.Println("Error getting changed services:", err)
		return
	}

	fmt.Println("Changed services:", strings.Join(services, ", "))
}
