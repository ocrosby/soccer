package main

import (
	"github.com/cucumber/godog"
	steps "github.com/ocrosby/soccer/cmd/tds-college-conference-service/features/step_definitions"
	"os"
	"testing"
)

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		// This function will run once before the test suite is started
	})

	ctx.AfterSuite(func() {
		// This function will run once after the test suite is completed
	})
}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format:    "pretty",               // Change to any format you prefer
		Paths:     []string{"./features"}, // Path to your feature files
		Randomize: 0,                      // Optional: specify a seed to randomize the scenario execution order
	}

	status := godog.TestSuite{
		Name:                 "godogs",
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
