package tds_college_conference_service

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	steps "github.com/ocrosby/soccer/cmd/tds-college-conference-service/features/step_definitions"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"testing"
)

// Global variable to hold the API container instance
var apiContainer testcontainers.Container

func StartAPIContainer() (testcontainers.Container, error) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "your-api-image:latest",                      // Replace with your actual API image
		ExposedPorts: []string{"8080/tcp"},                         // Adjust the port to match your API's port
		WaitingFor:   wait.ForHTTP("/health").WithPort("8080/tcp"), // Adjust the endpoint to your API's health check endpoint

	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	return container, nil
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		var err error

		// Start the API container before the test suite
		apiContainer, err = StartAPIContainer()
		if err != nil {
			panic(fmt.Errorf("Could not start API container: %v", err))
		}
	})

	ctx.AfterSuite(func() {
		// Terminate the API container after the test suite
		if err := apiContainer.Terminate(context.Background()); err != nil {
			panic(fmt.Errorf("Could not terminate API container: %v", err))
		}
	})
}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format:    "pretty",               // Change to any format you prefer
		Paths:     []string{"./features"}, // Path to your feature files
		Randomize: 0,                      // Optional: specify a seed to randomize the scenario execution order
	}

	status := godog.TestSuite{
		Name:                 "TopDrawerSoccer College Conference Service",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  steps.InitializeScenario,
		Options:              &opts,
	}.Run()

	// Optional: combine with unit test status
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
