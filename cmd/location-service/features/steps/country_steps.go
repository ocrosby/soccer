package steps

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/cucumber/godog"
	"github.com/docker/go-connections/nat"
	"github.com/ocrosby/soccer/internal/database"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	_ "github.com/lib/pq"
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

func executeDBScripts(db *sql.DB, scriptsDir string) error {
	var (
		files   []os.DirEntry
		content []byte
		err     error
	)

	if files, err = os.ReadDir(scriptsDir); err != nil {
		// Log the current directory at this point.
		// This will help you determine the location where the script is being executed.
		// You can use this information to debug the issue.
		dir, _ := os.Getwd()
		log.Printf("Current directory: %s", dir)
		return fmt.Errorf("reading scripts directory: %w", err)
	}

	// Sort files by name to ensure execution order
	// This is important if the scripts depend on each other
	// For example, creating tables before inserting data
	// You can also use a naming convention to ensure the correct order
	// For example, 01_create_table.sql, 02_insert_data.sql
	// This will ensure that the scripts are executed in the correct order
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, file := range files {
		// Skip files that are not SQL scripts
		if filepath.Ext(file.Name()) != ".sql" {
			continue
		}

		// Read the content of the SQL script
		if content, err = os.ReadFile(filepath.Join(scriptsDir, file.Name())); err != nil {
			return fmt.Errorf("reading file %s: %w", file.Name(), err)
		}

		// Execute the SQL script
		if _, err = db.Exec(string(content)); err != nil {
			return fmt.Errorf("executing script %s: %w", file.Name(), err)
		}

		log.Printf("Successfully executed script: %s", file.Name())
	}

	return nil
}

func startPostgresContainer() (testcontainers.Container, nat.Port, error) {
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
		return nil, "", err
	}

	port, err := container.MappedPort(context.Background(), "5432")
	if err != nil {
		return nil, "", err
	}

	// After the container is started and ready
	// Connect to the database and execute the SQL scripts
	db, err := sql.Open("postgres", fmt.Sprintf("host=localhost port=%s user=user password=password dbname=testdb sslmode=disable", port.Port()))
	if err != nil {
		return nil, "", fmt.Errorf("connecting to postgres: %w", err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Printf("Failed to close database connection: %v", err)
		}
	}(db)

	// Assuming your scripts are in "./scripts"
	if err = executeDBScripts(db, "./scripts"); err != nil {
		return nil, "", fmt.Errorf("executing database scripts: %w", err)
	}

	return container, port, nil
}

func startLocationServiceContainer(settings database.Settings) (testcontainers.Container, nat.Port, error) {
	// Construct the database connection string
	dbConnectionString := settings.ConnectionString()

	// https://golang.testcontainers.org/features/build_from_dockerfile/
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       "./../../../",
			Dockerfile:    "./cmd/location-service/Dockerfile",
			PrintBuildLog: true,
		},
		ExposedPorts: []string{"8080/tcp"},
		WaitingFor:   wait.ForListeningPort("8080/tcp").WithStartupTimeout(10 * time.Second),
		Env: map[string]string{
			"DB_CONNECTION_STRING": dbConnectionString,
		},
	}

	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return nil, "", err
	}

	port, err := container.MappedPort(context.Background(), "8080")
	if err != nil {
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

func InitializeCountriesTestSuite(ctx *godog.TestSuiteContext) {
	// BeforeSuite setup code
	ctx.BeforeSuite(func() {
		var err error

		settings := database.NewSettings("postgres", dbPort.Port(), "user", "password", "testdb", "disable")
		dbContainer, dbPort, err = startPostgresContainer()
		handleError(err, "Failed to start PostgreSQL container")
		logContainerDetails(dbPort, "PostgreSQL")

		locationServiceContainer, locationPort, err = startLocationServiceContainer(*settings)
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
