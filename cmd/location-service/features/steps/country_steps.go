package steps

import (
	"context"
	"github.com/cucumber/godog"
	"github.com/docker/go-connections/nat"
	"github.com/ocrosby/soccer/internal/database"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
)

var dbContainer testcontainers.Container
var locationServiceContainer testcontainers.Container

// var ctx context.Context
var dbPort, locationPort nat.Port

func iAmAnAuthorizedUser() error {
	// Implement authorization logic or mock
	return nil
}

func iSubmitARequestToAddANewCountryWithNameAndCode(name, code string) error {
	// Implement the logic to simulate adding a country
	return nil
}

func theCountryWithNameAndCodeShouldBeAddedToTheDatabase(name, code string) error {
	// Implement the logic to verify the country was added
	return nil
}

func iShouldReceiveTheDetailsOfTheCountry(arg1 string) error {
	return godog.ErrPending
}

func iSubmitARequestToDeleteTheCountryWithCode(arg1 string) error {
	return godog.ErrPending
}

func iSubmitARequestToRetrieveACountryWithCode(arg1 string) error {
	return godog.ErrPending
}

func iSubmitARequestToUpdateTheCountryWithCodeToChangeItsNameTo(arg1, arg2 string) error {
	return godog.ErrPending
}

func theCountryWithCodeExistsInTheDatabase(arg1 string) error {
	return godog.ErrPending
}

func theCountryWithCodeShouldBeRemovedFromTheDatabase(arg1 string) error {
	return godog.ErrPending
}

func theNameOfTheCountryWithCodeShouldBeUpdatedToInTheDatabase(arg1, arg2 string) error {
	return godog.ErrPending
}

func startPostgresContainer() (nat.Port, error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "user",
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"), // Wait until PostgreSQL is ready
	}

	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return "", err
	}

	port, err := container.MappedPort(context.Background(), "5432")
	if err != nil {
		return "", err
	}

	return port, nil
}

func startLocationServiceContainer(settings database.Settings) (nat.Port, error) {
	// Construct the database connection string
	dbConnectionString := settings.ConnectionString()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "./../",
			Dockerfile: "./Dockerfile",
		},
		ExposedPorts: []string{"8080/tcp"},
		WaitingFor:   wait.ForListeningPort("8080/tcp"),
		Env: map[string]string{
			"DB_CONNECTION_STRING": dbConnectionString,
		},
	}

	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return "", err
	}

	port, err := container.MappedPort(context.Background(), "8080")
	if err != nil {
		return "", err
	}

	return port, nil
}

func handleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func logContainerDetails(port nat.Port, containerName string) {
	log.Printf("Container %s is running on port %s", containerName, port.Port())
}

func InitializeCountriesTestSuite(ctx *godog.TestSuiteContext) {
	// BeforeSuite setup code
	ctx.BeforeSuite(func() {
		var err error

		settings := database.NewSettings("postgres", dbPort.Port(), "user", "password", "testdb", "disable")
		dbPort, err = startPostgresContainer()
		handleError(err, "Failed to start PostgreSQL container")
		logContainerDetails(dbPort, "PostgreSQL")

		locationPort, err = startLocationServiceContainer(*settings)
		handleError(err, "Failed to start Location Service container")
		logContainerDetails(locationPort, "Location Service")

		// Use dbPort in your database connection string
		// Example: "host=localhost port={dbPort.Port()} user=user password=password dbname=testdb sslmode=disable"
	})

	ctx.AfterSuite(func() {
		// Code to run after the test suite ends
		// For example, cleaning up the database or external resources
		if err := dbContainer.Terminate(context.Background()); err != nil {
			panic(err) // Handle error appropriately
		}

		if err := locationServiceContainer.Terminate(context.Background()); err != nil {
			log.Fatalf("Failed to terminate Location Service container: %v", err)
		}
	})
}

func InitializeCountriesScenario(ctx *godog.ScenarioContext) {
	ctx.StepContext().Before(func(ctx context.Context, st *godog.Step) (context.Context, error) {
		// Setup code before each step
		return nil, nil
	})

	ctx.StepContext().After(func(ctx context.Context, st *godog.Step, status godog.StepResultStatus, err error) (context.Context, error) {
		// Cleanup code after each step
		return nil, nil
	})

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		// Setup code before each scenario
		return nil, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		// Cleanup code after each scenario
		return nil, nil
	})

	ctx.Step(`^I am an authorized user$`, iAmAnAuthorizedUser)
	ctx.Step(`^I submit a request to add a new country with the name "([^"]*)" and code "([^"]*)"$`, iSubmitARequestToAddANewCountryWithNameAndCode)
	ctx.Step(`^the country "([^"]*)" with code "([^"]*)" should be added to the database$`, theCountryWithNameAndCodeShouldBeAddedToTheDatabase)
	ctx.Step(`^I should receive the details of the country "([^"]*)"$`, iShouldReceiveTheDetailsOfTheCountry)
	ctx.Step(`^I submit a request to delete the country with code "([^"]*)"$`, iSubmitARequestToDeleteTheCountryWithCode)
	ctx.Step(`^I submit a request to retrieve a country with code "([^"]*)"$`, iSubmitARequestToRetrieveACountryWithCode)
	ctx.Step(`^I submit a request to update the country with code "([^"]*)" to change its name to "([^"]*)"$`, iSubmitARequestToUpdateTheCountryWithCodeToChangeItsNameTo)
	ctx.Step(`^the country with code "([^"]*)" exists in the database$`, theCountryWithCodeExistsInTheDatabase)
	ctx.Step(`^the country with code "([^"]*)" should be removed from the database$`, theCountryWithCodeShouldBeRemovedFromTheDatabase)
	ctx.Step(`^the name of the country with code "([^"]*)" should be updated to "([^"]*)" in the database$`, theNameOfTheCountryWithCodeShouldBeUpdatedToInTheDatabase)
}
