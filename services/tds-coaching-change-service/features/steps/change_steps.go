package steps

import (
	"context"
	"github.com/cucumber/godog"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"time"
)

var coachingChangeServiceContainer testcontainers.Container

func allChangeResultsShouldBeFemale() error {
	return godog.ErrPending
}

func allChangeResultsShouldBeMale() error {
	return godog.ErrPending
}

func iRequestAllCoachingChanges() error {
	return godog.ErrPending
}

func iRequestFemaleCoachingChanges() error {
	return godog.ErrPending
}

func iRequestMaleCoachingChanges() error {
	return godog.ErrPending
}

func theResponseShouldBeSuccessful() error {
	return godog.ErrPending
}

func thereShouldBeBothMaleAndFemaleChangeResults() error {
	return godog.ErrPending
}

func thereShouldBeNoErrors() error {
	return godog.ErrPending
}

func startCoachingChangeServiceContainer() (testcontainers.Container, nat.Port, error) {
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       "./../../../",
			Dockerfile:    "./services/tds-coaching-change-service/Dockerfile",
			PrintBuildLog: true,
		},
		ExposedPorts: []string{"8080"},
		WaitingFor:   wait.ForHTTP("/ready").WithStartupTimeout(10 * time.Second),
	}

	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		log.Printf("Failed to start Coaching Change Service container: %v", err)
		return nil, "", err
	}

	port, err := container.MappedPort(context.Background(), "8080")
	if err != nil {
		log.Printf("Failed to get mapped port: %v", err)
		return nil, "", err
	}

	return container, port, nil
}

func handleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func logContainerDetails(port nat.Port, containerName string) {
	log.Printf("Container %s is running on port %s", containerName, port.Port())
}

func InitializeCoachingChangeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		// This function will run before the test suite is executed
		var (
			err  error
			port nat.Port
		)

		coachingChangeServiceContainer, port, err = startCoachingChangeServiceContainer()
		handleError(err, "Failed to start Coaching Change Service container")
		logContainerDetails(port, "Coaching Change Service")
	})

	ctx.AfterSuite(func() {
		// This function will run after the test suite is executed
		if err := coachingChangeServiceContainer.Terminate(context.Background()); err != nil {
			log.Fatalf("Failed to stop Coaching Change Service container: %v", err)
		}
	})
}

func InitializeCoachingChangeScenario(ctx *godog.ScenarioContext) {
	ctx.StepContext().Before(func(ctx context.Context, st *godog.Step) (context.Context, error) {
		// This function will run before each scenario is executed
		return nil, nil
	})

	ctx.StepContext().After(func(ctx context.Context, st *godog.Step, status godog.StepResultStatus, err error) (context.Context, error) {
		// This function will run after each scenario is executed
		return nil, nil
	})

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		// This function will run before each scenario is executed
		return nil, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		// Cleanup code after each scenario
		return nil, nil
	})

	ctx.Step(`^all change results should be female$`, allChangeResultsShouldBeFemale)
	ctx.Step(`^all change results should be male$`, allChangeResultsShouldBeMale)
	ctx.Step(`^I request all coaching changes$`, iRequestAllCoachingChanges)
	ctx.Step(`^I request female coaching changes$`, iRequestFemaleCoachingChanges)
	ctx.Step(`^I request male coaching changes$`, iRequestMaleCoachingChanges)
	ctx.Step(`^the response should be successful$`, theResponseShouldBeSuccessful)
	ctx.Step(`^there should be both male and female change results$`, thereShouldBeBothMaleAndFemaleChangeResults)
	ctx.Step(`^there should be no errors$`, thereShouldBeNoErrors)
}
